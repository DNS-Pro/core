package errs

import (
	"fmt"
)

type IAppErr interface {
	error
}
type baseAppErr struct {
	message string
	txt     string
}

func (e *baseAppErr) Error() string {
	if e.txt != "" {
		return fmt.Sprintf("%s: %s", e.message, e.txt)
	} else {
		return e.message
	}
}

type AppConfigValidationErr struct{ *baseAppErr }
type AppDefaultValueErr struct{ *baseAppErr }

func NewConfigValidationErr(e error) IAppErr {
	return AppConfigValidationErr{&baseAppErr{message: "error validating app config", txt: e.Error()}}
}
func NewConfigDefaultValueErr(e error) IAppErr {
	return AppDefaultValueErr{&baseAppErr{message: "error assigning default app config values", txt: e.Error()}}
}
