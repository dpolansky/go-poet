package poet

var _ CodeBlock = (*AliasSpec)(nil)
var _ TypeReference = (*AliasSpec)(nil)

// AliasSpec represents a type alias. *AliasSpec implements CodeBlock and TypeReference.
type AliasSpec struct {
	Name           string
	UnderlyingType TypeReference
	Comment        string
}

// NewAliasSpec returns a new spec representing a type alias.
func NewAliasSpec(name string, typeRef TypeReference) *AliasSpec {
	return &AliasSpec{
		Name:           name,
		UnderlyingType: typeRef,
	}
}

func (a *AliasSpec) AliasComment(comment string) *AliasSpec {
	a.Comment = comment
	return a
}

func (a *AliasSpec) GetName() string {
	return a.Name
}

func (a *AliasSpec) GetImports() []Import {
	return a.UnderlyingType.GetImports()
}

func (a *AliasSpec) String() string {
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
