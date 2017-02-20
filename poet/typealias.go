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

func (a *TypeAliasSpec) AliasComment(comment string) *TypeAliasSpec {
	a.Comment = comment
	return a
}

func (a *TypeAliasSpec) GetName() string {
	return a.Name
}

func (a *TypeAliasSpec) GetImports() []Import {
	return a.UnderlyingType.GetImports()
}

func (a *TypeAliasSpec) String() string {
	writer := newCodeWriter()
	if a.Comment != "" {
		writer.WriteStatement(statement{
			Format:    "// $L",
			Arguments: []interface{}{a.Comment},
		})
	}
	writer.WriteStatement(statement{
		Format:    "type $T $T",
		Arguments: []interface{}{a, a.UnderlyingType},
	})
	return writer.String()
}
