package bindings

type ModelSchemaBase struct {
	*Schema
	/** Model key */
	ModelId *ModelKeyV3 `json:"model_id"`
	/** The algo name for this Model. */
	Algo string `json:"algo,omitempty"`
	/** The pretty algo name for this Model (e.g., Generalized Linear Model, rather than GLM). */
	AlgoFullName string `json:"algo_full_name,omitempty"`
	/** The response column name for this Model (if applicable). Is null otherwise. */
	ResponseColumnName string `json:"response_column_name,omitempty"`
	/** The Model's training frame key */
	DataFrame *FrameKeyV3 `json:"data_frame,omitempty"`
	/** Timestamp for when this model was completed */
	Timestamp int64 `json:"timestamp,omitempty"`
}

func NewModelSchemaBase() *ModelSchemaBase {
	return &ModelSchemaBase{
		ModelId:            nil,
		Algo:               "",
		AlgoFullName:       "",
		ResponseColumnName: "",
		DataFrame:          nil,
		Timestamp:          0,
		Schema:             &Schema{},
	}
}
