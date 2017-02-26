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
		Value:    newStatement(0, 0, format, args...),
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
		Value:    newStatement(0, 0, format, args...),
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
		w.WriteCodeBlock(Comment(f.Comment))
	}
	w.WriteStatement(newStatement(0, 0, "package $L\n", f.Package))
}

func (f *FileSpec) writeImports(w *codeWriter) {
	imports := collectImports(f.InitializationPackages, f.CodeBlocks)
	if len(imports) == 0 {
		return
	}

	w.WriteStatement(newStatement(0, 1, "import ("))
	for _, i := range imports {
		var prefix string
		if i.GetAlias() != "" {
			prefix = i.GetAlias() + " "
		}
		w.WriteStatement(newStatement(0, 0, "$L$S", prefix, i.GetPackage()))
	}
	w.WriteStatement(newStatement(-1, 0, ")\n"))
}

func (f *FileSpec) writeInitFunc(w *codeWriter) {
	if f.Init != nil {
		w.WriteCodeBlock(f.Init)
		w.WriteStatement(Statement{})
	}
}

func (f *FileSpec) writeCodeBlocks(w *codeWriter) {
	for _, blk := range f.CodeBlocks {
		w.WriteCodeBlock(blk)
		w.WriteStatement(Statement{})
	}
}

func collectImports(initPackages []Import, codeBlocks []CodeBlock) []Import {
	// map[Package]map[Alias]Import
	packages := make(map[string]map[string]Import)
	for _, i := range initPackages {
		pkg := i.GetPackage()
		if _, exists := packages[pkg]; !exists {
			packages[pkg] = make(map[string]Import)
		}
		packages[i.GetPackage()][i.GetAlias()] = i
	}
	// Collect the imports from each code block
	for _, blk := range codeBlocks {
		for _, i := range blk.GetImports() {
			pkg := i.GetPackage()
			// external packages only
			if pkg != "" {
				if _, exists := packages[pkg]; !exists {
					packages[pkg] = make(map[string]Import)
				}
				packages[pkg][i.GetAlias()] = i
			}
		}
	}

	var pkgSlice []Import
	for _, aliasMap := range packages {
		for _, i := range aliasMap {
			pkgSlice = append(pkgSlice, i)
		}
	}
	return pkgSlice
}
