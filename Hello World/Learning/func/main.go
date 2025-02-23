package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function");
}

func add(a, b int) int {
	return a + b;
}

func multiply(a, b int) (result int)  {
	result = a * b;
	return;
}

func main() {
	fmt.Println("We are learning functions in GoLang");
	simpleFunction();

	sum := add(10, 20);
	fmt.Println("Addition of Two numbers is:", sum);

	data := multiply(5, 10);
	fmt.Println("Multiplication of Two numbers is:", data);
}