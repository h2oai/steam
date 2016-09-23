/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
