package bindings

import "encoding/json"

// This bypasses ModelMetricsBase used to get metric scalars
type ModelMetrics struct {
	Mse                  float64 `json:"MSE,omitempty"`
	R2                   float64 `json:"r2,omitempty"`
	Logloss              float64 `json:"logloss,omitempty"`
	Auc                  float64 `json:"AUC,omitempty"`
	Gini                 float64 `json:"Gini,omitempty"`
	MeanResidualDeviance float64 `json:"mean_residual_deviance,omitempty"`
}

func (o *ModelMetrics) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Mse                  interface{}
		R2                   interface{}
		Logloss              interface{}
		Auc                  interface{}
		Gini                 interface{}
		MeanResidualDeviance interface{}
	}{
		o.Mse,
		o.R2,
		o.Logloss,
		o.Auc,
		o.Gini,
		o.MeanResidualDeviance,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	o.Mse = jsonToDoubl(aux.Mse)
	o.R2 = jsonToDoubl(aux.R2)
	o.Logloss = jsonToDoubl(aux.Logloss)
	o.Auc = jsonToDoubl(aux.Auc)
	o.Gini = jsonToDoubl(aux.Gini)
	o.MeanResidualDeviance = jsonToDoubl(aux.MeanResidualDeviance)
	return nil
}
