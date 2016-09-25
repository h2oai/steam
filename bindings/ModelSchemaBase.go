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

package bindings

type ModelSchemaBase struct {
	*Schema
	/** Model key */
	ModelId *ModelKeyV3 `json:"model_id"`
	/** The algo name for this Model. */
	Algo string `json:"algo,omitempty"`
	/** The pretty algo name for this Model (e.g., Generalized Linear Model, rather than GLM). */
	AlgoFullName string `json:"algo_full_name,omitempty"`
	/** The response column name for this Model (if applicable). Is null otherwise. */
	ResponseColumnName string `json:"response_column_name,omitempty"`
	/** The Model's training frame key */
	DataFrame *FrameKeyV3 `json:"data_frame,omitempty"`
	/** Timestamp for when this model was completed */
	Timestamp int64 `json:"timestamp,omitempty"`
}

func NewModelSchemaBase() *ModelSchemaBase {
	return &ModelSchemaBase{
		ModelId:            nil,
		Algo:               "",
		AlgoFullName:       "",
		ResponseColumnName: "",
		DataFrame:          nil,
		Timestamp:          0,
		Schema:             &Schema{},
	}
}
