package yarn

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/h2oai/steam/master/data"
	"github.com/pkg/errors"
)

func jobList() (map[string]struct{}, error) {
	// Issue hadoop job -list to get job ids
	cmd := exec.Command("hadoop", "job", "-list")
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdOut.Close()

	// Scan the output for the job ids
	jobIds := make([]string, 0)
	go func() {
		scan := bufio.NewScanner(stdOut)
		scan.Split(bufio.ScanWords)

		for scan.Scan() {
			if strings.Contains(scan.Text(), "job_") {
				jobIds = append(jobIds, scan.Text())
			}
		}
	}()
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Make a map of the actual job ids.
	ret := make(map[string]struct{})
	for _, jobId := range jobIds {
		ret[jobId] = struct{}{}
	}
	return ret, nil
}

type clusterDatabase interface {
	ReadCluster(...data.QueryOpt) (data.Cluster, bool, error)
	ReadClusterYarnDetails(...data.QueryOpt) ([]data.ClusterYarnDetail, error)
	UpdateCluster(int64, ...data.QueryOpt) error
}

type Poll struct {
	ds clusterDatabase

	started string
	failed  string
}

func StartPoll(ds clusterDatabase, startedState, failedState string) error {
	poll := Poll{ds: ds, started: startedState, failed: failedState}

	t := time.NewTimer(time.Minute)
	for {
		if err := poll.pollFunc(); err != nil {
			return err
		}
		// Wait for a minute before next poll
		<-t.C
		t.Reset(time.Minute)
	}
}

func (p *Poll) pollFunc() error {
	// Retrieve job Ids
	jobIds, err := jobList()
	if err != nil {
		return errors.Wrap(err, "getting yarn job list")
	}

	// Get all yarn details of jobs in the started state
	details, err := p.ds.ReadClusterYarnDetails(data.ByState(p.started))
	if err != nil {
		return errors.Wrap(err, "reading yarn details from database")
	}
	// Compare "started" jobs (in steam) with "running" jobs (in yarn). In case
	// of a mismatch, always defer to yarn
	for _, detail := range details {
		if _, ok := jobIds[fmt.Sprintf("job_%s", detail.ApplicationId)]; !ok {
			cluster, exists, err := p.ds.ReadCluster(data.ByDetailId(detail.Id))
			if err != nil {
				return errors.Wrap(err, "reading cluster from database")
			} else if exists {
				if err := p.ds.UpdateCluster(cluster.Id, data.WithState(p.failed)); err != nil {
					return errors.Wrap(err, "updating cluster state")
				}
			}
		}
	}

	return nil
}
