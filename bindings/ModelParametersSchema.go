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

import "encoding/json"

type ModelParametersSchema struct {
	*Schema
	/** Destination id for this model; auto-generated if not specified */
	ModelId *ModelKeyV3 `json:"model_id,omitempty"`
	/** Training frame */
	TrainingFrame *FrameKeyV3 `json:"training_frame,omitempty"`
	/** Validation frame */
	ValidationFrame *FrameKeyV3 `json:"validation_frame,omitempty"`
	/** Number of folds for N-fold cross-validation */
	Nfolds int32 `json:"nfolds,omitempty"`
	/** Keep cross-validation model predictions */
	KeepCrossValidationPredictions bool `json:"keep_cross_validation_predictions,omitempty"`
	/** Keep cross-validation fold assignment */
	KeepCrossValidationFoldAssignment bool `json:"keep_cross_validation_fold_assignment,omitempty"`
	/** Allow parallel training of cross-validation models */
	ParallelizeCrossValidation bool `json:"parallelize_cross_validation,omitempty"`
	/** Response column */
	ResponseColumn *ColSpecifierV3 `json:"response_column,omitempty"`
	/** Column with observation weights */
	WeightsColumn *ColSpecifierV3 `json:"weights_column,omitempty"`
	/** Offset column */
	OffsetColumn *ColSpecifierV3 `json:"offset_column,omitempty"`
	/** Column with cross-validation fold index assignment per observation */
	FoldColumn *ColSpecifierV3 `json:"fold_column,omitempty"`
	/** Cross-validation fold assignment scheme, if fold_column is not specified */
	FoldAssignment ModelParametersFoldAssignmentScheme `json:"fold_assignment,omitempty"`
	/** Ignored columns */
	IgnoredColumns []string `json:"ignored_columns,omitempty"`
	/** Ignore constant columns */
	IgnoreConstCols bool `json:"ignore_const_cols,omitempty"`
	/** Whether to score during each iteration of model training */
	ScoreEachIteration bool `json:"score_each_iteration,omitempty"`
	/** Model checkpoint to resume training with */
	Checkpoint *ModelKeyV3 `json:"checkpoint,omitempty"`
	/** Early stopping based on convergence of stopping_metric. Stop if simple moving average of length k of the stopping_metric does not improve for k:=stopping_rounds scoring events (0 to disable) */
	StoppingRounds int32 `json:"stopping_rounds,omitempty"`
	/** Maximum allowed runtime in seconds for model training. Use 0 to disable. */
	MaxRuntimeSecs float64 `json:"max_runtime_secs,omitempty"`
	/** Metric to use for early stopping (AUTO: logloss for classification, deviance for regression) */
	StoppingMetric ScoreKeeperStoppingMetric `json:"stopping_metric,omitempty"`
	/** Relative tolerance for metric-based stopping criterion (stop if relative improvement is not at least this much) */
	StoppingTolerance float64 `json:"stopping_tolerance,omitempty"`
}

func NewModelParametersSchema() *ModelParametersSchema {
	return &ModelParametersSchema{
		ModelId:         nil,
		TrainingFrame:   nil,
		ValidationFrame: nil,
		Nfolds:          0,
		KeepCrossValidationPredictions:    false,
		KeepCrossValidationFoldAssignment: false,
		ParallelizeCrossValidation:        false,
		ResponseColumn:                    nil,
		WeightsColumn:                     nil,
		OffsetColumn:                      nil,
		FoldColumn:                        nil,
		FoldAssignment:                    ModelParametersFoldAssignmentScheme_NONE_,
		IgnoredColumns:                    nil,
		IgnoreConstCols:                   false,
		ScoreEachIteration:                false,
		Checkpoint:                        nil,
		StoppingRounds:                    0,
		MaxRuntimeSecs:                    0.0,
		StoppingMetric:                    ScoreKeeperStoppingMetric_NONE_,
		StoppingTolerance:                 0.0,
		Schema:                            &Schema{},
	}
}

// UnmarshalJSON to handle possible Infinity and NaN values
func (o *ModelParametersSchema) UnmarshalJSON(data []byte) error {
	type Alias ModelParametersSchema
	aux := &struct {
		MaxRuntimeSecs    interface{} `json:"max_runtime_secs,omitempty"`
		StoppingTolerance interface{} `json:"stopping_tolerance,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.MaxRuntimeSecs = jsonToDoubl(aux.MaxRuntimeSecs)
	o.StoppingTolerance = jsonToDoubl(aux.StoppingTolerance)
	return nil
}
