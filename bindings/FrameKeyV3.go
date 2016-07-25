package bindings

type FrameKeyV3 struct {
	*KeyV3
	/* INHERITED: Name (string representation) for this Key.
	Name string: "" `json:"name"`
	*/

	/* INHERITED: Name (string representation) for the type of Keyed this Key points to.
	Type string: "" `json:"type"`
	*/

	/* INHERITED: URL for the resource that this Key points to, if one exists.
	Url string: "" `json:"URL"`
	*/
}

func NewFrameKeyV3() *FrameKeyV3 {
	return &FrameKeyV3{
		KeyV3: &KeyV3{
			Name: "",
			Type: "",
			Url:  "",
		},
	}
}
