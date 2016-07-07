package bindings

import "encoding/json"

type ModelMetricsBase struct {
	*Schema
	/** The model used for this scoring run. */
	Model *ModelKeyV3 `json:"model,omitempty"`
	/** The checksum for the model used for this scoring run. */
	ModelChecksum int64 `json:"model_checksum,omitempty"`
	/** The frame used for this scoring run. */
	Frame *FrameKeyV3 `json:"frame,omitempty"`
	/** The checksum for the frame used for this scoring run. */
	FrameChecksum int64 `json:"frame_checksum,omitempty"`
	/** Optional description for this scoring run (to note out-of-bag, sampled data, etc.) */
	Description string `json:"description,omitempty"`
	/** The category (e.g., Clustering) for the model used for this scoring run. */
	ModelCategory ModelCategory `json:"model_category,omitempty"`
	/** The time in mS since the epoch for the start of this scoring run. */
	ScoringTime int64 `json:"scoring_time,omitempty"`
	/** Predictions Frame. */
	Predictions *FrameV3 `json:"predictions,omitempty"`
	/** The Mean Squared Error of the prediction for this scoring run. */
	Mse float64 `json:"MSE,omitempty"`
}

func NewModelMetricsBase() *ModelMetricsBase {
	return &ModelMetricsBase{
		Model:         nil,
		ModelChecksum: 0,
		Frame:         nil,
		FrameChecksum: 0,
		Description:   "",
		ModelCategory: 0,
		ScoringTime:   0,
		Predictions:   nil,
		Mse:           0.0,
		Schema:        &Schema{},
	}
}

// UnmarshalJSON to handle possible Infinity and NaN values
func (o *ModelMetricsBase) UnmarshalJSON(data []byte) error {
	type Alias ModelMetricsBase
	aux := &struct {
		Mse interface{} `json:"MSE,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.Mse = jsonToDoubl(aux.Mse)
	return nil
}
