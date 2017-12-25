package inmem

import (
	"fmt"
	"reflect"

	"github.com/segmentio/ksuid"
)

// KSUIDFieldIndex is used to extract a field from an object
// using reflection and builds an index on that field.
type KSUIDFieldIndex struct {
	Field string
}

func (s *KSUIDFieldIndex) FromObject(obj interface{}) (bool, []byte, error) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v) // Dereference the pointer if any

	fv := v.FieldByName(s.Field)
	if !fv.IsValid() {
		return false, nil, fmt.Errorf("field '%s' for %#v is invalid", s.Field, obj)
	}

	uid, ok := fv.Interface().(ksuid.KSUID)
	if !ok {
		return false, nil, fmt.Errorf("field must be a KSUID: %#v", uid)
	}

	return true, uid.Bytes(), nil
}

func (s *KSUIDFieldIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("must provide only a single argument")
	}

	arg, ok := args[0].(ksuid.KSUID)
	if !ok {
		return nil, fmt.Errorf("argument must be a string: %#v", args[0])
	}

	return arg.Bytes(), nil
}

func (s *KSUIDFieldIndex) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	val, err := s.FromArgs(args...)
	if err != nil {
		return nil, err
	}

	return val, nil
}
