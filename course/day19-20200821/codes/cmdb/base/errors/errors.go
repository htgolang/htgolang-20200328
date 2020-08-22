package errors

import (
	"github.com/astaxie/beego/validation"
)

type Errors struct {
	errors map[string][]string
}

func (e *Errors) Add(key, err string) {
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key], err)
}

func (e *Errors) AddValidation(valid *validation.Validation) {
	if valid.HasErrors() {
		for key, errs := range valid.ErrorsMap {
			for _, err := range errs {
				e.Add(key, err.Message)
			}
		}
	}
}

func (e *Errors) Errors() map[string][]string {
	return e.errors
}

func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}

func (e *Errors) HasErrors() bool {
	return len(e.errors) != 0
}

func New() *Errors {
	return &Errors{
		errors: make(map[string][]string),
	}
}
