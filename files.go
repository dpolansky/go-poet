package gopoet

import "bytes"

type FileSpec struct {
	Package                string
	InitializationPackages []ImportSpec
	Init                   CodeBlock
	CodeBlocks             []CodeBlock
}

func NewFileSpec(pkg string) *FileSpec {
	return &FileSpec{
		Package:                pkg,
		InitializationPackages: []ImportSpec{},
		CodeBlocks:             []CodeBlock{},
	}
}

func (f *FileSpec) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("package " + f.Package + "\n\n")
	seen := map[string]struct{}{}

	var packages []ImportSpec
	for _, blk := range f.CodeBlocks {
		packages = append(packages, blk.Packages()...)
	}

	if len(f.InitializationPackages) > 0 || len(packages) > 0 {
		buffer.WriteString("import (\n")

		for _, spec := range f.InitializationPackages {
			if _, found := seen[spec.GetPackage()]; !found && spec.GetPackage() != "" {
				buffer.WriteString("\t")
				if spec.GetPackageAlias() != "" {
					buffer.WriteString(spec.GetPackageAlias())
					buffer.WriteString(" ")
				}
				buffer.WriteString("\"" + spec.GetPackage() + "\"\n")
				seen[spec.GetPackage()] = struct{}{}

			}
		}

		for _, spec := range packages {
			if _, found := seen[spec.GetPackage()]; !found && spec.GetPackage() != "" {
				buffer.WriteString("\t")
				if spec.GetPackageAlias() != "" {
					buffer.WriteString(spec.GetPackageAlias())
					buffer.WriteString(" ")
				}
				buffer.WriteString("\"" + spec.GetPackage() + "\"\n")
				seen[spec.GetPackage()] = struct{}{}
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

func (f *FileSpec) Packages() []ImportSpec {
	return []ImportSpec{}
}

func (f *FileSpec) InitializationPackage(spec ImportSpec) *FileSpec {
	f.InitializationPackages = append(f.InitializationPackages, spec)
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
