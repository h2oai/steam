package bindings


type RequestSchema struct{
*Schema
    /** Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta" */
ExcludeFields string `json:"_exclude_fields"`
}

func newRequestSchema() *RequestSchema{
return &RequestSchema{
ExcludeFields: "",
}
}
