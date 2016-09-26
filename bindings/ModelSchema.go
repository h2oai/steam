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

type ModelSchema struct {
	*ModelSchemaBase
	/** The build parameters for the model (e.g. K for KMeans). */
	// FIXME hack: removed modelparameters because unused and problematic unmarshal (it's always empty anyway...)
	// Parameters *ModelParametersSchema `json:"parameters,omitempty"`
	/** The build output for the model (e.g. the cluster centers for KMeans). */
	Output *ModelOutputSchema `json:"output,omitempty"`
	/** Compatible frames, if requested */
	CompatibleFrames []string `json:"compatible_frames,omitempty"`
	/** Checksum for all the things that go into building the Model. */
	Checksum int64 `json:"checksum,omitempty"`
	/* INHERITED: Model key
	ModelId *ModelKeyV3 `json:"model_id"`
	*/
	/* INHERITED: The algo name for this Model.
	Algo string `json:"algo,omitempty"`
	*/
	/* INHERITED: The pretty algo name for this Model (e.g., Generalized Linear Model, rather than GLM).
	AlgoFullName string `json:"algo_full_name,omitempty"`
	*/
	/* INHERITED: The response column name for this Model (if applicable). Is null otherwise.
	ResponseColumnName string `json:"response_column_name,omitempty"`
	*/
	/* INHERITED: The Model's training frame key
	DataFrame *FrameKeyV3 `json:"data_frame,omitempty"`
	*/
	/* INHERITED: Timestamp for when this model was completed
	Timestamp int64 `json:"timestamp,omitempty"`
	*/
}

func NewModelSchema() *ModelSchema {
	return &ModelSchema{
		// Parameters:       nil,
		Output:           nil,
		CompatibleFrames: nil,
		Checksum:         0,
		ModelSchemaBase: &ModelSchemaBase{
			ModelId:            nil,
			Algo:               "",
			AlgoFullName:       "",
			ResponseColumnName: "",
			DataFrame:          nil,
			Timestamp:          0,
		},
	}
}
