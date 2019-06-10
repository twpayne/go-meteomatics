package meteomatics

import "strings"

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

// A ParameterSlice is a slice of ParameterStringers.
type ParameterSlice []ParameterStringer

// ParameterString returns s as a ParameterString.
func (s ParameterSlice) ParameterString() ParameterString {
	ss := make([]string, len(s))
	for i, ps := range s {
		ss[i] = string(ps.ParameterString())
	}
	return ParameterString(strings.Join(ss, ","))
}
