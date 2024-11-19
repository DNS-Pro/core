package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type HttpAuthenticator struct {
	Authenticator
	Url string `validate:"required,http_url"`
}

func (a *HttpAuthenticator) run(ctx context.Context) error {
	resp, err := http.Get(a.Url)
	if err != nil {
		return fmt.Errorf("error requesting url (%s): %s", a.Url, err)
	}
	sc := resp.StatusCode
	if sc < 200 || sc >= 300 {
		return fmt.Errorf("unexpected status code requesting url (%s): %d", a.Url, sc)
	}
	return nil
}
func NewHttpAuthenticator(runEvery time.Duration, url string) (*HttpAuthenticator, error) {
	v := &HttpAuthenticator{
		Url: url,
	}
	a := Authenticator{
		iAuthenticator: v,
		runEvery:       runEvery,
		aType:          AUTH_HTTP,
	}
	v.Authenticator = a
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(v); err != nil {
		return nil, fmt.Errorf("can not validate provided config: %s", err)
	}
	return v, nil
}
