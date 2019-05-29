package slice_test

import (
	"testing"

	"github.com/danielbintar/angel/server-library/slice"

	"github.com/stretchr/testify/assert"
)

func TestInStrings(t *testing.T) {
	cases := generateInStringTestCase()

	for _, c := range cases {
		assert.Equal(t, slice.InStrings(c.String, c.Strings), c.Result)
	}
}

type InStringTestCase struct {
	Strings []string
	String  string
	Result  bool
}

func generateInStringTestCase() []InStringTestCase {
	cases := []InStringTestCase {
		InStringTestCase {
			Strings: []string{},
			String: "",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{},
			String: "a",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"a"},
			String: "",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"a"},
			String: "A",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"A"},
			String: "a",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"ab"},
			String: "a",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"a"},
			String: "ab",
			Result: false,
		},
		InStringTestCase {
			Strings: []string{"a"},
			String: "a",
			Result: true,
		},
		InStringTestCase {
			Strings: []string{"A"},
			String: "A",
			Result: true,
		},
		InStringTestCase {
			Strings: []string{"b", "A"},
			String: "A",
			Result: true,
		},
		InStringTestCase {
			Strings: []string{"A", "b"},
			String: "A",
			Result: true,
		},
	}
	return cases
}
