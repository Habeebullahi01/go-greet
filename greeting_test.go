package greeting

import (
	"regexp"
	"testing"
)

// TestHelloName calls greeting.Hello function with a name and checks for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Habeeb"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Habeeb")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Habeeb")=%q,%v, wants match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty checks that the greeting.Hello function returns an error when called with an empty string.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if err == nil || msg != "" {
		t.Fatalf(`Hello("")=%v,%q wants "", error`, msg, err)
	}
}
