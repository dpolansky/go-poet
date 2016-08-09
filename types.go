package gopoet

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const UnqualifiedPrefix = "_unqualified"

// TypeReferenceFromInstance creates an TypeReference from an existing type
func TypeReferenceFromInstance(t interface{}) TypeReference {
	return newTypeReferenceFromInstance(t, "")
}

func TypeReferenceFromInstanceWithAlias(t interface{}, alias string) TypeReference {
	return newTypeReferenceFromInstance(t, alias)
}

func TypeReferenceFromInstanceWithCustomName(t interface{}, name string) TypeReference {
	typeRef := &typeReferenceWithCustomName{
		TypeReference: newTypeReferenceFromInstance(t, ""),
		name:          name,
	}

	return typeRef
}

type typeReferenceWithCustomName struct {
	TypeReference
	name string
}

func (t *typeReferenceWithCustomName) GetName() string {
	return t.name
}

func newTypeReferenceFromInstance(t interface{}, alias string) TypeReference {
	reflectType := reflect.TypeOf(t)
	if reflectType == nil {
		panic("Invalid nil instance without associated type")
	}

	if reflectType.Kind() == reflect.Func {
		return newTypeReferenceFromFunction(t, alias)
	}

	return newTypeReferenceFromValue(t, alias)
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
		KeyType:   newTypeReferenceFromInstance(reflect.New(refType.Key()).Elem().Interface(), ""),
		ValueType: newTypeReferenceFromInstance(reflect.New(refType.Elem()).Elem().Interface(), ""),
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

type TypeReferenceValue struct {
	Import *ImportSpec
	Name   string
	prefix string
}

var _ TypeReference = (*TypeReferenceValue)(nil)

func newTypeReferenceFromValue(t interface{}, alias string) TypeReference {
	refType := reflect.TypeOf(t)
	result := &TypeReferenceValue{}

	result.prefix, refType = dereferenceType("", refType)

	switch refType.Kind() {
	case reflect.Interface:
		fallthrough
	case reflect.Struct:
		result.Import = &ImportSpec{
			Qualified: !strings.HasPrefix(refType.Name(), UnqualifiedPrefix),
			Package:   refType.PkgPath(),
			Alias:     alias,
		}
	case reflect.Map:
		return newTypeReferenceFromMap(reflect.New(refType).Elem().Interface(), result.prefix)
	}

	result.Name = strings.TrimPrefix(refType.Name(), UnqualifiedPrefix)

	return result
}

func dereferenceType(prefix string, refType reflect.Type) (string, reflect.Type) {
	for {
		if refType.Kind() == reflect.Ptr {
			refType = refType.Elem()
			// interfaces are already pointers, so don't need to add prefix
			if refType.Kind() != reflect.Interface {
				prefix += "*"
			}
		} else if refType.Kind() == reflect.Slice || refType.Kind() == reflect.Array {
			prefix += "[]"
			refType = refType.Elem()
		} else if refType.Kind() == reflect.Chan {
			prefix += refType.ChanDir().String() + " "
			refType = refType.Elem()
		} else {
			break
		}
	}

	return prefix, refType
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

func newTypeReferenceFromFunction(t interface{}, alias string) TypeReference {
	funcPtr := runtime.FuncForPC(reflect.ValueOf(t).Pointer())
	fullyQualifiedPieces := strings.Split(funcPtr.Name(), ".")

	if len(fullyQualifiedPieces) < 2 {
		panic(fmt.Sprintf("Could not create type reference from function, %#v", t))
	}

	return &TypeReferenceFunc{
		Import: &ImportSpec{
			Qualified: true,
			Package:   fullyQualifiedPieces[0],
			Alias:     alias,
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
