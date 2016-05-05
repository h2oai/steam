package bindings

type FrameBase struct {
	*Schema
	/** Frame ID */
	FrameId *FrameKeyV3 `json:"frame_id"`

	/** Total data size in bytes */
	ByteSize int64 `json:"byte_size"`

	/** Is this Frame raw unparsed data? */
	IsText bool `json:"is_text"`
}

func newFrameBase() *FrameBase {
	return &FrameBase{
		FrameId:  nil,
		ByteSize: 0,
		IsText:   false,
	}
}
