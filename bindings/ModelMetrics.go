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
		Mse                  interface{} `json:"MSE,omitempty"`
		R2                   interface{} `json:"r2,omitempty"`
		Logloss              interface{} `json:"logloss,omitempty"`
		Auc                  interface{} `json:"AUC,omitempty"`
		Gini                 interface{} `json:"Gini,omitempty"`
		MeanResidualDeviance interface{} `json:"mean_residual_deviance,omitempty"`
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
