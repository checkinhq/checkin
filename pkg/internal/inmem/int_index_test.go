package inmem

import (
	"testing"

	"bytes"
	"encoding/binary"
	"strings"
)

type TestIntObject struct {
	ID       string
	Foo      string
	Bar      int
	Baz      string
	Bam      *bool
	Empty    string
	Qux      []string
	QuxEmpty []string
	Zod      map[string]string
	ZodEmpty map[string]string
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
}

func testObjInt() *TestIntObject {
	b := true
	obj := &TestIntObject{
		ID:  "my-cool-obj",
		Foo: "Testing",
		Bar: 42,
		Baz: "yep",
		Bam: &b,
		Qux: []string{"Test", "Test2"},
		Zod: map[string]string{
			"Role":          "Server",
			"instance_type": "m3.medium",
			"":              "asdf",
		},
		Int:   int(1),
		Int8:  int8(8),
		Int16: int16(16),
		Int32: int32(32),
		Int64: int64(64),
	}
	return obj
}

func TestIntFieldIndex_FromObject(t *testing.T) {
	obj := testObjInt()

	eint := make([]byte, 8)
	eint8 := make([]byte, 1)
	eint16 := make([]byte, 2)
	eint32 := make([]byte, 4)
	eint64 := make([]byte, 8)
	binary.PutUvarint(eint, uint64(obj.Int))
	binary.PutUvarint(eint8, uint64(obj.Int8))
	binary.PutUvarint(eint16, uint64(obj.Int16))
	binary.PutUvarint(eint32, uint64(obj.Int32))
	binary.PutUvarint(eint64, uint64(obj.Int64))

	cases := []struct {
		Field         string
		Expected      []byte
		ErrorContains string
	}{
		{
			Field:    "Int",
			Expected: eint,
		},
		{
			Field:    "Int8",
			Expected: eint8,
		},
		{
			Field:    "Int16",
			Expected: eint16,
		},
		{
			Field:    "Int32",
			Expected: eint32,
		},
		{
			Field:    "Int64",
			Expected: eint64,
		},
		{
			Field:         "IntGarbage",
			ErrorContains: "is invalid",
		},
		{
			Field:         "ID",
			ErrorContains: "want an int",
		},
	}

	for _, c := range cases {
		t.Run(c.Field, func(t *testing.T) {
			indexer := IntFieldIndex{c.Field}
			ok, val, err := indexer.FromObject(obj)

			if err != nil {
				if ok {
					t.Fatalf("okay and error")
				}

				if c.ErrorContains != "" && strings.Contains(err.Error(), c.ErrorContains) {
					return
				} else {
					t.Fatalf("Unexpected error %v", err)
				}
			}

			if !ok {
				t.Fatalf("not okay and no error")
			}

			if !bytes.Equal(val, c.Expected) {
				t.Fatalf("bad: %#v %#v", val, c.Expected)
			}

		})
	}
}

func TestIntFieldIndex_FromArgs(t *testing.T) {
	indexer := IntFieldIndex{"Foo"}
	_, err := indexer.FromArgs()
	if err == nil {
		t.Fatalf("should get err")
	}

	_, err = indexer.FromArgs(uint(1), uint(2))
	if err == nil {
		t.Fatalf("should get err")
	}

	_, err = indexer.FromArgs("foo")
	if err == nil {
		t.Fatalf("should get err")
	}

	obj := testObjInt()
	eint := make([]byte, 8)
	eint8 := make([]byte, 1)
	eint16 := make([]byte, 2)
	eint32 := make([]byte, 4)
	eint64 := make([]byte, 8)
	binary.PutUvarint(eint, uint64(obj.Int))
	binary.PutUvarint(eint8, uint64(obj.Int8))
	binary.PutUvarint(eint16, uint64(obj.Int16))
	binary.PutUvarint(eint32, uint64(obj.Int32))
	binary.PutUvarint(eint64, uint64(obj.Int64))

	val, err := indexer.FromArgs(obj.Int)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, eint) {
		t.Fatalf("bad: %#v %#v", val, eint)
	}

	val, err = indexer.FromArgs(obj.Int8)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, eint8) {
		t.Fatalf("bad: %#v %#v", val, eint8)
	}

	val, err = indexer.FromArgs(obj.Int16)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, eint16) {
		t.Fatalf("bad: %#v %#v", val, eint16)
	}

	val, err = indexer.FromArgs(obj.Int32)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, eint32) {
		t.Fatalf("bad: %#v %#v", val, eint32)
	}

	val, err = indexer.FromArgs(obj.Int64)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, eint64) {
		t.Fatalf("bad: %#v %#v", val, eint64)
	}
}
