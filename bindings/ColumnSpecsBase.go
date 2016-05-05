package bindings

type ColumnSpecsBase struct {
	*Schema
	/** Column Name */
	Name string `json:"name"`

	/** Column Type */
	Type string `json:"type"`

	/** Column Format (printf) */
	Format string `json:"format"`

	/** Column Description */
	Description string `json:"description"`
}

func newColumnSpecsBase() *ColumnSpecsBase {
	return &ColumnSpecsBase{
		Name:        "",
		Type:        "",
		Format:      "",
		Description: "",
	}
}
