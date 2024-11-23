package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DNS-Pro/core/pkg/errs"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/onsi/ginkgo/v2"
)

type httpAuther struct {
	Url string `validate:"required,http_url"`
}

func (a *httpAuther) Run(ctx context.Context) error {
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
func (a *httpAuther) GetType() AuthType {
	return AUTH_HTTP
}

// ...
// NewHttpAuther validates and creates auther.
//
// Using factory is the only way to create a HttpAuther, so validated configs are ensured.
func NewHttpAuther(url string) (IAuther, error) {
	v := httpAuther{
		Url: url,
	}
	if err := defaults.Set(&v); err != nil {
		return nil, errs.NewConfigDefaultValueErr(err)
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&v); err != nil {
		return nil, errs.NewConfigValidationErr(err)
	}
	return &v, nil
}
