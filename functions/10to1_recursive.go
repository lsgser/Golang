package main
import(
	"fmt"
)

func main(){
	printrec(10)
}

func printrec(n int){
	if n >= 1{
		fmt.Println("",n)
		printrec(n-1)
	}
}
