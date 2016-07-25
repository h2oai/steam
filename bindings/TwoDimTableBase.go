package bindings

type TwoDimTableBase struct {
	*Schema
	/** Table Name */
	Name string `json:"name"`

	/** Table Description */
	Description string `json:"description"`

	/** Column Specification */
	Columns []*ColumnSpecsBase `json:"columns"`

	/** Number of Rows */
	Rowcount int32 `json:"rowcount"`

	/** Table Data (col-major) */
	Data [][]Polymorphic `json:"data"`
}

func NewTwoDimTableBase() *TwoDimTableBase {
	return &TwoDimTableBase{
		Name:        "",
		Description: "",
		Columns:     nil,
		Rowcount:    0,
		Data:        nil,
	}
}
