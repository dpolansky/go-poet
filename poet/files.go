package poet

import (
	"bytes"
	"fmt"
)

// FileSpec represents a .go source file
type FileSpec struct {
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
	var buffer bytes.Buffer
	didStartImportBlock := false

	buffer.WriteString("package " + f.Package + "\n\n")
	seen := map[string]struct{}{}

	var packages []Import
	for _, blk := range f.CodeBlocks {
		packages = append(packages, blk.GetImports()...)
	}

	for _, imp := range f.InitializationPackages {
		if _, found := seen[imp.GetPackage()]; !found && imp.GetPackage() != "" {
			if !didStartImportBlock {
				buffer.WriteString("import (\n")
				didStartImportBlock = true
			}

			buffer.WriteString("\t_ ")
			buffer.WriteString("\"" + imp.GetPackage() + "\"\n")
			seen[imp.GetPackage()] = struct{}{}
		}
	}

	for _, imp := range packages {
		if _, found := seen[imp.GetPackage()]; !found && imp.GetPackage() != "" {
			if !didStartImportBlock {
				buffer.WriteString("import (\n")
				didStartImportBlock = true
			}

			buffer.WriteString("\t")
			if imp.GetAlias() != "" {
				buffer.WriteString(imp.GetAlias())
				buffer.WriteString(" ")
			}
			buffer.WriteString("\"" + imp.GetPackage() + "\"\n")
			seen[imp.GetPackage()] = struct{}{}
		}
	}

	if didStartImportBlock {
		buffer.WriteString(")\n\n")
	}

	// create a new array with codeBlocks
	var codeBlocks []CodeBlock
	if f.Init != nil {
		codeBlocks = append([]CodeBlock{f.Init}, f.CodeBlocks...)
	} else {
		codeBlocks = append([]CodeBlock(nil), f.CodeBlocks...)
	}

	for _, codeBlk := range codeBlocks {
		buffer.WriteString(codeBlk.String())
		buffer.WriteString("\n")
	}

	return buffer.String()
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
