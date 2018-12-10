package main

import(
	"fmt"
	"unicode/utf8"
)

func main(){
	//Haven't calculated the bytes yet,only did the length
	var str string = "asSASA ddd dsjkdsjs dk"
	var uniStr string = "asSASA ddd dsjkdsjsこん dk"
	var length int = len(str)
	var uniLength int = utf8.RuneCountInString(uniStr)
	fmt.Printf("The length and byte of the string is %d",length)
	fmt.Println()
	fmt.Printf("The length and byte of the unicode string is %d",uniLength)
	fmt.Println()
}
