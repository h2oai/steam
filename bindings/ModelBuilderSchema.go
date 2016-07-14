package bindings

type ModelBuilderSchema struct {
	*RequestSchema
	/** Model builder parameters. */
	Parameters []ModelParameterSchemaV3
	/** The algo name for this ModelBuilder. */
	Algo string `json:"algo,omitempty"`
	/** The pretty algo name for this ModelBuilder (e.g., Generalized Linear Model, rather than GLM). */
	AlgoFullName string `json:"algo_full_name,omitempty"`
	/** Model categories this ModelBuilder can build. */
	CanBuild []ModelCategory `json:"can_build,omitempty"`
	/** Should the builder always be visible, be marked as beta, or only visible if the user starts up with the experimental flag? */
	Visibility ModelBuilderBuilderVisibility `json:"visibility,omitempty"`
	/** Job Key */
	Job *JobV3 `json:"job,omitempty"`
	/** Parameter validation messages */
	Messages []*ValidationMessageBase `json:"messages,omitempty"`
	/** Count of parameter validation errors */
	ErrorCount int32 `json:"error_count,omitempty"`
	/** HTTP status to return for this build. */
	HttpStatus int32 `json:"__http_status,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewModelBuilderSchema() *ModelBuilderSchema {
	return &ModelBuilderSchema{
		Algo:         "",
		AlgoFullName: "",
		CanBuild:     nil,
		Visibility:   ModelBuilderBuilderVisibility_NONE_,
		Job:          nil,
		Messages:     nil,
		ErrorCount:   0,
		HttpStatus:   0,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
