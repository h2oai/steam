package bindings

type InitIDV3 struct {
	*RequestSchema
	/** Session ID */
	SessionKey string `json:"session_key,omitempty"`
	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string `json:"_exclude_fields,omitempty"`
	*/
}

func NewInitIDV3() *InitIDV3 {
	return &InitIDV3{
		SessionKey: "",
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
