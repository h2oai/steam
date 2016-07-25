package bindings

type ModelSchema struct {
	*ModelSchemaBase
	/** The build parameters for the model (e.g. K for KMeans). */
	Parameters *ModelParametersSchema `json:"parameters,omitempty"`
	/** The build output for the model (e.g. the cluster centers for KMeans). */
	Output *ModelOutputSchema `json:"output,omitempty"`
	/** Compatible frames, if requested */
	CompatibleFrames []string `json:"compatible_frames,omitempty"`
	/** Checksum for all the things that go into building the Model. */
	Checksum int64 `json:"checksum,omitempty"`
	/* INHERITED: Model key
	ModelId *ModelKeyV3 `json:"model_id"`
	*/
	/* INHERITED: The algo name for this Model.
	Algo string `json:"algo,omitempty"`
	*/
	/* INHERITED: The pretty algo name for this Model (e.g., Generalized Linear Model, rather than GLM).
	AlgoFullName string `json:"algo_full_name,omitempty"`
	*/
	/* INHERITED: The response column name for this Model (if applicable). Is null otherwise.
	ResponseColumnName string `json:"response_column_name,omitempty"`
	*/
	/* INHERITED: The Model's training frame key
	DataFrame *FrameKeyV3 `json:"data_frame,omitempty"`
	*/
	/* INHERITED: Timestamp for when this model was completed
	Timestamp int64 `json:"timestamp,omitempty"`
	*/
}

func NewModelSchema() *ModelSchema {
	return &ModelSchema{
		Parameters:       nil,
		Output:           nil,
		CompatibleFrames: nil,
		Checksum:         0,
		ModelSchemaBase: &ModelSchemaBase{
			ModelId:            nil,
			Algo:               "",
			AlgoFullName:       "",
			ResponseColumnName: "",
			DataFrame:          nil,
			Timestamp:          0,
		},
	}
}
