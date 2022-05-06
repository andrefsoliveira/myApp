package sampleModule

import (
	"errors"
	"fmt"
)

func SumTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func CheckName(name string) (string, error) {

	if name == "" {
		return "", errors.New("The name cannot be empty")
	}

	return fmt.Sprintf("The name is %v", name), nil
}
