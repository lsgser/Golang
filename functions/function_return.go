package main;
import(
	"fmt"
)

func main(){
	p2 := Add2()//returns an anonymous function(closure i.e a function returned by another function)
	fmt.Printf("Call Add2 for 3 gives: %v\n",p2(3))
	TwoAdder := Adder(2)//returns an anonymous function(closure i.e a function returned by another function),but the main function has an input value
	fmt.Printf("The result is: %v\n",TwoAdder(3))
}

func Add2() (func(b int) int){
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int)(func(b int) int){
	return func(b int) int{
		return a + b
	}
}


