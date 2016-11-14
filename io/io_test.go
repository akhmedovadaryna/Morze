package io

import "testing"

func TestSerial_Input_from_file(t *testing.T) {
	siff, err := Serial_Input_from_file("input_text1")
	if siff != "Rose in Bloom" {
		t.Error("Expected Rose in Bloom", siff , err)
	}
}

func TestSerial_Output_from_file(t *testing.T) {
	soff := Serial_Output_from_file("output", "res")
	if soff == nil {
		t.Error("Expected not nil, got ", soff)
	}
}