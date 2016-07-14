package bindings

import "encoding/json"

type ModelMetricsBinomialV3 struct {
	*ModelMetricsBase
	/** The R^2 for this scoring run. */
	R2 float64 `json:"r2,omitempty"`
	/** The logarithmic loss for this scoring run. */
	Logloss float64 `json:"logloss,omitempty"`
	/** The AUC for this scoring run. */
	Auc float64 `json:"AUC,omitempty"`
	/** The Gini score for this scoring run. */
	Gini float64 `json:"Gini,omitempty"`
	/** The class labels of the response. */
	Domain []string `json:"domain,omitempty"`
	/** The Metrics for various thresholds. */
	ThresholdsAndMetricScores *TwoDimTableV3 `json:"thresholds_and_metric_scores,omitempty"`
	/** The Metrics for various criteria. */
	MaxCriteriaAndMetricScores *TwoDimTableV3 `json:"max_criteria_and_metric_scores,omitempty"`
	/** Gains and Lift table. */
	GainsLiftTable *TwoDimTableV3 `json:"gains_lift_table,omitempty"`
	/* INHERITED: The model used for this scoring run.
	Model *ModelKeyV3 `json:"model,omitempty"`
	*/
	/* INHERITED: The checksum for the model used for this scoring run.
	ModelChecksum int64 `json:"model_checksum,omitempty"`
	*/
	/* INHERITED: The frame used for this scoring run.
	Frame *FrameKeyV3 `json:"frame,omitempty"`
	*/
	/* INHERITED: The checksum for the frame used for this scoring run.
	FrameChecksum int64 `json:"frame_checksum,omitempty"`
	*/
	/* INHERITED: Optional description for this scoring run (to note out-of-bag, sampled data, etc.)
	Description string `json:"description,omitempty"`
	*/
	/* INHERITED: The category (e.g., Clustering) for the model used for this scoring run.
	ModelCategory ModelCategory `json:"model_category,omitempty"`
	*/
	/* INHERITED: The time in mS since the epoch for the start of this scoring run.
	ScoringTime int64 `json:"scoring_time,omitempty"`
	*/
	/* INHERITED: Predictions Frame.
	Predictions *FrameV3 `json:"predictions,omitempty"`
	*/
	/* INHERITED: The Mean Squared Error of the prediction for this scoring run.
	Mse float64 `json:"MSE,omitempty"`
	*/
}

func NewModelMetricsBinomialV3() *ModelMetricsBinomialV3 {
	return &ModelMetricsBinomialV3{
		R2:      0.0,
		Logloss: 0.0,
		Auc:     0.0,
		Gini:    0.0,
		Domain:  nil,
		ThresholdsAndMetricScores:  nil,
		MaxCriteriaAndMetricScores: nil,
		GainsLiftTable:             nil,
		ModelMetricsBase: &ModelMetricsBase{
			Model:         nil,
			ModelChecksum: 0,
			Frame:         nil,
			FrameChecksum: 0,
			Description:   "",
			ModelCategory: ModelCategory_NONE_,
			ScoringTime:   0,
			Predictions:   nil,
			Mse:           0.0,
		},
	}
}

// UnmarshalJSON to handle possible Infinity and NaN values
func (o *ModelMetricsBinomialV3) UnmarshalJSON(data []byte) error {
	type Alias ModelMetricsBinomialV3
	aux := &struct {
		R2      interface{} `json:"r2,omitempty"`
		Logloss interface{} `json:"logloss,omitempty"`
		Auc     interface{} `json:"AUC,omitempty"`
		Gini    interface{} `json:"Gini,omitempty"`
		Mse     interface{} `json:"MSE,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.R2 = jsonToDoubl(aux.R2)
	o.Logloss = jsonToDoubl(aux.Logloss)
	o.Auc = jsonToDoubl(aux.Auc)
	o.Gini = jsonToDoubl(aux.Gini)
	o.Mse = jsonToDoubl(aux.Mse)
	return nil
}
