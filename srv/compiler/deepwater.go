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
package compiler

import (
	"mime/multipart"

	"github.com/h2oai/steam/lib/fs"
	"github.com/pkg/errors"
)

const fileTypeDWDep = "deepwater"

type Deepwater struct {
	*Model

	deepwaterDep string
}

// NewDeepWater returns a new instance of Deepwater.
func NewDeepwater(workingDirectory string, modelId int64, logicalName, modelType string, pythonFiles pythonPackage) *Deepwater {
	return &Deepwater{
		Model:        NewModel(workingDirectory, modelId, logicalName, modelType, pythonFiles),
		deepwaterDep: fs.GetDeepwaterDepPath(workingDirectory, modelId),
	}
}

func (c *Deepwater) AttachFiles(w *multipart.Writer) error {
	if err := c.Model.AttachFiles(w); err != nil {
		return err
	}

	return errors.Wrap(
		attachFile(w, c.deepwaterDep, fileTypeDWDep),
		"attaching deepwater dependency",
	)
}
