package hash_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/coolestowl/go-base/hash"
)

var (
	sha256Cases = []struct {
		src      string
		expected string
	}{{
		src:      "",
		expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}, {
		src:      "a",
		expected: "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb",
	}, {
		src:      "ab",
		expected: "fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603",
	}, {
		src:      "abc",
		expected: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
	}, {
		src:      "abcd",
		expected: "88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589",
	}}
)

func TestSHA256(t *testing.T) {
	for _, c := range sha256Cases {
		got := hash.SHA256([]byte(c.src))
		expected, _ := hex.DecodeString(c.expected)
		if !bytes.Equal(got, expected) {
			t.Errorf("expected '%v' got '%v'", expected, got)
		}
	}
}

func TestHexSHA256(t *testing.T) {
	for _, c := range sha256Cases {
		got := hash.HexSHA256([]byte(c.src))
		if got != c.expected {
			t.Errorf("expected '%v' got '%v'", c.expected, got)
		}
	}
}
