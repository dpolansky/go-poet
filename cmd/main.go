package main

import (
	"bytes"
	"fmt"

	"gitlab.com/kdeenanauth/gopoet"
)

type blah struct {
}

func main() {
	typeRef := gopoet.TypeReferenceFromInstance(make(chan<- *bytes.Buffer))
	fmt.Println(typeRef.GetName())
}

func oldmain() {
	fmtImport := gopoet.TypeReferenceFromInstance(fmt.Println)
	byteRef := gopoet.TypeReferenceFromInstance(&bytes.Buffer{})

	sampleStruct := gopoet.NewStructSpec("A")
	sampleMethodSpec := gopoet.NewMethodSpec("sampleMethod", "a", false, sampleStruct)
	sampleStruct.MethodAndAttach("blahMethod", "b", false)

	mainSpec := gopoet.NewFuncSpec("main")
	mainSpec.Statement("$T($S)", fmtImport, "Calling hello...")

	helloSpec := gopoet.NewFuncSpec("hello")
	helloSpec.Parameter("byteThing", byteRef)
	helloSpec.BlockStart("for i:= 0; i < 5; i++")
	helloSpec.Statement("$T($L)", fmtImport, "i")
	helloSpec.BlockEnd()

	mainSpec.Statement("$L()", helloSpec.Name)

	fileSpec := gopoet.NewFileSpec("main")
	fileSpec.CodeBlock(sampleStruct)
	fileSpec.CodeBlock(mainSpec)
	fileSpec.CodeBlock(helloSpec)
	fileSpec.CodeBlock(sampleMethodSpec)

	fmt.Println(fileSpec.String())
}
