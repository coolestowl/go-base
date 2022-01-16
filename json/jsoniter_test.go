package json_test

import (
	"bytes"
	stdjson "encoding/json"
	"testing"

	jsoniter "github.com/coolestowl/go-base/json"
)

type jsonTest struct {
	Object struct {
		SubSlice []int
	}
	Slice []string
	Str   string
}

func TestMarshal(t *testing.T) {
	v := jsonTest{
		Object: struct{ SubSlice []int }{
			SubSlice: []int{2, 3, 3},
		},
		Slice: []string{"2", "3", "3"},
		Str:   "hello",
	}

	got, err := jsoniter.Marshal(v)
	if err != nil {
		t.Error(err)
	}

	expected, _ := stdjson.Marshal(v)
	if !bytes.Equal(expected, got) {
		t.Errorf("expected '%v' got '%v'", expected, got)
	}
}

func TestUnmarshal(t *testing.T) {
	src := `{"object":{"sub_slice":[2,3,3]},"slice":["2","3","3"],"str":"hello"}`

	var got jsonTest
	if err := jsoniter.Unmarshal([]byte(src), &got); err != nil {
		t.Error(err)
	}

	var expected jsonTest
	_ = stdjson.Unmarshal([]byte(src), &expected)

	if !intSliceEqual(got.Object.SubSlice, expected.Object.SubSlice) {
		t.Errorf("expected '%v' got '%v'", expected, got)
	}
}

func BenchmarkMarshal(b *testing.B) {
	v := jsonTest{
		Object: struct{ SubSlice []int }{
			SubSlice: []int{2, 3, 3},
		},
		Slice: []string{"2", "3", "3"},
		Str:   "hello",
	}

	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(v)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	var (
		bytes = []byte(`{"object":{"sub_slice":[2,3,3]},"slice":["2","3","3"],"str":"hello"}`)
		obj   = jsonTest{}
	)
	for i := 0; i < b.N; i++ {
		_ = jsoniter.Unmarshal(bytes, &obj)
	}
}

func intSliceEqual(this, other []int) bool {
	if len(this) != len(other) {
		return false
	}
	for idx := range this {
		if this[idx] != other[idx] {
			return false
		}
	}
	return true
}
