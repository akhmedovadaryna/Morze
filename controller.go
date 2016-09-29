package main

import (
	"./logic"
	"./io"
	"fmt"
	"flag"
)

func main() {

	m := flag.String("morze", "", "morze to text")
	t := flag.String("text", "", "text to morze")
	flag.Parse()


	if *t != "" {

		var name_file = "./files/"
		name_file += *t
		fmt.Println(*t)
		text, err := io.Input_from_file(name_file)
		if err != nil {
			return
		}
		result := logic.TextToMorze(text)
		fmt.Println(result)

		var filename = "./files/output.txt"
		err = nil
		err = io.Output_from_file(filename, result)
		if err != nil {
			fmt.Println("error")
			return
		}
	}
	if *m != "" {

		var name_file = "./files/"
		name_file += *m
		fmt.Println(*m)
		text, err := io.Input_from_file(name_file)
		if err != nil {
			return
		}
		result := logic.Morze(text)
		fmt.Println(result)

		var filename = "./files/output.txt"
		err = nil
		err = io.Output_from_file(filename, result)
		if err != nil {
			fmt.Println("error")
			return
		}

	}


}