package gopoet

import "bytes"

type ImportSpec struct {
	Package   string
	Alias     string
	Qualified bool
}

func (i *ImportSpec) GetQualifier() string {
	if i == nil {
		return ""
	}

	result := bytes.Buffer{}

	if i.Qualified {
		if i.Alias != "" {
			result.WriteString(i.Alias)
		} else {
			result.WriteString(i.Package)
		}
		result.WriteString(".")
	}

	return result.String()
}

var _ Import = (*ImportSpec)(nil)

func (i *ImportSpec) GetAlias() string {
	return i.Alias
}

func (i *ImportSpec) GetPackage() string {
	return i.Package
}
