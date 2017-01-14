package yarn

import (
	"bufio"
	"os/exec"
	"strings"
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
