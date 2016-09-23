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

type ModelMetricsListSchemaV3 struct {
	*RequestSchema
	/** Key of Model of interest (optional) */
	Model *ModelKeyV3 `json:"model,omitempty"`
	/** Key of Frame of interest (optional) */
	Frame *FrameKeyV3 `json:"frame,omitempty"`
	/** Key of predictions frame, if predictions are requested (optional) */
	PredictionsFrame *FrameKeyV3 `json:"predictions_frame,omitempty"`
	/** Compute reconstruction error (optional, only for Deep Learning AutoEncoder models) */
	ReconstructionError bool `json:"reconstruction_error,omitempty"`
	/** Compute reconstruction error per feature (optional, only for Deep Learning AutoEncoder models) */
	ReconstructionErrorPerFeature bool `json:"reconstruction_error_per_feature,omitempty"`
	/** Extract Deep Features for given hidden layer (optional, only for Deep Learning models) */
	DeepFeaturesHiddenLayer int32 `json:"deep_features_hidden_layer,omitempty"`
	/** Reconstruct original training frame (optional, only for GLRM models) */
	ReconstructTrain bool `json:"reconstruct_train,omitempty"`
	/** Project GLRM archetypes back into original feature space (optional, only for GLRM models) */
	ProjectArchetypes bool `json:"project_archetypes,omitempty"`
	/** Reverse transformation applied during training to model output (optional, only for GLRM models) */
	ReverseTransform bool `json:"reverse_transform,omitempty"`
	/** Return the leaf node assignment (optional, only for DRF/GBM models) */
	LeafNodeAssignment bool `json:"leaf_node_assignment,omitempty"`
	/** Retrieve all members for a given exemplar (optional, only for Aggregator models) */
	ExemplarIndex int32 `json:"exemplar_index,omitempty"`
	/** ModelMetrics */
	ModelMetrics []*ModelMetricsBase `json:"model_metrics,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewModelMetricsListSchemaV3() *ModelMetricsListSchemaV3 {
	return &ModelMetricsListSchemaV3{
		Model:                         nil,
		Frame:                         nil,
		PredictionsFrame:              nil,
		ReconstructionError:           false,
		ReconstructionErrorPerFeature: false,
		DeepFeaturesHiddenLayer:       -1,
		ReconstructTrain:              false,
		ProjectArchetypes:             false,
		ReverseTransform:              false,
		LeafNodeAssignment:            false,
		ExemplarIndex:                 -1,
		ModelMetrics:                  nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
