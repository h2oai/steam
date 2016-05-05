package bindings

type ModelsBase struct {
	*RequestSchema
	/** Name of Model of interest */
	ModelId *ModelKeyV3 `json:"model_id"`

	/** Return potentially abridged model suitable for viewing in a browser */
	Preview bool `json:"preview"`

	/** Find and return compatible frames? */
	FindCompatibleFrames bool `json:"find_compatible_frames"`

	/** Models */
	Models []*ModelSchemaBase `json:"models"`

	/** Compatible frames */
	CompatibleFrames []*FrameV3 `json:"compatible_frames"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func newModelsBase() *ModelsBase {
	return &ModelsBase{
		ModelId:              nil,
		Preview:              false,
		FindCompatibleFrames: false,
		Models:               nil,
		CompatibleFrames:     nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
