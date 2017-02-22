package poet

var _ CodeBlock = (*TypeAliasSpec)(nil)
var _ TypeReference = (*TypeAliasSpec)(nil)

// TypeAliasSpec represents a type alias. *AliasSpec implements CodeBlock and TypeReference.
type TypeAliasSpec struct {
	Name           string
	UnderlyingType TypeReference
	Comment        string
}

// NewTypeAliasSpec returns a new spec representing a type alias.
func NewTypeAliasSpec(name string, typeRef TypeReference) *TypeAliasSpec {
	return &TypeAliasSpec{
		Name:           name,
		UnderlyingType: typeRef,
	}
}

// AliasComment adds a comment to a type alias.
func (a *TypeAliasSpec) AliasComment(comment string) *TypeAliasSpec {
	a.Comment = comment
	return a
}

// GetName returns the alias for this Type Alias.
func (a *TypeAliasSpec) GetName() string {
	return a.Name
}

// GetImports returns a slice of imports that the aliased type requires.
func (a *TypeAliasSpec) GetImports() []Import {
	return a.UnderlyingType.GetImports()
}

func (a *TypeAliasSpec) String() string {
	writer := newCodeWriter()
	if a.Comment != "" {
		writer.WriteComment(a.Comment)
	}

	writer.WriteStatement(statement{
		Format:    "type $T $T",
		Arguments: []interface{}{a, a.UnderlyingType},
	})
	return writer.String()
}
