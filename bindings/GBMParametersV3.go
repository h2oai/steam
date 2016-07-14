package bindings

import (
	"encoding/json"
	"math"
)

type GBMParametersV3 struct {
	*SharedTreeParametersV3
	/** Learning rate (from 0.0 to 1.0) */
	LearnRate float64 `json:"learn_rate,omitempty"`
	/** Scale down the learning rate by this factor after each tree */
	LearnRateAnnealing float64 `json:"learn_rate_annealing,omitempty"`
	/** Distribution function */
	Distribution DistributionFamily `json:"distribution,omitempty"`
	/** Desired quantile for quantile regression (from 0.0 to 1.0) */
	QuantileAlpha float64 `json:"quantile_alpha,omitempty"`
	/** Tweedie Power (between 1 and 2) */
	TweediePower float64 `json:"tweedie_power,omitempty"`
	/** Column sample rate (from 0.0 to 1.0) */
	ColSampleRate float64 `json:"col_sample_rate,omitempty"`
	/** Maximum absolute value of a leaf node prediction */
	MaxAbsLeafnodePred float64 `json:"max_abs_leafnode_pred,omitempty"`
	/* INHERITED: Balance training data class counts via over/under-sampling (for imbalanced data).
	BalanceClasses bool `json:"balance_classes,omitempty"`
	*/
	/* INHERITED: Desired over/under-sampling ratios per class (in lexicographic order). If not specified, sampling factors will be automatically computed to obtain class balance during training. Requires balance_classes.
	ClassSamplingFactors []float32 `json:"class_sampling_factors,omitempty"`
	*/
	/* INHERITED: Maximum relative size of the training data after balancing class counts (can be less than 1.0). Requires balance_classes.
	MaxAfterBalanceSize float32 `json:"max_after_balance_size,omitempty"`
	*/
	/* INHERITED: Maximum size (# classes) for confusion matrices to be printed in the Logs
	MaxConfusionMatrixSize int32 `json:"max_confusion_matrix_size,omitempty"`
	*/
	/* INHERITED: Max. number (top K) of predictions to use for hit ratio computation (for multi-class only, 0 to disable)
	MaxHitRatioK int32 `json:"max_hit_ratio_k,omitempty"`
	*/
	/* INHERITED: Number of trees.
	Ntrees int32 `json:"ntrees,omitempty"`
	*/
	/* INHERITED: Maximum tree depth.
	MaxDepth int32 `json:"max_depth,omitempty"`
	*/
	/* INHERITED: Fewest allowed (weighted) observations in a leaf (in R called 'nodesize').
	MinRows float64 `json:"min_rows,omitempty"`
	*/
	/* INHERITED: For numerical columns (real/int), build a histogram of (at least) this many bins, then split at the best point
	Nbins int32 `json:"nbins,omitempty"`
	*/
	/* INHERITED: For numerical columns (real/int), build a histogram of (at most) this many bins at the root level, then decrease by factor of two per level
	NbinsTopLevel int32 `json:"nbins_top_level,omitempty"`
	*/
	/* INHERITED: For categorical columns (factors), build a histogram of this many bins, then split at the best point. Higher values can lead to more overfitting.
	NbinsCats int32 `json:"nbins_cats,omitempty"`
	*/
	/* INHERITED: Stop making trees when the R^2 metric equals or exceeds this
	R2Stopping float64 `json:"r2_stopping,omitempty"`
	*/
	/* INHERITED: Seed for pseudo random number generator (if applicable)
	Seed int64 `json:"seed,omitempty"`
	*/
	/* INHERITED: Run on one node only; no network overhead but fewer cpus used.  Suitable for small datasets.
	BuildTreeOneNode bool `json:"build_tree_one_node,omitempty"`
	*/
	/* INHERITED: Row sample rate per tree (from 0.0 to 1.0)
	SampleRate float64 `json:"sample_rate,omitempty"`
	*/
	/* INHERITED: Row sample rate per tree per class (from 0.0 to 1.0)
	SampleRatePerClass []float64 `json:"sample_rate_per_class,omitempty"`
	*/
	/* INHERITED: Column sample rate per tree (from 0.0 to 1.0)
	ColSampleRatePerTree float64 `json:"col_sample_rate_per_tree,omitempty"`
	*/
	/* INHERITED: Relative change of the column sampling rate for every level (from 0.0 to 2.0)
	ColSampleRateChangePerLevel float64 `json:"col_sample_rate_change_per_level,omitempty"`
	*/
	/* INHERITED: Score the model after every so many trees. Disabled if set to 0.
	ScoreTreeInterval int32 `json:"score_tree_interval,omitempty"`
	*/
	/* INHERITED: Minimum relative improvement in squared error reduction for a split to happen.
	MinSplitImprovement float64 `json:"min_split_improvement,omitempty"`
	*/
	/* INHERITED: Whether to use random split points for histograms (to pick the best split from).
	RandomSplitPoints bool `json:"random_split_points,omitempty"`
	*/
	/* INHERITED: Destination id for this model; auto-generated if not specified
	ModelId *ModelKeyV3 `json:"model_id,omitempty"`
	*/
	/* INHERITED: Training frame
	TrainingFrame *FrameKeyV3 `json:"training_frame,omitempty"`
	*/
	/* INHERITED: Validation frame
	ValidationFrame *FrameKeyV3 `json:"validation_frame,omitempty"`
	*/
	/* INHERITED: Number of folds for N-fold cross-validation
	Nfolds int32 `json:"nfolds,omitempty"`
	*/
	/* INHERITED: Keep cross-validation model predictions
	KeepCrossValidationPredictions bool `json:"keep_cross_validation_predictions,omitempty"`
	*/
	/* INHERITED: Keep cross-validation fold assignment
	KeepCrossValidationFoldAssignment bool `json:"keep_cross_validation_fold_assignment,omitempty"`
	*/
	/* INHERITED: Allow parallel training of cross-validation models
	ParallelizeCrossValidation bool `json:"parallelize_cross_validation,omitempty"`
	*/
	/* INHERITED: Response column
	ResponseColumn *ColSpecifierV3 `json:"response_column,omitempty"`
	*/
	/* INHERITED: Column with observation weights
	WeightsColumn *ColSpecifierV3 `json:"weights_column,omitempty"`
	*/
	/* INHERITED: Offset column
	OffsetColumn *ColSpecifierV3 `json:"offset_column,omitempty"`
	*/
	/* INHERITED: Column with cross-validation fold index assignment per observation
	FoldColumn *ColSpecifierV3 `json:"fold_column,omitempty"`
	*/
	/* INHERITED: Cross-validation fold assignment scheme, if fold_column is not specified
	FoldAssignment ModelParametersFoldAssignmentScheme `json:"fold_assignment,omitempty"`
	*/
	/* INHERITED: Ignored columns
	IgnoredColumns []string `json:"ignored_columns,omitempty"`
	*/
	/* INHERITED: Ignore constant columns
	IgnoreConstCols bool `json:"ignore_const_cols,omitempty"`
	*/
	/* INHERITED: Whether to score during each iteration of model training
	ScoreEachIteration bool `json:"score_each_iteration,omitempty"`
	*/
	/* INHERITED: Model checkpoint to resume training with
	Checkpoint *ModelKeyV3 `json:"checkpoint,omitempty"`
	*/
	/* INHERITED: Early stopping based on convergence of stopping_metric. Stop if simple moving average of length k of the stopping_metric does not improve for k:=stopping_rounds scoring events (0 to disable)
	StoppingRounds int32 `json:"stopping_rounds,omitempty"`
	*/
	/* INHERITED: Maximum allowed runtime in seconds for model training. Use 0 to disable.
	MaxRuntimeSecs float64 `json:"max_runtime_secs,omitempty"`
	*/
	/* INHERITED: Metric to use for early stopping (AUTO: logloss for classification, deviance for regression)
	StoppingMetric ScoreKeeperStoppingMetric `json:"stopping_metric,omitempty"`
	*/
	/* INHERITED: Relative tolerance for metric-based stopping criterion (stop if relative improvement is not at least this much)
	StoppingTolerance float64 `json:"stopping_tolerance,omitempty"`
	*/
}

func NewGBMParametersV3() *GBMParametersV3 {
	return &GBMParametersV3{
		LearnRate:          0.1,
		LearnRateAnnealing: 1.0,
		Distribution:       DistributionFamilyAUTO,
		QuantileAlpha:      0.5,
		TweediePower:       1.5,
		ColSampleRate:      1.0,
		MaxAbsLeafnodePred: math.Inf(1),
		SharedTreeParametersV3: &SharedTreeParametersV3{
			BalanceClasses:              false,
			ClassSamplingFactors:        nil,
			MaxAfterBalanceSize:         5.0,
			MaxConfusionMatrixSize:      20,
			MaxHitRatioK:                0,
			Ntrees:                      50,
			MaxDepth:                    5,
			MinRows:                     10.0,
			Nbins:                       20,
			NbinsTopLevel:               1024,
			NbinsCats:                   1024,
			R2Stopping:                  0.999999,
			Seed:                        -1,
			BuildTreeOneNode:            false,
			SampleRate:                  1.0,
			SampleRatePerClass:          nil,
			ColSampleRatePerTree:        1.0,
			ColSampleRateChangePerLevel: 1.0,
			ScoreTreeInterval:           0,
			MinSplitImprovement:         0.0,
			RandomSplitPoints:           false,
			ModelParametersSchema: &ModelParametersSchema{
				ModelId:         nil,
				TrainingFrame:   nil,
				ValidationFrame: nil,
				Nfolds:          0,
				KeepCrossValidationPredictions:    false,
				KeepCrossValidationFoldAssignment: false,
				ParallelizeCrossValidation:        true,
				ResponseColumn:                    nil,
				WeightsColumn:                     nil,
				OffsetColumn:                      nil,
				FoldColumn:                        nil,
				FoldAssignment:                    ModelParametersFoldAssignmentSchemeAUTO,
				IgnoredColumns:                    nil,
				IgnoreConstCols:                   true,
				ScoreEachIteration:                false,
				Checkpoint:                        nil,
				StoppingRounds:                    0,
				MaxRuntimeSecs:                    0.0,
				StoppingMetric:                    ScoreKeeperStoppingMetricAUTO,
				StoppingTolerance:                 0.001,
			},
		},
	}
}

// UnmarshalJSON to handle possible Infinity and NaN values
func (o *GBMParametersV3) UnmarshalJSON(data []byte) error {
	type Alias GBMParametersV3
	aux := &struct {
		LearnRate                   interface{}   `json:"learn_rate,omitempty"`
		LearnRateAnnealing          interface{}   `json:"learn_rate_annealing,omitempty"`
		QuantileAlpha               interface{}   `json:"quantile_alpha,omitempty"`
		TweediePower                interface{}   `json:"tweedie_power,omitempty"`
		ColSampleRate               interface{}   `json:"col_sample_rate,omitempty"`
		MaxAbsLeafnodePred          interface{}   `json:"max_abs_leafnode_pred,omitempty"`
		ClassSamplingFactors        []interface{} `json:"class_sampling_factors,omitempty"`
		MaxAfterBalanceSize         interface{}   `json:"max_after_balance_size,omitempty"`
		MinRows                     interface{}   `json:"min_rows,omitempty"`
		R2Stopping                  interface{}   `json:"r2_stopping,omitempty"`
		SampleRate                  interface{}   `json:"sample_rate,omitempty"`
		SampleRatePerClass          []interface{} `json:"sample_rate_per_class,omitempty"`
		ColSampleRatePerTree        interface{}   `json:"col_sample_rate_per_tree,omitempty"`
		ColSampleRateChangePerLevel interface{}   `json:"col_sample_rate_change_per_level,omitempty"`
		MinSplitImprovement         interface{}   `json:"min_split_improvement,omitempty"`
		MaxRuntimeSecs              interface{}   `json:"max_runtime_secs,omitempty"`
		StoppingTolerance           interface{}   `json:"stopping_tolerance,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.LearnRate = jsonToDoubl(aux.LearnRate)
	o.LearnRateAnnealing = jsonToDoubl(aux.LearnRateAnnealing)
	o.QuantileAlpha = jsonToDoubl(aux.QuantileAlpha)
	o.TweediePower = jsonToDoubl(aux.TweediePower)
	o.ColSampleRate = jsonToDoubl(aux.ColSampleRate)
	o.MaxAbsLeafnodePred = jsonToDoubl(aux.MaxAbsLeafnodePred)
	o.ClassSamplingFactors = jsonToFloats(aux.ClassSamplingFactors)
	o.MaxAfterBalanceSize = jsonToFloat(aux.MaxAfterBalanceSize)
	o.MinRows = jsonToDoubl(aux.MinRows)
	o.R2Stopping = jsonToDoubl(aux.R2Stopping)
	o.SampleRate = jsonToDoubl(aux.SampleRate)
	o.SampleRatePerClass = jsonToDoubls(aux.SampleRatePerClass)
	o.ColSampleRatePerTree = jsonToDoubl(aux.ColSampleRatePerTree)
	o.ColSampleRateChangePerLevel = jsonToDoubl(aux.ColSampleRateChangePerLevel)
	o.MinSplitImprovement = jsonToDoubl(aux.MinSplitImprovement)
	o.MaxRuntimeSecs = jsonToDoubl(aux.MaxRuntimeSecs)
	o.StoppingTolerance = jsonToDoubl(aux.StoppingTolerance)
	return nil
}
