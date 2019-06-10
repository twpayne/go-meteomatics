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

// A Parameter is a parameter.
type Parameter struct {
	Name     ParameterName
	Level    LevelStringer
	Interval IntervalStringer
	Units    Units
}

// ParameterString returns p as a ParameterString.
func (p Parameter) ParameterString() ParameterString {
	ps := string(p.Name)
	if p.Level != nil {
		if level := p.Level.LevelString(); level != "" {
			ps += "_" + string(level)
		}
	}
	if p.Interval != nil {
		if interval := p.Interval.IntervalString(); interval != "" {
			ps += "_" + string(interval)
		}
	}
	ps += ":" + string(p.Units)
	return ParameterString(ps)
}
