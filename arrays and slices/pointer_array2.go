package main;
import(
	"fmt"
)

func main(){
	var arr=[3]int{1,2,3}
	var arr2=[3]int{2,4,6}

	for i:=0;i<len(arr);i++{
		mult2(&arr)		
		fp(&arr)
	}
	
	for i:=0;i<len(arr2);i++{
		mult2(&arr2)
		fp(&arr2)
	}
	
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func fp(a *[3]int){fmt.Println(a)}

func mult2(a *[3]int){
	for i:=0;i < len(a);i++{
		a[i]=a[i]*2
	}
}
