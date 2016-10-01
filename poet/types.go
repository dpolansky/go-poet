package poet

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// UnqualifiedPrefix The prefix for type aliases that will be interpreted as unqualified
const UnqualifiedPrefix = "_unqualified"

var (
	// String A TypeReference for string
	String = TypeReferenceFromInstance("")
	// Bool A TypeReference for bool
	Bool = TypeReferenceFromInstance(false)
	// Int A TypeReference for int
	Int = TypeReferenceFromInstance(0)
	// Int8 A TypeReference for int8
	Int8 = TypeReferenceFromInstance(int8(0))
	// Int16 A TypeReference for int16
	Int16 = TypeReferenceFromInstance(int16(0))
	// Int32 A TypeReference for int32
	Int32 = TypeReferenceFromInstance(int32(0))
	// Int64 A TypeReference for int64
	Int64 = TypeReferenceFromInstance(int64(0))
	// Uint A TypeReference for uint
	Uint = TypeReferenceFromInstance(uint(0))
	// Uint8 A TypeReference for uint8
	Uint8 = TypeReferenceFromInstance(uint8(0))
	// Uint16 A TypeReference for uint16
	Uint16 = TypeReferenceFromInstance(uint16(0))
	// Uint32 A TypeReference for uint32
	Uint32 = TypeReferenceFromInstance(uint32(0))
	// Uint64 A TypeReference for uint64
	Uint64 = TypeReferenceFromInstance(uint64(0))
	// Uintptr A TypeReference for uintptr
	Uintptr = TypeReferenceFromInstance(uintptr(0))
	// Float32 A TypeReference for float32
	Float32 = TypeReferenceFromInstance(float32(0))
	// Float64 A TypeReference for float64
	Float64 = TypeReferenceFromInstance(float64(0))
	// Complex64 A TypeReference for complex64
	Complex64 = TypeReferenceFromInstance(complex64(0))
	// Complex128 A TypeReference for complex128
	Complex128 = TypeReferenceFromInstance(complex128(0))
	// Byte A TypeReference for byte
	Byte = TypeReferenceFromInstanceWithCustomName(uint8(0), "byte")
	// Rune A TypeReference for rune
	Rune = TypeReferenceFromInstanceWithCustomName(int32(0), "rune")
)

// TypeReferenceFromInstance creates a TypeReference from an instance of a variable
func TypeReferenceFromInstance(t interface{}) TypeReference {
	return newTypeReferenceFromInstance(t, "")
}

// TypeReferenceFromInstanceWithAlias creates a TypeReference from an instance of a variable
// with the given package alias
func TypeReferenceFromInstanceWithAlias(t interface{}, alias string) TypeReference {
	return newTypeReferenceFromInstance(t, alias)
}

// TypeReferenceFromInstanceWithCustomName creates a TypeReference from an instance of a variable
// with the given custom name, for use of a type alias's name rather than the underlying
// reflect type.
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

type typeReferenceMap struct {
	KeyType   TypeReference
	ValueType TypeReference
	prefix    string
}

var _ TypeReference = (*typeReferenceMap)(nil)

func newTypeReferenceFromMap(t interface{}, prefix string) TypeReference {
	refType := reflect.TypeOf(t)

	return &typeReferenceMap{
		KeyType:   newTypeReferenceFromInstance(reflect.New(refType.Key()).Elem().Interface(), ""),
		ValueType: newTypeReferenceFromInstance(reflect.New(refType.Elem()).Elem().Interface(), ""),
		prefix:    prefix,
	}
}

func (t *typeReferenceMap) GetImports() []Import {
	imports := []Import{}

	imports = append(imports, t.KeyType.GetImports()...)
	imports = append(imports, t.ValueType.GetImports()...)
	return imports
}

func (t *typeReferenceMap) GetName() string {
	return fmt.Sprintf("%smap[%s]%s", t.prefix, t.KeyType.GetName(), t.ValueType.GetName())
}

type typeReferenceValue struct {
	Import *ImportSpec
	Name   string
	prefix string
}

var _ TypeReference = (*typeReferenceValue)(nil)

func newTypeReferenceFromValue(t interface{}, alias string) TypeReference {
	refType := reflect.TypeOf(t)
	result := &typeReferenceValue{}

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

func (t *typeReferenceValue) GetImports() []Import {
	return []Import{t.Import}
}

func (t *typeReferenceValue) GetName() string {
	result := bytes.Buffer{}

	result.WriteString(t.prefix)
	result.WriteString(t.Import.getQualifier())
	result.WriteString(t.Name)

	return result.String()
}

type typeReferenceFunc struct {
	Import *ImportSpec
	Name   string
}

var _ TypeReference = (*typeReferenceFunc)(nil)

func newTypeReferenceFromFunction(t interface{}, alias string) TypeReference {
	// split up the function's name from its package path
	n := runtime.FuncForPC(reflect.ValueOf(t).Pointer()).Name()
	ndxOfLastDot := strings.LastIndex(n, ".")

	return &typeReferenceFunc{
		Import: &ImportSpec{
			Qualified: true,
			Package:   n[:ndxOfLastDot],
			Alias:     alias,
		},
		Name: n[ndxOfLastDot+1:],
	}
}

func (t *typeReferenceFunc) GetImports() []Import {
	return []Import{t.Import}
}

func (t *typeReferenceFunc) GetName() string {
	result := bytes.Buffer{}

	result.WriteString(t.Import.getQualifier())
	result.WriteString(t.Name)

	return result.String()
}
