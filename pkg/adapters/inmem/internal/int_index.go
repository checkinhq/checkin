package internal

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

// IntFieldIndex is used to extract a int field from an object using
// reflection and builds an index on that field.
type IntFieldIndex struct {
	Field string
}

func (in *IntFieldIndex) FromObject(obj interface{}) (bool, []byte, error) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v) // Dereference the pointer if any

	fv := v.FieldByName(in.Field)
	if !fv.IsValid() {
		return false, nil,
			fmt.Errorf("field '%s' for %#v is invalid", in.Field, obj)
	}

	// Check the type
	k := fv.Kind()
	size, ok := IsIntType(k)
	if !ok {
		return false, nil, fmt.Errorf("field %q is of type %v; want an int", in.Field, k)
	}

	// Get the value and encode it
	val := fv.Int()
	buf := make([]byte, size)
	binary.PutUvarint(buf, uint64(val))

	return true, buf, nil
}

func (in *IntFieldIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("must provide only a single argument")
	}

	v := reflect.ValueOf(args[0])
	if !v.IsValid() {
		return nil, fmt.Errorf("%#v is invalid", args[0])
	}

	k := v.Kind()
	size, ok := IsIntType(k)
	if !ok {
		return nil, fmt.Errorf("arg is of type %v; want an int", k)
	}
	buf := make([]byte, size)
	val := v.Int()
	binary.PutUvarint(buf, uint64(val))

	return buf, nil
}

// IsIntType returns whether the passed type is a type of Int and the number
// of bytes the type is. To avoid platform specific sizes, the uint type returns
// 8 bytes regardless of if it is smaller.
func IsIntType(k reflect.Kind) (size int, okay bool) {
	switch k {
	case reflect.Int:
		return 8, true
	case reflect.Int8:
		return 1, true
	case reflect.Int16:
		return 2, true
	case reflect.Int32:
		return 4, true
	case reflect.Int64:
		return 8, true
	default:
		return 0, false
	}
}
