package main

import (
	"./logic"
	"./io"
	"fmt"
)

func main() {
	//var text2 string = ".-___.-_.-_.-___.-"
	var name_file = "./files/input_text"
	text, err := io.Input_from_file(name_file)
	if err != nil {
        	return
    	}
	result := Morze.TextToMorze(text)
	fmt.Println(result)


	var filename = "./files/output.txt"
	err = nil
	err = io.Output_from_file(filename,result)
	if err != nil {
		fmt.Println("error")
		return
	}

	result = Morze.Morze(result)
	fmt.Println(result)

}