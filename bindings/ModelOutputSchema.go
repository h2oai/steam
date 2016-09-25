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

import (
	"encoding/json"
	"fmt"
)

type ModelOutputSchema struct {
	*Schema
	/** Column names */
	Names []string `json:"names,omitempty"`
	/** Domains for categorical columns */
	Domains [][]string `json:"domains,omitempty"`
	/** Cross-validation models (model ids) */
	CrossValidationModels []*ModelKeyV3 `json:"cross_validation_models,omitempty"`
	/** Cross-validation predictions, one per cv model (deprecated, use cross_validation_holdout_predictions_frame_id instead) */
	CrossValidationPredictions []*FrameKeyV3 `json:"cross_validation_predictions,omitempty"`
	/** Cross-validation holdout predictions (full out-of-sample predictions on training data) */
	CrossValidationHoldoutPredictionsFrameId *FrameKeyV3 `json:"cross_validation_holdout_predictions_frame_id,omitempty"`
	/** Cross-validation fold assignment (each row is assigned to one holdout fold) */
	CrossValidationFoldAssignmentFrameId *FrameKeyV3 `json:"cross_validation_fold_assignment_frame_id,omitempty"`
	/** Category of the model (e.g., Binomial) */
	ModelCategory ModelCategory `json:"model_category,omitempty"`
	/** Model summary */
	ModelSummary *TwoDimTableBase `json:"model_summary,omitempty"`
	/** Scoring history */
	ScoringHistory *TwoDimTableBase `json:"scoring_history,omitempty"`
	/** Training data model metrics */
	TrainingMetrics *ModelMetrics `json:"training_metrics,omitempty"`
	/** Validation data model metrics */
	ValidationMetrics *ModelMetrics `json:"validation_metrics,omitempty"`
	/** Cross-validation model metrics */
	CrossValidationMetrics *ModelMetrics `json:"cross_validation_metrics,omitempty"`
	/** Cross-validation model metrics summary */
	CrossValidationMetricsSummary *TwoDimTableBase `json:"cross_validation_metrics_summary,omitempty"`
	/** Job status */
	Status string `json:"status,omitempty"`
	/** Start time in milliseconds */
	StartTime int64 `json:"start_time,omitempty"`
	/** End time in milliseconds */
	EndTime int64 `json:"end_time,omitempty"`
	/** Runtime in milliseconds */
	RunTime int64 `json:"run_time,omitempty"`
	/** Help information for output fields */
	Help map[string]string `json:"help,omitempty"`
}

func NewModelOutputSchema() *ModelOutputSchema {
	return &ModelOutputSchema{
		Names:                                    nil,
		Domains:                                  nil,
		CrossValidationModels:                    nil,
		CrossValidationPredictions:               nil,
		CrossValidationHoldoutPredictionsFrameId: nil,
		CrossValidationFoldAssignmentFrameId:     nil,
		ModelCategory:                            ModelCategory_NONE_,
		ModelSummary:                             nil,
		ScoringHistory:                           nil,
		TrainingMetrics:                          nil,
		ValidationMetrics:                        nil,
		CrossValidationMetrics:                   nil,
		CrossValidationMetricsSummary:            nil,
		Status:    "",
		StartTime: 0,
		EndTime:   0,
		RunTime:   0,
		Help:      nil,
		Schema:    &Schema{},
	}
}

// ToString returns the contents of this object as a JSON String.
func (o *ModelOutputSchema) ToString() string {
	j, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return fmt.Sprint(err)
	}
	return string(j)
}
