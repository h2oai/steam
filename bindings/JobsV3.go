package bindings

type JobsV3 struct {
	*RequestSchema
	/** Optional Job identifier */
	JobId *JobKeyV3 `json:"job_id"`

	/** jobs */
	Jobs []*JobV3 `json:"jobs"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func newJobsV3() *JobsV3 {
	return &JobsV3{
		JobId: nil,
		Jobs:  nil,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
