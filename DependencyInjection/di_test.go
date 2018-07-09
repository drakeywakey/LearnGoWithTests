package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Drake")

	got := buffer.String()
	want := "Sup, Drake"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
