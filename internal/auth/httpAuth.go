package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/onsi/ginkgo/v2"
)

type HttpAuther struct {
	Url string `validate:"required,http_url"`
}

func (a *HttpAuther) Run(ctx context.Context) error {
	resp, err := http.Get(a.Url)
	if err != nil {
		return fmt.Errorf("error requesting url (%s): %s", a.Url, err)
	}
	sc := resp.StatusCode
	ginkgo.GinkgoWriter.Printf("%s:%d", a.Url, sc)
	if sc < 200 || sc >= 300 {
		return fmt.Errorf("unexpected status code requesting url (%s): %d", a.Url, sc)
	}
	return nil
}
func (a *HttpAuther) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(a)
}
func (a *HttpAuther) SetDefaults() error {
	return defaults.Set(a)
}
func NewHttpAuther(url string) *HttpAuther {
	return &HttpAuther{
		Url: url,
	}
}
