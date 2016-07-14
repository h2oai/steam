package bindings

type ValidationMessageBase struct {
	*Schema
	/** Type of validation message (ERROR, WARN, INFO, HIDE) */
	MessageType string `json:"message_type,omitempty"`
	/** Field to which the message applies */
	FieldName string `json:"field_name,omitempty"`
	/** Message text */
	Message string `json:"message,omitempty"`
}

func NewValidationMessageBase() *ValidationMessageBase {
	return &ValidationMessageBase{
		MessageType: "",
		FieldName:   "",
		Message:     "",
		Schema:      &Schema{},
	}
}
