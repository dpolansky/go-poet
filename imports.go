package gopoet

import "bytes"

// ImportSpec implements Import to represent an imported go package
type ImportSpec struct {
	Package   string
	Alias     string
	Qualified bool
}

// getQualifier returns the fully qualified package (e.g. bytes.) for use in a qualified
// declared type
func (i *ImportSpec) getQualifier() string {
	if i == nil || !i.Qualified {
		return ""
	}

	result := bytes.Buffer{}

	if i.Alias != "" {
		result.WriteString(i.Alias)
	} else {
		result.WriteString(i.Package)
	}
	result.WriteString(".")

	return result.String()
}

var _ Import = (*ImportSpec)(nil)

// GetAlias returns the alias associated with the package
func (i *ImportSpec) GetAlias() string {
	if i == nil {
		return ""
	}

	return i.Alias
}

// GetPackage returns the package
func (i *ImportSpec) GetPackage() string {
	if i == nil {
		return ""
	}

	return i.Package
}
