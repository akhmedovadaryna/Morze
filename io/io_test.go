package io

import "testing"

func TestInput_from_file(t *testing.T) {
	iff, err := Input_from_file("C:/Users/Дарина/PycharmProjects/Morze/files/input_text")
	if iff != "Go" {
		t.Error("Expected Go", iff , err)
	}
}
