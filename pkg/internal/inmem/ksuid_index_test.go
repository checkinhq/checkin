package inmem

import (
	"testing"

	"bytes"

	"github.com/segmentio/ksuid"
)

type TestKsuidObject struct {
	ID       ksuid.KSUID
	Foo      string
	Bar      int
	Baz      string
	Bam      *bool
	Empty    string
	Qux      []string
	QuxEmpty []string
	Zod      map[string]string
	ZodEmpty map[string]string
}

func testObjKsuid() *TestKsuidObject {
	b := true
	obj := &TestKsuidObject{
		ID:  ksuid.Max,
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
	}
	return obj
}

func TestKsuidFieldIndex_FromObject(t *testing.T) {
	obj := testObjKsuid()

	indexer := KSUIDFieldIndex{"ID"}
	ok, val, err := indexer.FromObject(obj)

	if err != nil {
		if ok {
			t.Fatalf("okay and error")
		}

		t.Fatalf("Unexpected error %v", err)
	}

	if !ok {
		t.Fatalf("not okay and no error")
	}

	if !bytes.Equal(val, ksuid.Max.Bytes()) {
		t.Fatalf("bad: %#v %#v", val, ksuid.Max.Bytes())
	}
}

func TestKsuidFieldIndex_FromArgs(t *testing.T) {
	indexer := KSUIDFieldIndex{"ID"}
	_, err := indexer.FromArgs()
	if err == nil {
		t.Fatalf("should get err")
	}

	_, err = indexer.FromArgs("foo")
	if err == nil {
		t.Fatalf("should get err")
	}

	obj := testObjKsuid()

	val, err := indexer.FromArgs(obj.ID)
	if err != nil {
		t.Fatalf("bad: %v", err)
	}
	if !bytes.Equal(val, ksuid.Max.Bytes()) {
		t.Fatalf("bad: %#v %#v", val, ksuid.Max.Bytes())
	}
}
