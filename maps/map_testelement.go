package main
import(
	"fmt"
)

func main(){
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["Johannesburg"]=55
	map1["Pretoria"]=20
	map1["Durban"]=25

	value,isPresent = map1["Cape Town"]

	if isPresent{
		fmt.Printf("The value of \"Cape Town\" is: %d ",value)
	}else{
		fmt.Println("map1 doesn't contain Cape Town")
	}

	value,isPresent=map1["Johannesburg"]
	fmt.Printf("Is \"Johnannesburg\" in map1?: %t\n",isPresent)
	fmt.Printf("Value is: %d\n",value)

	//delete an item
	delete(map1,"Johannesburg")
	
	
	value,isPresent=map1["Johannesburg"]
	if isPresent{
		fmt.Printf("The value of \"Johannesburg\" is: %d ",value)
	}else{
		fmt.Println("map1 doesn't contain Johannesburg")
	}
}
