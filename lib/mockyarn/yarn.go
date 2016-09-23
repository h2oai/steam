package yarn

import (
	"github.com/h2oai/steam/lib/fs"
	"math/rand"
	"strconv"
)

// StartCloud starts a yarn cloud by shelling out to hadoop
//
// This process needs to store the job-ID to kill the process in the future
func StartCloud(size int, kerberos bool, mem, name, username, keytab string) (string, string, error) {

	applicationID, err := fs.NewID()
	if err != nil {
		return "", "", err
	}
	address := strconv.Itoa(192+rand.Intn(63)) + "." + strconv.Itoa(192+rand.Intn(63)) + "." + strconv.Itoa(192+rand.Intn(63)) + "." + strconv.Itoa(192+rand.Intn(63))

	return applicationID, address, nil
}

// StopCloud kills a hadoop cloud by shelling out a command based on the job-ID
//
// In the future this
func StopCloud(kerberos bool, name, id, username, keytab string) error {
	return nil
}
