package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {

	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("kerem@aol.com")
	if err == nil {
		t.Error("it is an email")
	}

	_, err = IsEmail("kerem@aol")
	if err != nil {
		t.Error("it is not an email")
	}

}
