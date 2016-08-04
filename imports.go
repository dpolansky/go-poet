package gopoet

import (
	"reflect"
)

// Import represents information needed to qualify a type and build imports
// for a FileSpec
type Import struct {
	Package     string
	Alias       string
	Name        string
	Unqualified bool
}

// ImportFromInstance creates an Import from an existing type
func ImportFromInstance(t interface{}) Import {
	reflectType := reflect.TypeOf(t)

	if reflectType == nil {
		panic("Invalid nil instance without associated type")
	}

	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	return Import{
		Package: reflectType.PkgPath(),
		Name:    reflectType.Name(),
	}
}

// GetPackage returns the package name for this import
func (i *Import) GetPackage() string {
	return i.Package
}

// GetPackageAlias returns the alias for referencing the package
func (i *Import) GetPackageAlias() string {
	return i.Alias
}

// NeedsQualifier returns whether the name needs to be qualified with the package/alias
func (i *Import) NeedsQualifier() bool {
	return !i.Unqualified
}

// GetName returns the name of the type/function being referenced within the package.
func (i *Import) GetName() string {
	return i.Name
}
