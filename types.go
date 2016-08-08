package gopoet

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type ImportSpec struct {
	Package   string
	Alias     string
	Qualified bool
}

func (i *ImportSpec) GetQualifier() string {
	if i == nil {
		return ""
	}

	result := bytes.Buffer{}

	if i.Qualified {
		if i.Alias != "" {
			result.WriteString(i.Alias)
		} else {
			result.WriteString(i.Package)
		}
		result.WriteString(".")
	}

	return result.String()
}

var _ Import = (*ImportSpec)(nil)

func (i *ImportSpec) GetAlias() string {
	return i.Alias
}

func (i *ImportSpec) GetPackage() string {
	return i.Package
}

type TypeReferenceMap struct {
	KeyType   TypeReference
	ValueType TypeReference
	prefix    string
}

var _ TypeReference = (*TypeReferenceMap)(nil)

func newTypeReferenceFromMap(t interface{}, prefix string) TypeReference {
	refType := reflect.TypeOf(t)

	return &TypeReferenceMap{
		KeyType:   newTypeReferenceFromInstance(reflect.New(refType.Key()).Elem().Interface()),
		ValueType: newTypeReferenceFromInstance(reflect.New(refType.Elem()).Elem().Interface()),
		prefix:    prefix,
	}
}

func (t *TypeReferenceMap) GetImports() []Import {
	imports := []Import{}

	imports = append(imports, t.KeyType.GetImports()...)
	imports = append(imports, t.ValueType.GetImports()...)
	return imports
}

func (t *TypeReferenceMap) GetName() string {
	return fmt.Sprintf("%smap[%s]%s", t.prefix, t.KeyType.GetName(), t.ValueType.GetName())
}

// TypeReferenceFromInstance creates an TypeReference from an existing type
func TypeReferenceFromInstance(t interface{}) TypeReference {
	return newTypeReferenceFromInstance(t)
}

// TypeReferenceFromInstance creates an TypeReference from an existing type
func TypeReferenceFromInstanceWithAlias(t interface{}, alias string) TypeReference {
	typeRef := &typeReferenceWithAlias{
		TypeReference: newTypeReferenceFromInstance(t),
		alias:         alias,
	}

	return typeRef
}

type typeReferenceWithAlias struct {
	TypeReference
	alias string
}

func (t *typeReferenceWithAlias) GetName() string {
	return t.alias
}

func newTypeReferenceFromInstance(t interface{}) TypeReference {
	reflectType := reflect.TypeOf(t)
	if reflectType == nil {
		panic("Invalid nil instance without associated type")
	}

	if reflectType.Kind() == reflect.Func {
		return newTypeReferenceFromFunction(t)
	}

	return newTypeReferenceFromValue(t)
}

type TypeReferenceValue struct {
	Import *ImportSpec
	Name   string
	prefix string
}

var _ TypeReference = (*TypeReferenceValue)(nil)

func newTypeReferenceFromValue(t interface{}) TypeReference {
	refType := reflect.TypeOf(t)
	result := &TypeReferenceValue{}

	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		// interfaces are already pointers, so don't need to add prefix
		if refType.Kind() != reflect.Interface {
			result.prefix += "*"
		}
	}

	if refType.Kind() == reflect.Slice || refType.Kind() == reflect.Array {
		result.prefix += "[]"
		refType = refType.Elem()
	}

	switch refType.Kind() {
	case reflect.Interface:
		fallthrough
	case reflect.Struct:
		result.Import = &ImportSpec{
			Qualified: true,
			Package:   refType.PkgPath(),
		}
		break
	case reflect.Map:
		return newTypeReferenceFromMap(reflect.New(refType).Elem().Interface(), result.prefix)
	}

	result.Name = refType.Name()
	return result
}

func (t *TypeReferenceValue) GetImports() []Import {
	return []Import{t.Import}
}

func (t *TypeReferenceValue) GetName() string {
	result := bytes.Buffer{}

	result.WriteString(t.prefix)
	result.WriteString(t.Import.GetQualifier())
	result.WriteString(t.Name)

	return result.String()
}

type TypeReferenceFunc struct {
	Import *ImportSpec
	Name   string
}

var _ TypeReference = (*TypeReferenceFunc)(nil)

func newTypeReferenceFromFunction(t interface{}) TypeReference {
	funcPtr := runtime.FuncForPC(reflect.ValueOf(t).Pointer())
	fullyQualifiedPieces := strings.Split(funcPtr.Name(), ".")

	if len(fullyQualifiedPieces) < 2 {
		panic(fmt.Sprintf("Could not create type reference from function, %#v", t))
	}

	return &TypeReferenceFunc{
		Import: &ImportSpec{
			Qualified: true,
			Package:   fullyQualifiedPieces[0],
		},
		Name: fullyQualifiedPieces[1],
	}
}

func (t *TypeReferenceFunc) GetImports() []Import {
	return []Import{t.Import}
}

func (t *TypeReferenceFunc) GetName() string {
	result := bytes.Buffer{}

	result.WriteString(t.Import.GetQualifier())
	result.WriteString(t.Name)

	return result.String()
}
