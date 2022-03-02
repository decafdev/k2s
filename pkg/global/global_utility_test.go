package global_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techdecaf/k2s/v2/pkg/global"
)

func TestUtilityFunctions(t *testing.T) {
	// given struct
	type given struct {
		params []string
		expect string
	}

	tests := make(map[string]given)

	tests["when more than one string is valid"] = given{
		params: []string{"a", "b"},
		expect: "a",
	}

	tests["when more than one empty string"] = given{
		params: []string{"", "", "c"},
		expect: "c",
	}

	tests["when no strings are valid"] = given{
		params: []string{"", "", ""},
		expect: "",
	}

	for when, given := range tests {
		t.Run(when, func(t *testing.T) {
			res := global.StringDefault(given.params...)
			assert.Equal(t, res, given.expect)
		})
	}
}
