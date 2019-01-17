package main
import(
	"fmt"
)
/*
	code prints out 10-1 in recursively
*/
func main(){
	printrec(10)
}

func printrec(n int){
	if n >= 1{
		fmt.Println("",n)
		printrec(n-1)
	}
}
