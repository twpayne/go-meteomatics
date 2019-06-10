package meteomatics

// A ParameterString is a string representing a parameter.
type ParameterString string

// A ParameterStringer can be converted to a ParameterString.
type ParameterStringer interface {
	ParameterString() ParameterString
}

// ParameterString returns s as a ParameterString.
func (s ParameterString) ParameterString() ParameterString {
	return s
}
