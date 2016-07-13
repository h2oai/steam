package bindings

type FramesBase struct {
	*RequestSchema
	/** Name of Frame of interest */
	FrameId *FrameKeyV3 `json:"frame_id,omitempty"`
	/** Name of column of interest */
	Column string `json:"column,omitempty"`
	/** Row offset to return */
	RowOffset int64 `json:"row_offset,omitempty"`
	/** Number of rows to return */
	RowCount int32 `json:"row_count,omitempty"`
	/** Column offset to return */
	ColumnOffset int32 `json:"column_offset,omitempty"`
	/** Number of columns to return */
	ColumnCount int32 `json:"column_count,omitempty"`
	/** Find and return compatible models? */
	FindCompatibleModels bool `json:"find_compatible_models,omitempty"`
	/** File output path */
	Path string `json:"path,omitempty"`
	/** Overwrite existing file */
	Force bool `json:"force,omitempty"`
	/** Job for export file */
	Job *JobV3 `json:"job,omitempty"`
	/** Frames */
	Frames []*FrameBase `json:"frames,omitempty"`
	/** Compatible models */
	CompatibleModels []*ModelSchema `json:"compatible_models,omitempty"`
	/** Domains */
	Domain [][]string `json:"domain,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewFramesBase() *FramesBase {
	return &FramesBase{
		FrameId:              nil,
		Column:               "",
		RowOffset:            0,
		RowCount:             0,
		ColumnOffset:         0,
		ColumnCount:          0,
		FindCompatibleModels: false,
		Path:                 "",
		Force:                false,
		Job:                  nil,
		Frames:               nil,
		CompatibleModels:     nil,
		Domain:               nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
