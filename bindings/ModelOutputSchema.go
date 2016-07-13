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
	TrainingMetrics *ModelMetricsBase `json:"training_metrics,omitempty"`
	/** Validation data model metrics */
	ValidationMetrics *ModelMetricsBase `json:"validation_metrics,omitempty"`
	/** Cross-validation model metrics */
	CrossValidationMetrics *ModelMetricsBase `json:"cross_validation_metrics,omitempty"`
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
