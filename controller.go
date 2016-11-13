package main

import (
	"./logic"
	"./io"
//	"time"
//	"fmt"
	"flag"

	"fmt"
	"time"
)

const serial bool = false

func main() {

	m := flag.String("morze", "", "morze to text")
	t := flag.String("text", "", "text to morze")
	flag.Parse()

	if serial {
		if *t != "" {

			var name_file = "./files/"
			name_file += *t
			fmt.Println(*t)
			text, err := io.Serial_Input_from_file(name_file)
			if err != nil {
				return
			}
			result := logic.Serial_TextToMorze(text)

			var filename = "./files/output.txt"
			err = nil
			err = io.Serial_Output_from_file(filename, result)
			if err != nil {
				fmt.Println("error")
				return
			}
		}
		if *m != "" {
			tt := time.Now().Nanosecond()
			var name_file = "./files/"
			name_file += *m
			fmt.Println(*m)
			text, err := io.Serial_Input_from_file(name_file)
			if err != nil {
				return
			}
			result := logic.Serial_Morze(text)


			var filename = "./files/output.txt"
			err = nil
			err = io.Serial_Output_from_file(filename, result)
			if err != nil {
				fmt.Println("error")
				return
			}
			fmt.Println(time.Now().Nanosecond() - tt)
		}

	} else {
		c_in := make(chan string);
		c_out := make(chan string);

		if *t != "" {
			var finish bool = false
			var name_file = "./files/"
			name_file += *t
			go io.Input_from_file(name_file, c_in)
			go logic.TextToMorze(c_in, c_out)

			var filename = "./files/output_"
			filename += *t
			go io.Output_from_file(filename, c_out, &finish)
			for !finish {}

		}
		if *m != "" {
			tt := time.Now().Nanosecond()
			var finish bool = false
			var name_file = "./files/"
			name_file += *m
			go io.Input_from_file(name_file, c_in)

			go logic.Morze(c_in, c_out)


			var filename = "./files/output_"
			filename += *m
			go io.Output_from_file(filename, c_out, &finish)
			for !finish {}
			fmt.Println(time.Now().Nanosecond() - tt)
		}

	}

}