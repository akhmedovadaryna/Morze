package Morze

import "testing"

func TestMorze(t *testing.T) {
	var m string
	m = Morze("--._---____--._._-_")
	if m != "GO GET"{
		t.Error("Expected GO GET, got", m)
	}
}

func TestTextToMorze(t *testing.T) {
	var ttm string
	ttm = TextToMorze("GO GET")
	if ttm != "--._---____--._._-_" {
		t.Error("Expecteg --._---____--._._-_, got", ttm)
	}
}

func BenchmarkMorze(b *testing.B) {
	for n := 0; n < b.N; n++{
		Morze("--._---____--._._-_")
	}
}
func BenchmarkTextToMorze(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TextToMorze("GO GET")

	}
}