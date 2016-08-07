package main

import (
	"bytes"
	"fmt"

	"gitlab.com/kdeenanauth/gopoet"
)

type blah struct {
}

func main() {
	fileSpec := gopoet.NewFileSpec("main")
	funcSpec := gopoet.NewFuncSpec("blah")
	funcSpec.Parameter("a", gopoet.ImportFromInstance(&bytes.Buffer{}))
	NewParameter("a", isPtr)
	fileSpec.CodeBlock(funcSpec)
	fmt.Println(fileSpec)
}

func mainOld() {
	fmtImport := gopoet.Import{
		Package: "fmt",
		Name:    "Println",
	}

	sampleStruct := gopoet.NewStructSpec("A")
	blahStruct := gopoet.ImportFromInstance(bytes.Buffer{})
	typeReference := gopoet.TypeReferenceFromInstance(&bytes.Buffer{})

	blahStruct.Unqualified = true

	sampleMethodSpec := gopoet.NewMethodSpec("sampleMethod", "a", false, sampleStruct)
	blahSpec := gopoet.NewMethodSpec("blahMethod", "b", true, blahStruct)

	sampleStruct.MethodAndAttach("blahMethod", "b", false)

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
