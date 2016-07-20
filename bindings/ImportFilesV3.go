package bindings

type ImportFilesV3 struct {
	*RequestSchema
	/** path */
	Path string `json:"path"`
	/** files */
	Files []string `json:"files,omitempty"`
	/** names */
	DestinationFrames []string `json:"destination_frames,omitempty"`
	/** fails */
	Fails []string `json:"fails,omitempty"`
	/** dels */
	Dels []string `json:"dels,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewImportFilesV3() *ImportFilesV3 {
	return &ImportFilesV3{
		Path:              "",
		Files:             nil,
		DestinationFrames: nil,
		Fails:             nil,
		Dels:              nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
