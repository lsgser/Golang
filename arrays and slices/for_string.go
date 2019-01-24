package main
import(
	"fmt"
)

func main(){
	s := "\u00ff\u754c"
	for i,c:= range s{
		fmt.Printf("%d:%c ",i,c)
	}
	fmt.Println()
	fmt.Println(s)
	
	//immutability of strings,converting a character value in a string
	h:="Hello"
	c:=[]byte(h)
	c[0]='C'
	s2:=string(c)
	fmt.Println(h)
	fmt.Println(s2)
}	
