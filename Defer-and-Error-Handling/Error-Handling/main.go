package main

import (
	"fmt"
	"os"
)

/**
type PathError struct {
    Op   string
    Path string
    Err  error
}*/

/*
n this tutorial we discussed how to handle errors that occur in our program and also how to inspect the errors to get more information from them. A quick recap of what we discussed in this tutorial,

What are errors?
Error representation
Various ways of extracting more information from errors
Do not ignore errors
*/

func main() {
	f, err1 := os.Open("/test.txt")
	if err1 != nil {
		fmt.Println(err1)
		//	return
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		// why err.Path look PathError struct
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
