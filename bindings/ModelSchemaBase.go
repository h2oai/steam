package bindings

type ModelSchemaBase struct {
	*Schema
	/** Model key */
	ModelId *ModelKeyV3 `json:"model_id"`

	/** The algo name for this Model. */
	Algo string `json:"algo"`

	/** The pretty algo name for this Model (e.g., Generalized Linear Model, rather than GLM). */
	AlgoFullName string `json:"algo_full_name"`

	/** The response column name for this Model (if applicable). Is null otherwise. */
	ResponseColumnName string `json:"response_column_name"`

	/** The Model's training frame key */
	DataFrame *FrameKeyV3 `json:"data_frame"`

	/** Timestamp for when this model was completed */
	Timestamp int64 `json:"timestamp"`
}

func newModelSchemaBase() *ModelSchemaBase {
	return &ModelSchemaBase{
		ModelId:            nil,
		Algo:               "",
		AlgoFullName:       "",
		ResponseColumnName: "",
		DataFrame:          nil,
		Timestamp:          0,
	}
}
