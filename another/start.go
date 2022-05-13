package main
import (
	"fmt"
	"./../modal"
)

func main()  {
	fmt.Println("hi, what can you be called")
	var name string
	fmt.Scanf("%s",&name)
	fmt.Println("hi, ",name)
	fmt.Println("what is the temperature in fahrenheit")
	var temp int
	fmt.Scanf("%s",&temp)

	 result := Celsius_converter(temp)
	println(result)
	println(modal.Mahmud)
}

func Celsius_converter(temp int) (int) {
	cels:= (temp -32)*5/9
	
	return cels
}