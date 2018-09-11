package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	data1, err1 := ioutil.ReadFile("test.txt")
	if err1 != nil {
		fmt.Println("File reading error", err1)
		return
	}
	fmt.Println("Contents of file:", string(data1))

	/* 	fptr1 := flag.String("fpath", "test.txt", "file path to read from1")
	   	flag.Parse()
	   	fmt.Println("value of fpath is", *fptr1) */

	fptr := flag.String("fpath", "test.txt", "file path to read from2")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
