package bindings

type TwoDimTableV3 struct {
	*TwoDimTableBase
	/* INHERITED: Table Name
	Name string: "" `json:"name"`
	*/

	/* INHERITED: Table Description
	Description string: "" `json:"description"`
	*/

	/* INHERITED: Column Specification
	Columns []*ColumnSpecsBase: nil `json:"columns"`
	*/

	/* INHERITED: Number of Rows
	Rowcount int32: 0 `json:"rowcount"`
	*/

	/* INHERITED: Table Data (col-major)
	Data [][]Polymorphic: nil `json:"data"`
	*/
}

func newTwoDimTableV3() *TwoDimTableV3 {
	return &TwoDimTableV3{
		TwoDimTableBase: &TwoDimTableBase{
			Name:        "",
			Description: "",
			Columns:     nil,
			Rowcount:    0,
			Data:        nil,
		},
	}
}
