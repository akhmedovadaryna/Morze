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
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	"-----": "0",
}
var final string


func Morze(text string) string{

	spl0 := strings.Split(text, "\n")
	for e := 0; e < len(spl0); e = e + 1 {
		lines(spl0[e], &final)
	}

	return final
}

func lines(spl_e string, final *string)  {
	spl := strings.Split(spl_e, "___")
	for e := 0; e < len(spl); e = e + 1 {

		// words
		words(spl[e])
		*final = strings.Join([]string{*final, " "}, "")
	}
	*final += "\n"

}


func words(spl_e string)  {
	spl1 := strings.Split(spl_e, "_")
		for i := 0; i < len(spl1); i = i + 1 {
			//word
			word(spl1[i])
		}
}


func word(spl1_i string) {
		for i, j := range morze_tab{
			if i == spl1_i {
				final = strings.Join([]string{final, j}, "")
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