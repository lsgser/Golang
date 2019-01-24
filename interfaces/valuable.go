package main
import(
	"fmt"
)

type StockPosition struct{
	ticker string
	sharePrice float32
	count float32
}

type Car struct{
	make string
	model string
	price float32
}

type valuable interface{
	getValue() float32
}

func main(){
	var o valuable = StockPosition{"GOOG",500.2,4.0}
	showValue(o)
	o = Car{"FORD","RANGER",44000.00}
	showValue(o)	
}

func (s StockPosition) getValue() float32{
	return s.sharePrice * s.count
}

func (c Car) getValue() float32{
	return c.price
}

func showValue(asset valuable){
	fmt.Printf("Value of the asset is %0.2f\n",asset.getValue())
}
