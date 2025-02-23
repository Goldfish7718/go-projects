package main

import (
	"calc-api/utils"
	"fmt"
)

func main () {
	var input1 int;
	var input2 int;
	var ch int;

	
	fmt.Println("Welcome to calculator API")
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide")
	fmt.Scanln(&ch)

	fmt.Println("Enter 1st number")
	fmt.Scanln(&input1)

	fmt.Println("Enter 2nd number")
	fmt.Scanln(&input2)

	switch ch {
	case 1:
		fmt.Println(utils.Add(input1, input2))
			
	case 2:
		fmt.Println(utils.Subtract(input1, input2))

	case 3:
		fmt.Println(utils.Multiply(input1, input2))

	case 4:
		fmt.Println(utils.Divide(input1, input2))
	}
}