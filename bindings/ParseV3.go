package bindings

type ParseV3 struct {
	*RequestSchema
	/** Final frame name */
	DestinationFrame *FrameKeyV3 `json:"destination_frame"`
	/** Source frames */
	SourceFrames []*FrameKeyV3 `json:"source_frames"`
	/** Parser type */
	ParseType ApiParseTypeValuesProvider `json:"parse_type,omitempty"`
	/** Field separator */
	Separator int8 `json:"separator,omitempty"`
	/** Single Quotes */
	SingleQuotes bool `json:"single_quotes,omitempty"`
	/** Check header: 0 means guess, +1 means 1st line is header not data, -1 means 1st line is data not header */
	CheckHeader int32 `json:"check_header,omitempty"`
	/** Number of columns */
	NumberColumns int32 `json:"number_columns,omitempty"`
	/** Column names */
	ColumnNames []string `json:"column_names,omitempty"`
	/** Value types for columns */
	ColumnTypes []string `json:"column_types,omitempty"`
	/** Domains for categorical columns */
	Domains [][]string `json:"domains,omitempty"`
	/** NA strings for columns */
	NaStrings [][]string `json:"na_strings,omitempty"`
	/** Size of individual parse tasks */
	ChunkSize int32 `json:"chunk_size,omitempty"`
	/** Delete input key after parse */
	DeleteOnDone bool `json:"delete_on_done,omitempty"`
	/** Block until the parse completes (as opposed to returning early and requiring polling */
	Blocking bool `json:"blocking,omitempty"`
	/** Parse job */
	Job *JobV3 `json:"job,omitempty"`
	/** Rows */
	Rows int64 `json:"rows,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewParseV3() *ParseV3 {
	return &ParseV3{
		DestinationFrame: nil,
		SourceFrames:     nil,
		ParseType:        ApiParseTypeValuesProvider_NONE_,
		Separator:        0,
		SingleQuotes:     false,
		CheckHeader:      0,
		NumberColumns:    0,
		ColumnNames:      nil,
		ColumnTypes:      nil,
		Domains:          nil,
		NaStrings:        nil,
		ChunkSize:        0,
		DeleteOnDone:     false,
		Blocking:         false,
		Job:              nil,
		Rows:             0,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}

func (p *ParseV3) FromParseSetup(parseSetup ParseSetupV3) {
	newFr := NewFrameKeyV3()
	newFr.Name = parseSetup.DestinationFrame
	p.DestinationFrame = newFr

	p.CheckHeader = parseSetup.CheckHeader
	p.ChunkSize = parseSetup.ChunkSize
	p.ColumnNames = parseSetup.ColumnNames
	p.ColumnTypes = parseSetup.ColumnTypes
	p.ExcludeFields = parseSetup.ExcludeFields
	p.NaStrings = parseSetup.NaStrings
	p.NumberColumns = parseSetup.NumberColumns
	p.ParseType = parseSetup.ParseType
	p.Separator = parseSetup.Separator
	p.SingleQuotes = parseSetup.SingleQuotes
	p.SourceFrames = parseSetup.SourceFrames
}
