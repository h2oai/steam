package bindings

type JobV3 struct {
	*RequestSchema
	/** Job Key */
	Key *JobKeyV3 `json:"key"`

	/** Job description */
	Description string `json:"description"`

	/** job status */
	Status string `json:"status"`

	/** progress, from 0 to 1 */
	Progress float32 `json:"progress"`

	/** current progress status description */
	ProgressMsg string `json:"progress_msg"`

	/** Start time */
	StartTime int64 `json:"start_time"`

	/** Runtime in milliseconds */
	Msec int64 `json:"msec"`

	/** destination key */
	Dest *KeyV3 `json:"dest"`

	/** exception */
	Warnings []string `json:"warnings"`

	/** exception */
	Exception string `json:"exception"`

	/** stacktrace */
	Stacktrace string `json:"stacktrace"`

	/** ready for view */
	ReadyForView bool `json:"ready_for_view"`

	/* INHERITED: Comma-separated list of JSON field paths to exclude from the result, used like: "/3/Frames?_exclude_fields=frames/frame_id/URL,__meta"
	ExcludeFields string: "" `json:"_exclude_fields"`
	*/
}

func NewJobV3() *JobV3 {
	return &JobV3{
		Key:          nil,
		Description:  "",
		Status:       "",
		Progress:     0.0,
		ProgressMsg:  "",
		StartTime:    0,
		Msec:         0,
		Dest:         nil,
		Warnings:     nil,
		Exception:    "",
		Stacktrace:   "",
		ReadyForView: false,
		RequestSchema: &RequestSchema{
			ExcludeFields: "",
		},
	}
}
