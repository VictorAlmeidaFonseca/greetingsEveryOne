package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName call greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Glady"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Glady")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Glady")  = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty call greetings.Heloo with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
