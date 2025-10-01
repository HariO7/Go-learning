package main

import "fmt"

func main() {
	transformFuncDouble := createTransformFunc(2)

	fmt.Println(transformFuncDouble(4))
}

func createTransformFunc(factor int) func(number int) int {
	return func(number int) int {
		return number * factor
	}
}
