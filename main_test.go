package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestHello(t *testing.T) {
	out, in, _ := makeUI()
	if out.Text != "Hello world" {
		t.Errorf("Incorrect initial hello")
	}

	test.Type(in, "foo")
	if out.Text != "Hello foo !" {
		t.Errorf("Incorrect hello %s", out.Text)
	}
}
