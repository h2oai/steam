package yarn

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/h2oai/steam/lib/kerberos"
	"github.com/h2oai/steam/master/data"
	"github.com/pkg/errors"
)

func jobList(keytabFile, principal string, uid, gid uint32) (map[string]struct{}, error) {
	// Create timeout for hadoop job -list
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// Issue hadoop job -list to get job ids
	cmd := exec.CommandContext(ctx, "hadoop", "job", "-list")
	if keytabFile != "" {
		if err := kerberos.Kinit(keytabFile, principal, uid, gid); err != nil {
			return nil, errors.Wrap(err, "kinit")
		}
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid}
	}
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
	ReadKeytab(...data.QueryOpt) (data.Keytab, bool, error)
}

type Poll struct {
	ds *data.Datastore

	workingDirectory string
	started          string
	failed           string
}

type viewKerberosFn func(*data.Datastore) (bool, error)

func StartPoll(ds *data.Datastore, startedState, failedState, workgingDir string, viewKerb viewKerberosFn) error {
	poll := Poll{ds: ds, started: startedState, failed: failedState, workingDirectory: workgingDir}
	cmd := exec.Command("hadoop", "version")
	if err := cmd.Run(); err != nil {
		log.Println("Hadoop executable not found; Steam will run without polling")
		return nil
	}

	t := time.NewTimer(time.Minute)
	for {
		kEnable, err := viewKerb(poll.ds)
		if err != nil {
			log.Println("POLL", err, "trying withouth kerberos")
		}
		if err := poll.pollFunc(kEnable); err != nil {
			log.Println("Poll ERROR", err)
		}
		// Wait for a minute before next poll
		<-t.C
		t.Reset(time.Minute)
	}
}

func (p *Poll) pollFunc(kerberosEnabled bool) error {
	var ktPath, principal string
	var uid, gid uint32
	if kerberosEnabled {
		keytab, exists, err := p.ds.ReadKeytab(data.ByPrincipalSteam)
		if err != nil {
			return errors.Wrap(err, "reading keytab")
		} else if exists {
			principal = keytab.Principal.String
			uid, gid, err = GetUser(principal)
			if err != nil {
				return errors.Wrap(err, "getting user")
			}
			ktPath, err = kerberos.WriteKeytab(keytab, p.workingDirectory, int(uid), int(gid))
			if err != nil {
				return errors.Wrap(err, "writing keytab")
			}
		}
	}
	// Retrieve job Ids
	jobIds, err := jobList(ktPath, principal, uid, gid)
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
			log.Println("POLL", detail.ApplicationId, "Yarn state mismatch with Steam state")
			cluster, exists, err := p.ds.ReadCluster(data.ByDetailId(detail.Id))
			if err != nil {
				return errors.Wrap(err, "reading cluster from database")
			} else if exists {
				log.Printf("POLL %s user: %s cluster: %s stopped externally; changing Steam state",
					detail.ApplicationId, cluster.Username.String, cluster.Name)
				if err := p.ds.UpdateCluster(cluster.Id, data.WithState(p.failed)); err != nil {
					return errors.Wrap(err, "updating cluster state")
				}
			}
		}
	}

	return nil
}
