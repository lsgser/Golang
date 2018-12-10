package main

import(
	"fmt"
)

type TZ int

func main(){
	var a,b TZ =3,4 //TZ stands for an int type here
	c := a+b
	fmt.Printf("c has a value: %d\n",c)
}
