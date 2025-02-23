package main

import "fmt"

// defer will execute before completion of execution of main function;
// when there are multiple defer keyword it will take them in LIFO order format

func add(a,b int) int {
	return a + b;
}

func main() {
	fmt.Println("Starting of the program");
	data := add(1,2);
	defer fmt.Println("Addition data is:", data);
	defer fmt.Println("Middle of the program");
	fmt.Println("End of the program");
}