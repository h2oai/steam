package bindings

type ColSpecifierV3 struct {
	*Schema
	/** Name of the column */
	ColumnName string `json:"column_name,omitempty"`
	/** List of fields which specify columns that must contain this column */
	IsMemberOfFrames []string `json:"is_member_of_frames,omitempty"`
}

func NewColSpecifierV3() *ColSpecifierV3 {
	return &ColSpecifierV3{
		ColumnName:       "",
		IsMemberOfFrames: nil,
		Schema:           &Schema{},
	}
}
