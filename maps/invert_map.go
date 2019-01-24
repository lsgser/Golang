package main
import(
	"fmt"
)

/*
	By this we mean switching the values and keys. If the value type of a map is acceptable as a key
	type, and the map values are unique, this can be done easily
*/
var(
	barVal = map[string]int{
		"alpha":34,"bravo":56,"charlie":23,"delta":87,
		"echo":56,"foxtrot":12,"golf":34,"hotel":16,
		"indio":87,"juliet":65,"kilo":43,"lima":98,
	}
)

func main(){
	invMap:=make(map[int]string,len(barVal))
	
	for key,value:=range barVal{
		invMap[value]=key
	}

	fmt.Println("Inverted:")
	for key,value:=range invMap{
		fmt.Printf("The key %d contains the value: %s\n",key,value)
	}	
}
