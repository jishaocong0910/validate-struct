package validate_test

import (
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

func TestCheckEmbed_Dive(t *testing.T) {
	r := require.New(t)
	checklistSuccess(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url", DemoBase: &validate.DemoBase{Id: "id"}})
	})
	checklistFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url", DemoBase: &validate.DemoBase{}})
	}, "id must not be blank")
	checklistFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{Url: "url"})
	}, "id is required")
}

func TestCheckEmbed_ShouldStop(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Embed(&validate.DemoImage{})
	}, "url must not be blank")
}
