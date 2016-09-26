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

package web

import "testing"

func TestGetClusterDatasets(tt *testing.T) {
	t := newTest(tt)

	id, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	datasets, err := t.svc.GetDatasetsFromCluster(t.su, id)
	t.nil(err)

	// VERIFY INFORMATION IN DATASETS HERE (INTEGRATION)
	t.log("Datasets", datasets)
}

func TestGetClusterModels(tt *testing.T) {
	t := newTest(tt)

	id, err := t.svc.RegisterCluster(t.su, ClusterAddress)
	t.nil(err)

	models, err := t.svc.GetModelsFromCluster(t.su, id, h2oFrames[0].name)
	t.nil(err)

	// VERIFY INFORMATION HERE
	t.log("Models", models)
}
