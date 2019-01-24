package main
import(
	"fmt"
)

func main(){
	slice1 := make([]int,4)
	seasons := []string{"Spring","Summer","Autumn","Winter"}

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix,value := range slice1{
		fmt.Printf("Slice at %d is: %d\n",ix,value)
	}
	
	fmt.Println()

	for ix,season := range seasons{
		fmt.Printf("Season %d is %s\n",ix,season)
	}
	
	fmt.Println()
	
	for _,season := range seasons{
		fmt.Printf("%s\n",season)
	}
}
