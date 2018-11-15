package main

import "fmt"

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharsOld(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

//func mutate(s string) string {
func mutate(s []rune) string {
	//	s[0] = "a" //any valid unicode character within single quote is a rune
	s[0] = 'a'
	//return s
	return string(s)
}

func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Println("")
	printChars(name)
	fmt.Println("")

	name = "Señor"
	printBytes(name)
	fmt.Println("")
	printChars(name)

	fmt.Println("")
	printCharsAndBytes(name)

	fmt.Println("")
	//byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice := []byte{67, 97, 102, 195, 169}
	str1 := string(byteSlice)
	fmt.Println(str1)

	fmt.Println("")
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)

	fmt.Println("")
	h := "hello"
	//fmt.Println(mutate(h))

	fmt.Println(mutate([]rune(h)))
}
