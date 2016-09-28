package main

import (
	"./logic"
	"./io"
	"fmt"
)

func main() {
	//var text2 string = ".-___.-_.-_.-___.-"
	var name_file = "input.txt"
	text, err := io.Input_from_file(name_file)
	if err != nil {
        	return
    	}
	result := Morze.TextToMorze(text)
	fmt.Println(result)
	result = Morze.Morze(result)
	fmt.Println(result)

}