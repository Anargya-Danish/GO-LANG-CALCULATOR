// MADE BY ANARGYA-DANISH
// GITHUB: https://github.com/Anargya-Danish
// YOU CAN DEVELOP THIS PROJECT TO BE BETTER

package main

import "fmt"

func increase() {
	//variable
	var num1 int
	var num2 int

	fmt.Print("ENTER THE FIRST NUMBER: ")
	fmt.Scanln(&num1)

	fmt.Print("ENTER THE SECOND NUMBER: ")
	fmt.Scanln(&num2)

	total := num1 + num2

	fmt.Println(num1, "+", num2, "=", total)

}

func subTraction() {
	//variable
	var num1 int
	var num2 int

	fmt.Print("ENTER THE FIRST NUMBER: ")
	fmt.Scanln(&num1)

	fmt.Print("ENTER THE SECOND NUMBER: ")
	fmt.Scanln(&num2)

	total := num1 - num2

	fmt.Println(num1, "-", num2, "=", total)

}

func division() {
	//variable
	var num1 int
	var num2 int

	fmt.Print("ENTER THE FIRST NUMBER: ")
	fmt.Scanln(&num1)

	fmt.Print("ENTER THE SECOND NUMBER: ")
	fmt.Scanln(&num2)

	total := num1 / num2

	fmt.Println(num1, "/", num2, "=", total)

}

func multiplication() {
	//variable
	var num1 int
	var num2 int

	fmt.Print("ENTER THE FIRST NUMBER: ")
	fmt.Scanln(&num1)

	fmt.Print("ENTER THE SECOND NUMBER: ")
	fmt.Scanln(&num2)

	total := num1 * num2

	fmt.Println(num1, "x", num2, "=", total)

}

func main() {
	fmt.Println("===WELCOME TO ANARGYA-DANISWARA CALCULATOR===")
	fmt.Println("YOU CAN DEVELOP THIS CALCULATOR TO BE BETTER AND BETTER : )")
	var Operator string
	fmt.Println("Select Your Operator\n==> + \n==> -\n==> x\n==> / ")
	fmt.Print(": ")
	fmt.Scanln(&Operator)

	switch Operator {
	case "+":
		increase()
	case "-":
		subTraction()
	case "/":
		division()
	case "x":
		multiplication()
	default:
		fmt.Println("ERROR, Not a math operator")
	}

}
