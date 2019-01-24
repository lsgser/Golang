package main
import(
	"fmt"
)

type B struct{
	thing int
}

func main(){
	var b1 B //b1 is a value
	b1.Change()
	fmt.Println(b1.Write())
	
	b2 := new(B)//b2 is a pointer
	b2.Change()
	fmt.Println(b2.Write())
}

func (b *B) Change(){
	b.thing=1
}

func (b B) Write() string{
	return fmt.Sprint(b)
}
