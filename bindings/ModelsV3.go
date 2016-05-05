package bindings

type ModelsV3 struct {
	*ModelsBase
	/* INHERITED: Name of Model of interest
	ModelId *ModelKeyV3: null `json:"model_id"`
	*/

	/* INHERITED: Return potentially abridged model suitable for viewing in a browser
	Preview bool: false `json:"preview"`
	*/

	/* INHERITED: Find and return compatible frames?
	FindCompatibleFrames bool: false `json:"find_compatible_frames"`
	*/

	/* INHERITED: Models
	Models []*ModelSchemaBase: nil `json:"models"`
	*/

	/* INHERITED: Compatible frames
	CompatibleFrames []*FrameV3: nil `json:"compatible_frames"`
	*/

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func newModelsV3() *ModelsV3 {
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
