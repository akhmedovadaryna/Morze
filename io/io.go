package io

import (


	"io/ioutil"
	"os"
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



	bs, err := ioutil.ReadFile(file_name)
    	if err != nil {
        	return "", err
    	}
    	str := string(bs)
  //  	fmt.Println(str)
	return str, err
}


func Output_from_file(filename string, result string) error{

	file, err := os.Create(filename)
	if err != nil {
       		return err
    	}
    	defer file.Close()
   	file.WriteString(result)
	return err

}
