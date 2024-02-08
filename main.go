package main

import "fmt"

func main() {

	for {
		var username string
		var password string

		fmt.Println("Username: ")
		fmt.Scan(&username)

		fmt.Println("Password: ")
		fmt.Scan(&password)

		if username == "maxkaiser" && password == "max98" {
			fmt.Println("Succsefully Logged In!")
			break
		} else {
			fmt.Println("Wrong Username or Password!, Please try again.")
		}
	}

	fmt.Println("Which calculation would you like to do?(+,-,/,*)")
	var calculationInput string
	fmt.Scan(&calculationInput)

	firstNumber := "Please enter the first number."
	secondNumber := "Please enter the second number"

	switch calculationInput {
	case "+":
		fmt.Println(firstNumber)
		var add1 int64
		fmt.Scan(&add1)
		fmt.Println(secondNumber)
		var add2 int64
		fmt.Scan(&add2)

		result := addition(add1, add2)
		fmt.Printf("The sum of %v + %v is: %v", add1, add2, result)
	case "-":
		fmt.Println(firstNumber)
		var sub1 int64
		fmt.Scan(&sub1)

		fmt.Println(secondNumber)
		var sub2 int64
		fmt.Scan(&sub2)

		result := subtraction(sub1, sub2)
		fmt.Printf("The sum of %v - %v is: %v", sub1, sub2, result)
	case "*":
		fmt.Println(firstNumber)
		var multi1 float64
		fmt.Scan(&multi1)

		fmt.Println(secondNumber)
		var multi2 float64
		fmt.Scan(&multi2)

		result := multiplication(multi1, multi2)
		fmt.Printf("The sum of %v * %v is: %v", multi1, multi2, result)
	default:

		fmt.Println(firstNumber)
		var div1 float64
		fmt.Scan(&div1)

		fmt.Println(secondNumber)
		var div2 float64
		fmt.Scan(&div2)

		result := division(div1, div2)
		fmt.Printf("The sum of %v / %v is: %v", div1, div2, result)
	}
}

func addition(n1, n2 int64) int64 {
	sum := n1 + n2
	return sum
}

func subtraction(n1, n2 int64) int64 {
	sum := n1 - n2
	return sum
}

func multiplication(n1, n2 float64) float64 {
	sum := n1 * n2
	return sum
}

func division(n1, n2 float64) float64 {
	sum := n1 / n2
	return sum
}
