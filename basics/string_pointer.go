package main

import(
	"fmt"
)

func main(){
	s := "good bye"
	var p *string = &s
	//assigning a new value in *p also changes the value of s
	*p = "Ciao"

	fmt.Printf("Here is a the pointer p: %p\n",p)
	fmt.Printf("Here is the string *p: %s\n",*p)
	fmt.Printf("Here is the string s: %s\n",s)
}
