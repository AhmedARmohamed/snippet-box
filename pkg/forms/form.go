package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// EmailRX is pattern to check if string is a valid email
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Form is a struct made to hold form data and form errors
type Form struct {
	url.Values
	Errors errors
}

// New is func to create new Form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required method to check that specific fields in the form
// data are present and not blank.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MaxLength checks that a specific field in the form contains
// a appropriate number of characters.
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

// MinLength checks that a specific field in the form contains
// a appropriate number of characters.
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too small (minimum is %d characters)", d))
	}
}

// PermittedValues checks that a specific field in the form matches
// one fo a set a specific values.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

// Valid returns true if previous error checks went okay.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// MatchesPattern checks if specific field matches the pattern
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}
