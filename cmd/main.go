package main

import (
	"fmt"

	"gitlab.com/kdeenanauth/gopoet"
)

func main() {
	f := gopoet.NewFuncSpec("main")
	fmtImport := gopoet.Import{
		Package: "fmt",
		Name:    "Println",
	}

	f.Statement("$T($S)", fmtImport, "Hello Kevin")

	fileSpec := gopoet.NewFileSpec("blah")
	fileSpec.CodeBlock(f)

	fmt.Println(fileSpec.String())
}

// myFmt := ImportDef.Get("fmt", nil)
// blankImport := ImportDef.GetInitialization("fmt")
// aliasImport := ImportDef.GetAlias("fmt", "tmf")
// aliasImport := ImportDef.GetAliasedType("fmt", "tmf", "PrintLn")
// unqualifiedFmt := ImportDef.GetUnqualified("fmt", "PrintLn")
// myFmtPrintLn := ImportDef.Get("fmt", "PrintLn")

// inlineStruct := StructSpec{}

// mainSpec := FuncSpec { Name="main"}
// mainSpec.AddStatement("$T.PrintLn($S)", ImportDef.Get("fmt"), "Hello World")
// mainSpec.AddStatement("$T.($S)", myFmtPrintLn, "Hello World")
// mainSpec.AddStatement("$T.helloStatic()", ImportDef.Get(A{}))
// mainSpec.AddStatement("a := &$T{}", ImportDef.Get(A{}))
// mainSpec.AddStatement("a.hello()")

// fileSpec := FileSpec { Package="main" }
// fileSpec.AddCodeBlock(mainSpec)
// fileSpec.String()

// intSpec := InterfaceSpec
// intSpec.addMethod("Hello")
