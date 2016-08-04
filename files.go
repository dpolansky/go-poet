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
	buffer.WriteString("package " + f.Package + "\n")

	for _, codeBlk := range f.CodeBlocks {
		buffer.WriteString(codeBlk.String())
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
