package main

import (
	"fmt"

	sample "github.com/andrefsoliveira/myApp/sampleModule"
)

func main() {

	fmt.Println(sample.SumTwoNumbers(1, 1))

	sample.CheckName("")
	sample.CheckName("André")
}
