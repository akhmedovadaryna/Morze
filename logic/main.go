package Morze
import (
	"strings"
)

// this is a comment
var morze_tab = map[string]string{
	".-"   : "A",
	"-..." : "B",
	"-.-." : "C",
	"-.."  : "D",
	"."    : "E",
	"..-." : "F",
	"--."  : "G",
	"...." : "H",
	".."   : "I",
	".---" : "J",
	"-.-"  : "K",
	".-.." : "L",
	"--"   : "M",
	"-."   : "N",
	"---"  : "O",
	".--." : "P",
	"--.-" : "Q",
	".-."  : "R",
	"..."  : "S",
	"-"    : "T",
	"..-"  : "U",
	"...-" : "V",
	".--"  : "W",
	"-..-" : "X",
	"-.--" : "Y",
	"--.." : "Z",
}
var final string


func Morze(text string) string{
	keys := make([]string, 0, len(morze_tab))
	for k := range morze_tab {
		keys = append(keys, k)

	}

	spl0 := strings.Split(text, "\n")
	for e := 0; e < len(spl0); e = e + 1 {
		lines(keys, spl0[e], &final)
	}

	return final
}

func lines(keys []string, spl_e string, final *string)  {
	spl := strings.Split(spl_e, "___")
	for e := 0; e < len(spl); e = e + 1 {
		// words
		words(keys, spl[e])
		*final = strings.Join([]string{*final, " "}, "")
	}
	*final += "\n"

}


func words(keys []string, spl_e string)  {
	spl1 := strings.Split(spl_e, "_")
		for i := 0; i < len(spl1); i = i + 1 {
			//word
			word(keys, spl1[i])
		}
}


func word(keys []string, spl1_i string) {
	for j := 0; j < len(keys); j = j + 1 {
		if keys[j] == spl1_i {
			final = strings.Join([]string{final, morze_tab[keys[j]]}, "")
			break
		}
	}
}

func TextToMorze(text string) string{
	var result string
	text = format_text(text)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		for _, r := range line {
			text_to_morze_line(r, &result)
		}
		result += "\n"
	}
	return result
}

func format_text(text string) string {
	text = strings.Map(func (r rune) rune {
		if strings.Index(".,?!:\"'\\/", string(r)) >= 0 {
			return ' '
		}
		return r
	}, text)
	return  strings.ToUpper(text)
}


func text_to_morze_line(r rune, result *string) {
	if r != ' ' {
		for code ,let := range morze_tab {
			if string(r) == let {
				*result += code + "_"
				break
			}
		}
	} else {
		*result += "___"
	}
}