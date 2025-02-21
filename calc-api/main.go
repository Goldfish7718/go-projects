package main

import (
	"calc-api/utils"
	"fmt"
)

func main () {
	var input1 int;
	var input2 int;

	fmt.Println("Welcome to calculator API")

	fmt.Println("Enter 1st number")
	fmt.Scanln(&input1)

	fmt.Println("Enter 1st number")
	fmt.Scanln(&input2)

	sum := utils.Add(input1, input2)
	fmt.Println(sum)
}