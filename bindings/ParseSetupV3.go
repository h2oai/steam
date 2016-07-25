package bindings

type ParseSetupV3 struct {
	*RequestSchema
	/** Source frames */
	SourceFrames []*FrameKeyV3 `json:"source_frames"`
	/** Parser type */
	ParseType ApiParseTypeValuesProvider `json:"parse_type,omitempty"`
	/** Field separator */
	Separator int8 `json:"separator,omitempty"`
	/** Single quotes */
	SingleQuotes bool `json:"single_quotes,omitempty"`
	/** Check header: 0 means guess, +1 means 1st line is header not data, -1 means 1st line is data not header */
	CheckHeader int32 `json:"check_header,omitempty"`
	/** Column names */
	ColumnNames []string `json:"column_names,omitempty"`
	/** Value types for columns */
	ColumnTypes []string `json:"column_types,omitempty"`
	/** NA strings for columns */
	NaStrings [][]string `json:"na_strings,omitempty"`
	/** Regex for names of columns to return */
	ColumnNameFilter string `json:"column_name_filter,omitempty"`
	/** Column offset to return */
	ColumnOffset int32 `json:"column_offset,omitempty"`
	/** Number of columns to return */
	ColumnCount int32 `json:"column_count,omitempty"`
	/** Suggested name */
	DestinationFrame string `json:"destination_frame,omitempty"`
	/** Number of header lines found */
	HeaderLines int64 `json:"header_lines,omitempty"`
	/** Number of columns */
	NumberColumns int32 `json:"number_columns,omitempty"`
	/** Sample data */
	Data [][]string `json:"data,omitempty"`
	/** Warnings */
	Warnings []string `json:"warnings,omitempty"`
	/** Size of individual parse tasks */
	ChunkSize int32 `json:"chunk_size,omitempty"`
	/** Total number of columns we would return with no column pagination */
	TotalFilteredColumnCount int32 `json:"total_filtered_column_count,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewParseSetupV3() *ParseSetupV3 {
	return &ParseSetupV3{
		SourceFrames:             nil,
		ParseType:                ApiParseTypeValuesProviderGUESS,
		Separator:                0,
		SingleQuotes:             false,
		CheckHeader:              0,
		ColumnNames:              nil,
		ColumnTypes:              nil,
		NaStrings:                nil,
		ColumnNameFilter:         "",
		ColumnOffset:             0,
		ColumnCount:              0,
		DestinationFrame:         "",
		HeaderLines:              0,
		NumberColumns:            0,
		Data:                     nil,
		Warnings:                 nil,
		ChunkSize:                4194304,
		TotalFilteredColumnCount: 0,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
