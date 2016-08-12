package poet

import (
	"bytes"
	"fmt"
)

// FileSpec represents a .go source file
type FileSpec struct {
	Comment                string
	Package                string      // Package that the file belongs to
	InitializationPackages []Import    // InitializationPackages include any imports that need to be included for their side effects
	Init                   *FuncSpec   // Init is a single function to be outputted before all CodeBlocks
	CodeBlocks             []CodeBlock // CodeBlocks are appeneded and when outputted separated by a newline
}

// NewFileSpec constructs a new FileSpec with the given package name
func NewFileSpec(pkg string) *FileSpec {
	return &FileSpec{
		Package:                pkg,
		InitializationPackages: []Import{},
		CodeBlocks:             []CodeBlock{},
	}
}

// String produces the final go file string
func (f *FileSpec) String() string {
	seen := map[string]struct{}{}
	b := &bytes.Buffer{}

	if f.Comment != "" {
		b.WriteString("// " + f.Comment + "\n")
	}
	b.WriteString("package " + f.Package + "\n\n")

	// Collect the imports from each code block
	var packages []Import
	for _, blk := range f.CodeBlocks {
		packages = append(packages, blk.GetImports()...)
	}

	// if there are packages to write
	if hasExternalImports(packages) || hasExternalImports(f.InitializationPackages) {
		b.WriteString("import (\n")

		// initialization packages should have an _ for an alias if no alias is specified
		writeImports(b, f.InitializationPackages, seen, true)
		writeImports(b, packages, seen, false)

		b.WriteString(")\n\n")
	}

	if f.Init != nil {
		b.WriteString(f.Init.String() + "\n")
	}

	for _, blk := range f.CodeBlocks {
		b.WriteString(blk.String() + "\n")
	}

	return b.String()
}

// Returns whether a slice of imports has external or non-primitive imports
func hasExternalImports(imports []Import) bool {
	for _, i := range imports {
		if i.GetPackage() != "" {
			return true
		}
	}

	return false
}

// Writes imports to a buffer, using the specified fallbackAlias if an import does not have a specified
// alias (in the case of initialization imports)
func writeImports(b *bytes.Buffer, imports []Import, seen map[string]struct{}, isInitialization bool) {
	for _, i := range imports {
		// skip any non-external packages
		if i.GetPackage() == "" {
			continue
		}

		b.WriteString("\t")

		if isInitialization {
			// initialization packages should be prefixed with an _
			b.WriteString("_ ")
		} else if i.GetAlias() != "" {
			b.WriteString(i.GetAlias() + " ")
		}

		b.WriteString("\"" + i.GetPackage() + "\"\n")
		seen[i.GetPackage()] = struct{}{}
	}
}

// InitializationPackage appends an initialization package for its side effects
func (f *FileSpec) InitializationPackage(imp Import) *FileSpec {
	f.InitializationPackages = append(f.InitializationPackages, imp)
	return f
}

// CodeBlock adds a code block to the file
func (f *FileSpec) CodeBlock(blk CodeBlock) *FileSpec {
	f.CodeBlocks = append(f.CodeBlocks, blk)
	return f
}

// InitFunction assign an init function
func (f *FileSpec) InitFunction(blk *FuncSpec) *FileSpec {
	if blk.Name != "init" {
		panic(fmt.Sprintf("the init function must be named 'init' (got '%s')", f.Init.Name))
	}

	f.Init = blk
	return f
}

// GlobalVariable adds a global variable to the file with the given name, type reference, format string
// for the value of the variable, and arguments for the format string.
func (f *FileSpec) GlobalVariable(name string, typ TypeReference, format string, args ...interface{}) *FileSpec {
	v := &Variable{
		Identifier: Identifier{
			Name: name,
			Type: typ,
		},
		Constant: false,
		Format:   format,
		Args:     args,
	}
	f.CodeBlocks = append(f.CodeBlocks, v)
	return f
}

// GlobalConstant adds a global constant to the file with the given name, type reference, format string
// for the value of the constant, and arguments for the format string.
func (f *FileSpec) GlobalConstant(name string, typ TypeReference, format string, args ...interface{}) *FileSpec {
	v := &Variable{
		Identifier: Identifier{
			Name: name,
			Type: typ,
		},
		Constant: true,
		Format:   format,
		Args:     args,
	}
	f.CodeBlocks = append(f.CodeBlocks, v)
	return f
}

// VariableGrouping adds a variable grouping to the file, which can have variables appended to it
// that will be formatted in groups of variables and constants.
func (f *FileSpec) VariableGrouping() *VariableGrouping {
	v := &VariableGrouping{}
	f.CodeBlocks = append(f.CodeBlocks, v)
	return v
}

// FileComment sets the file's comment to the given input string.
func (f *FileSpec) FileComment(comment string) *FileSpec {
	f.Comment = comment
	return f
}
