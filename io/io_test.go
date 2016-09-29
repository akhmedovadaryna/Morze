package io

import "testing"

func TestInput_from_file(t *testing.T) {
	iff, err := Input_from_file("input_text")
	if iff != "GoLang Examples \n go get" {
		t.Error("Expected GoLang Examples go get; got", iff)
	}
}