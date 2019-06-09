package meteomatics

type ParameterString string

type ParameterStringer interface {
	ParameterString() ParameterString
}

func (s ParameterString) ParameterString() ParameterString {
	return s
}