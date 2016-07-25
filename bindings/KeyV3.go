package bindings

type KeyV3 struct {
	*Schema
	/** Name (string representation) for this Key. */
	Name string `json:"name"`

	/** Name (string representation) for the type of Keyed this Key points to. */
	Type string `json:"type"`

	/** URL for the resource that this Key points to, if one exists. */
	Url string `json:"URL"`
}

func NewKeyV3() *KeyV3 {
	return &KeyV3{
		Name: "",
		Type: "",
		Url:  "",
	}
}
