
package main

import (
	. "fmt"
	"reflect"
    "runtime"
)

var (
	test int = 5
	
)

func GetFunctionName(i interface{}) string {
	
	
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	// return reflect.TypeOf(i)
}

type A struct {
	
}

func (a A) helloStatic() {

}

func (a *A) hello() {

}

func main() {
	//myClass = ClassName.forName("fmt", "Println")
	fmt.Println("Hello World!")
	//helloType := reflect.TypeOf(hello)
	a := &A{}
	// addStatement("$.hello()", fmt)
	fmt.Println("name:", GetFunctionName(A.hello))
}



myFmt := ImportDef.Get("fmt")
blankImport := ImportDef.GetInitialization("fmt")
aliasImport := ImportDef.GetAlias("fmt", "tmf")
aliasImport := ImportDef.GetAliasedType("fmt", "tmf", "PrintLn")
unqualifiedFmt := ImportDef.GetUnqualified("fmt", "PrintLn")
myFmtPrintLn := ImportDef.Get("fmt", "PrintLn")


inlineStruct := StructSpec{}

mainSpec := FuncSpec { Name="main"}
mainSpec.AddStatement("$T.PrintLn($S)", ImportDef.Get("fmt"), "Hello World")
mainSpec.AddStatement("$T.($S)", myFmtPrintLn, "Hello World")
mainSpec.AddStatement("$T.helloStatic()", ImportDef.Get(A{}))
mainSpec.AddStatement("a := &$T{}", ImportDef.Get(A{}))
mainSpec.AddStatement("a.hello()")

fileSpec := FileSpec { Package="main" }
fileSpec.AddCodeBlock(mainSpec)
fileSpec.String()

