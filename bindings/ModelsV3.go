package bindings

type ModelsV3 struct {
	*ModelsBase
	/* INHERITED: Name of Model of interest
	ModelId *ModelKeyV3 `json:"model_id,omitempty"`
	*/
	/* INHERITED: Return potentially abridged model suitable for viewing in a browser
	Preview bool `json:"preview,omitempty"`
	*/
	/* INHERITED: Find and return compatible frames?
	FindCompatibleFrames bool `json:"find_compatible_frames,omitempty"`
	*/
	/* INHERITED: Models
	Models []*ModelSchemaBase `json:"models,omitempty"`
	*/
	/* INHERITED: Compatible frames
	CompatibleFrames []*FrameV3 `json:"compatible_frames,omitempty"`
	*/
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewModelsV3() *ModelsV3 {
	return &ModelsV3{
		ModelsBase: &ModelsBase{
			ModelId:              nil,
			Preview:              false,
			FindCompatibleFrames: false,
			Models:               nil,
			CompatibleFrames:     nil,
			RequestSchema: &RequestSchema{
				ExcludeFields: "",
			},
		},
	}
}
