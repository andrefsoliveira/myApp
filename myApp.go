package main

import (
	"fmt"

	sample "github.com/andrefsoliveira/myApp/sampleModule"
)

func main() {

	fmt.Println(sample.SumTwoNumbers(1, 1))

	fmt.Println(sample.CheckName(""))
	fmt.Println(sample.CheckName("André"))
}
