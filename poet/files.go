package poet

import (
	"fmt"
)

// FileSpec represents a .go source file
type FileSpec struct {
	Comment                string
	Package                string      // Package that the file belongs to
	InitializationPackages []Import    // InitializationPackages include any imports that need to be included for their side effects
	Init                   *FuncSpec   // Init is a single function to be outputted before all CodeBlocks
	CodeBlocks             []CodeBlock // CodeBlocks are appended and when outputted separated by a newline
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
	w := newCodeWriter()

	f.writeHeader(w)
	f.writeImports(w)
	f.writeInitFunc(w)
	f.writeCodeBlocks(w)

	return w.String()
}

// InitializationPackage appends an initialization package for its side effects
func (f *FileSpec) InitializationPackage(imp Import) *FileSpec {
	if imp.GetPackage() != "" {
		f.InitializationPackages = append(f.InitializationPackages, &ImportSpec{
			Package: imp.GetPackage(),
			Alias:   "_",
		})
	}
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

func (f *FileSpec) writeHeader(w *codeWriter) {
	if f.Comment != "" {
		w.WriteStatement(statement{
			Format:    "// $L",
			Arguments: []interface{}{f.Comment},
		})
	}
	w.WriteStatement(statement{
		Format:    "package $L\n",
		Arguments: []interface{}{f.Package},
	})
}

func (f *FileSpec) writeImports(w *codeWriter) {
	imports := collectImports(f.InitializationPackages, f.CodeBlocks)
	if len(imports) == 0 {
		return
	}

	w.WriteStatement(statement{
		Format:      "import (",
		AfterIndent: 1,
	})
	for _, i := range imports {
		var prefix string
		if i.GetAlias() != "" {
			prefix = i.GetAlias() + " "
		}
		w.WriteStatement(statement{
			Format:    "$L$S",
			Arguments: []interface{}{prefix, i.GetPackage()},
		})
	}
	w.WriteStatement(statement{
		Format:       ")\n",
		BeforeIndent: -1,
	})
}

func (f *FileSpec) writeInitFunc(w *codeWriter) {
	if f.Init != nil {
		w.WriteStatement(statement{Format: f.Init.String()})
	}
}

func (f *FileSpec) writeCodeBlocks(w *codeWriter) {
	for _, blk := range f.CodeBlocks {
		w.WriteStatement(statement{Format: blk.String()})
	}
}

func collectImports(initPackages []Import, codeBlocks []CodeBlock) []Import {
	var packages []Import
	packages = append(packages, initPackages...)
	// Collect the imports from each code block
	for _, blk := range codeBlocks {
		for _, i := range blk.GetImports() {
			// external packages only
			if i.GetPackage() != "" {
				packages = append(packages, i)
			}
		}
	}
	return packages
}
