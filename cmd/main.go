package main

import (
	"fmt"

	"gitlab.com/kdeenanauth/gopoet"
)

type blah struct {
}

func main() {
	fmtImport := gopoet.Import{
		Package: "fmt",
		Name:    "Println",
	}

	sampleStruct := gopoet.NewStructSpec("A")
	blahStruct := gopoet.ImportFromInstance(blah{})
	blahStruct.Unqualified = true

	sampleMethodSpec := gopoet.NewMethodSpec("sampleMethod", "a", false, sampleStruct)
	blahSpec := gopoet.NewMethodSpec("blahMethod", "b", true, blahStruct)

	sampleStruct.AttachMethod("blahMethod", "b", false)

	mainSpec := gopoet.NewFuncSpec("main")
	mainSpec.Statement("$T($S)", fmtImport, "Calling hello...")

	helloSpec := gopoet.NewFuncSpec("hello")
	helloSpec.BlockStart("for i:= 0; i < 5; i++")
	helloSpec.Statement("$T($L)", fmtImport, "i")
	helloSpec.BlockEnd()

	mainSpec.Statement("$L()", helloSpec.Name)

	fileSpec := gopoet.NewFileSpec("main")
	fileSpec.CodeBlock(sampleStruct)
	fileSpec.CodeBlock(mainSpec)
	fileSpec.CodeBlock(helloSpec)
	fileSpec.CodeBlock(sampleMethodSpec)
	fileSpec.CodeBlock(blahSpec)

	fmt.Println(fileSpec.String())
}
