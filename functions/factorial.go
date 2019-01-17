package main;
import(
	"fmt"
)

func main(){
	result := 0;

	for i := 0;i<=29;i++{
		result = factorial(i);
		fmt.Printf("%d!\n",result);
	}
}

func factorial(n int)(res int){
	if n <= 1{
		res = 1
	}else
	{
		res = n*factorial(n-1);
	}
	
	return
}
