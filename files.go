package gopoet

import "bytes"

type FileSpec struct {
	Package                string
	InitializationPackages []Import
	Init                   CodeBlock
	CodeBlocks             []CodeBlock
}

func NewFileSpec(pkg string) *FileSpec {
	return &FileSpec{
		Package:                pkg,
		InitializationPackages: []Import{},
		CodeBlocks:             []CodeBlock{},
	}
}

func (f *FileSpec) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("package " + f.Package + "\n\n")
	seen := map[string]struct{}{}

	var packages []Import
	for _, blk := range f.CodeBlocks {
		packages = append(packages, blk.GetImports()...)
	}

	if len(f.InitializationPackages) > 0 || len(packages) > 0 {
		buffer.WriteString("import (\n")

		for _, imp := range f.InitializationPackages {
			if _, found := seen[imp.GetPackage()]; !found && imp.GetPackage() != "" {
				buffer.WriteString("\t")
				if imp.GetAlias() != "" {
					buffer.WriteString(imp.GetAlias())
					buffer.WriteString(" ")
				}
				buffer.WriteString("\"" + imp.GetPackage() + "\"\n")
				seen[imp.GetPackage()] = struct{}{}
			}
		}

		for _, imp := range packages {
			if _, found := seen[imp.GetPackage()]; !found && imp.GetPackage() != "" {
				buffer.WriteString("\t")
				if imp.GetAlias() != "" {
					buffer.WriteString(imp.GetAlias())
					buffer.WriteString(" ")
				}
				buffer.WriteString("\"" + imp.GetPackage() + "\"\n")
				seen[imp.GetPackage()] = struct{}{}
			}
		}

		buffer.WriteString(")\n\n")
	}

	for _, codeBlk := range f.CodeBlocks {
		buffer.WriteString(codeBlk.String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (f *FileSpec) Packages() []TypeReference {
	return []TypeReference{}
}

func (f *FileSpec) InitializationPackage(imp Import) *FileSpec {
	f.InitializationPackages = append(f.InitializationPackages, imp)
	return f
}

func (f *FileSpec) CodeBlock(blk CodeBlock) *FileSpec {
	f.CodeBlocks = append(f.CodeBlocks, blk)
	return f
}

func (f *FileSpec) InitFunction(blk CodeBlock) *FileSpec {
	f.Init = blk
	return f
}
