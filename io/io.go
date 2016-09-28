package io

import (


	"io/ioutil"
)

func Input_from_file(file_name string) (string, error){

/*
	file, err := os.Open("./files/input")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

	stat, err := file.Stat()
    	if err != nil {
        	return
    	}
	fmt.Println(stat.Size())

	f := bufio.NewReader(file)
    	for  {
        read_line, err := f.ReadString('\n')
		if err != nil {
			break
		}
        fmt.Print(read_line)
    }
*/

	bs, err := ioutil.ReadFile("./files/input_text")
    	if err != nil {
        	return "", err
    	}
    	str := string(bs)
  //  	fmt.Println(str)
	return str, err
}
