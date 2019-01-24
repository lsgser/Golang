package main
import(
	"fmt"
	"./person"
)

func main(){
	p := new(person.Person)
	p.SetFirstName("Lesego")
	p.SetLastName("Seritili")
	fmt.Printf("%s %s\n",p.FirstName(),p.LastName())
}
