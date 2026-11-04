package validate_test

import (
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

type embedTest struct {
	field *string
}

func (e embedTest) Checklist(v *validate.Validator) {
	v.String(e.field, "field").Required()
}

func TestCheckEmbed_Required(t *testing.T) {
	r := require.New(t)
	testSuccess(r, func(v *validate.Validator) {
		v.Embed((*embedTest)(nil))
	})
	testFail(r, func(v *validate.Validator) {
		v.Embed(&embedTest{})
	}, "field is required")
}
