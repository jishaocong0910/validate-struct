package validate_test

import (
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

func TestCheckEmbed_Required(t *testing.T) {
	r := require.New(t)
	testSuccess(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url", DemoBase: &validate.DemoBase{Id: "id"}})
	})
	testFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url", DemoBase: &validate.DemoBase{}})
	}, "id must not be blank")
	testFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url"})
	}, "id must not be blank")
}

func TestCheckEmbed_ShouldStop(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{})
	}, "url must not be blank")
}
