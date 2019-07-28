package main
import (
	"fmt"
)

func main() {
	fmt.Println("Addition is ",add(42, 13))
	fmt.Println("Subtraction is ",sub(42, 13))
	fmt.Println("Multiplication is ",mul(42, 13))
	fmt.Println("Division is ",div(42, 13))
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x - y
}

func div(x, y int) int {
	return x - y
}


