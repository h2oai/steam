package bindings

type SharedTreeV3 struct {
	*ModelBuilderSchema
	/* INHERITED: Model builder parameters.
	   Parameters []ModelParameterSchemaV3
	*/
	/* INHERITED: The algo name for this ModelBuilder.
	Algo string `json:"algo,omitempty"`
	*/
	/* INHERITED: The pretty algo name for this ModelBuilder (e.g., Generalized Linear Model, rather than GLM).
	AlgoFullName string `json:"algo_full_name,omitempty"`
	*/
	/* INHERITED: Model categories this ModelBuilder can build.
	CanBuild []ModelCategory `json:"can_build,omitempty"`
	*/
	/* INHERITED: Should the builder always be visible, be marked as beta, or only visible if the user starts up with the experimental flag?
	Visibility ModelBuilderBuilderVisibility `json:"visibility,omitempty"`
	*/
	/* INHERITED: Job Key
	Job *JobV3 `json:"job,omitempty"`
	*/
	/* INHERITED: Parameter validation messages
	Messages []*ValidationMessageV3 `json:"messages,omitempty"`
	*/
	/* INHERITED: Count of parameter validation errors
	ErrorCount int32 `json:"error_count,omitempty"`
	*/
	/* INHERITED: HTTP status to return for this build.
	HttpStatus int32 `json:"__http_status,omitempty"`
	*/
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewSharedTreeV3() *SharedTreeV3 {
	return &SharedTreeV3{
		ModelBuilderSchema: &ModelBuilderSchema{
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
		},
	}
}
