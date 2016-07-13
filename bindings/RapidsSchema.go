package bindings

type RapidsSchema struct {
	*Schema
	/** An Abstract Syntax Tree. */
	Ast string `json:"ast"`
	/** Key name to assign Frame results */
	Id string `json:"id,omitempty"`
	/** Session key */
	SessionId string `json:"session_id,omitempty"`
}

func NewRapidsSchema() *RapidsSchema {
	return &RapidsSchema{
		Ast:       "",
		Id:        "",
		SessionId: "",
		Schema:    &Schema{},
	}
}
