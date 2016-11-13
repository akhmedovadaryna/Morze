package io

import (
	"os"
	"bufio"
	"io/ioutil"
)

func Input_from_file(file_name string, c chan string) {
	file, err := os.Open(file_name)
    if err != nil {
        return
    }
    defer file.Close()



	f := bufio.NewReader(file)
    	for  {
        read_line, err := f.ReadString('\n')
		if err != nil {
			c <- "//"
			break
		}
	c <- read_line
    }

	return

/*

	bs, err := ioutil.ReadFile(file_name)
    	if err != nil {
        	return "", err
    	}
    	str := string(bs)
  //  	fmt.Println(str)
	return str, err
	*/
}


func Output_from_file(filename string, c chan string, finish *bool) {
	file, err := os.Create(filename)
	if err != nil {
       		return
    	}
    	defer file.Close()
	for {
		str := <- c
		if str == "//"{
			*finish = true
			break
		}
		file.WriteString(str)
	}
	return

}


func Serial_Input_from_file(file_name string) (string, error){

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


func Serial_Output_from_file(filename string, result string) error{

	file, err := os.Create(filename)
	if err != nil {
       		return err
    	}
    	defer file.Close()
   	file.WriteString(result)
	return err

}

