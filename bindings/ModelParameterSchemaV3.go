package bindings

type ModelParameterSchemaV3 struct {
	*Schema
	/** name in the JSON, e.g. "lambda" */
	Name string `json:"name,omitempty"`
	/** label in the UI, e.g. "lambda" */
	Label string `json:"label,omitempty"`
	/** help for the UI, e.g. "regularization multiplier, typically used for foo bar baz etc." */
	Help string `json:"help,omitempty"`
	/** the field is required */
	Required bool `json:"required,omitempty"`
	/** Java type, e.g. "double" */
	Type string `json:"type,omitempty"`
	/** default value, e.g. 1 */
	DefaultValue interface{} // Polymorphic class `json:"default_value,omitempty"`
	/** actual value as set by the user and / or modified by the ModelBuilder, e.g., 10 */
	ActualValue interface{} // Polymorphic class `json:"actual_value,omitempty"`
	/** the importance of the parameter, used by the UI, e.g. "critical", "extended" or "expert" */
	Level string `json:"level,omitempty"`
	/** list of valid values for use by the front-end */
	Values []string `json:"values,omitempty"`
	/** For Vec-type fields this is the set of other Vec-type fields which must contain mutually exclusive values; for example, for a SupervisedModel the response_column must be mutually exclusive with the weights_column */
	IsMemberOfFrames []string `json:"is_member_of_frames,omitempty"`
	/** For Vec-type fields this is the set of Frame-type fields which must contain the named column; for example, for a SupervisedModel the response_column must be in both the training_frame and (if it's set) the validation_frame */
	IsMutuallyExclusiveWith []string `json:"is_mutually_exclusive_with,omitempty"`
	/** Parameter can be used in grid call */
	Gridable bool `json:"gridable,omitempty"`
}

func NewModelParameterSchemaV3() *ModelParameterSchemaV3 {
	return &ModelParameterSchemaV3{
		Name:                    "",
		Label:                   "",
		Help:                    "",
		Required:                false,
		Type:                    "",
		DefaultValue:            nil,
		ActualValue:             nil,
		Level:                   "",
		Values:                  nil,
		IsMemberOfFrames:        nil,
		IsMutuallyExclusiveWith: nil,
		Gridable:                false,
		Schema:                  &Schema{},
	}
}
