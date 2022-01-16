package json

import "testing"

func TestDiffString(t *testing.T) {
	cases := []struct {
		argA     string
		argB     string
		expected string
	}{{
		argA:     `{"obj":{"a":"aaa","b":"bbb"}}`,
		argB:     `{"obj":{"a":"","b":"bbb"}}`,
		expected: "obj.a",
	}, {
		argA:     `{"obj":{"a":"aaa","b":"bbb"}}`,
		argB:     `{"obj":{"a":"","b":""}}`,
		expected: "obj",
	}, {
		argA:     `{"obj":{"a":"aaa","b":"bbb"}}`,
		argB:     `{"obj":{"a":"aaa","b":"bbb","c":"ccc"}}`,
		expected: "obj.c",
	}}
	for _, c := range cases {
		got, err := DiffString(c.argA, c.argB)
		if err != nil {
			t.Error(err)
		}
		if c.expected != got {
			t.Errorf("expected '%v' got '%v'", c.expected, got)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		strs     []string
		expected string
	}{{
		strs: []string{
			"", "a", "abc",
		},
		expected: "",
	}, {
		strs: []string{
			"a", "ab", "abc",
		},
		expected: "a",
	}, {
		strs: []string{
			"ajns", "ajns.a", "ajns.b",
		},
		expected: "ajns",
	}}

	for _, c := range cases {
		got := longestCommonPrefix(c.strs)
		if got != c.expected {
			t.Errorf("expected '%v' got '%v'", c.expected, got)
		}
	}
}
