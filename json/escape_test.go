package json_test

import (
	"testing"

	"github.com/coolestowl/go-base/json"
)

func TestEscape(t *testing.T) {
	cases := []struct {
		Src      string
		Expected string
	}{
		{
			Src:      "hhhh",
			Expected: `"hhhh"`,
		},
		{
			Src:      `{"text":"hhhh"}`,
			Expected: `"{\"text\":\"hhhh\"}"`,
		},
	}
	for _, c := range cases {
		got, err := json.Escape(c.Src)
		if err != nil {
			t.Fatal(err)
		}
		if got != c.Expected {
			t.Fatalf("expected '%s' got '%s'", c.Expected, got)
		}
	}
}
