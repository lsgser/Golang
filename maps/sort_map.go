package main
import(
	"fmt"
	"sort"
)
/*
	If you want a sorted map, copy the keys (or values) to a slice, sort the slice (using the sort package), and then print out the keys and/or values using a for-range on the slice.
*/
var(
	barVal = map[string]int{
		"alpha":34,"bravo":56,"charlie":23,"delta":87,"echo":56,"foxtrot":12,"golf":34,"hotel":16,"indio":87,"juliet":65,"kilo":43,"lima":98}
)

func main(){
	fmt.Println("Unsorted:")
	for k,v:=range barVal{
		fmt.Printf("Key: %v,Value: %v /",k,v)
	}

	keys := make([]string,len(barVal))

	i:=0
	for k,_:=range barVal{
		keys[i]=k
		i++
	}

	sort.Strings(keys)
	fmt.Println()
	fmt.Println("Sorted:")
	for _,k:= range keys{
		fmt.Printf("Key: %v, Value: %v /",k,barVal[k])
	}

	
}
