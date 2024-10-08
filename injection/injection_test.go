package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gerald")

	got := buffer.String()
	want := "Hello, Gerald"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
