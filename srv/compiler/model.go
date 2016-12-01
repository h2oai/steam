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
	"fmt"
	"mime/multipart"

	"github.com/h2oai/steam/lib/fs"
	"github.com/pkg/errors"
)

type Model struct {
	modelPath string
	modelType string
	javaDep   string

	pythonFiles pythonPackage
}

func NewModel(workingDirectory string, modelId int64, logicalName, modelType string, pythonFiles pythonPackage) *Model {
	m := &Model{javaDep: fs.GetGenModelPath(workingDirectory, modelId)}

	switch modelType {
	case "pojo":
		m.modelPath = fs.GetJavaModelPath(workingDirectory, modelId, logicalName)
		m.modelType = fileTypeJava
	case "mojo":
		m.modelPath = fs.GetMOJOPath(workingDirectory, modelId, logicalName)
		m.modelType = fileTypeMOJO
	case "":
		panic("model type unset")
	default:
		panic(fmt.Errorf("invalid model type %q", modelType))
	}

	m.pythonFiles = pythonFiles

	return m
}

func (c *Model) AttachFiles(w *multipart.Writer) error {
	// Attach Java files
	if err := attachFile(w, c.modelPath, c.modelType); err != nil {
		return errors.Wrap(err, "attaching model")
	}
	if err := attachFile(w, c.javaDep, fileTypeJavaDep); err != nil {
		return errors.Wrap(err, "attaching java dependency")
	}

	// Attach Python Package Files
	if c.pythonFiles.Main != "" {
		if err := attachFile(w, c.pythonFiles.Main, fileTypePythonMain); err != nil {
			return errors.Wrap(err, "attaching Python main file")
		}
		for _, file := range c.pythonFiles.Other {
			if err := attachFile(w, file, fileTypePythonOther); err != nil {
				return errors.Wrap(err, "attaching Python file")
			}
		}
		if c.pythonFiles.Yaml != "" {
			if err := attachFile(w, c.pythonFiles.Yaml, fileTypePythonEnv); err != nil {
				return errors.Wrap(err, "attaching Python env file")
			}
		}
	}
	return nil
}
