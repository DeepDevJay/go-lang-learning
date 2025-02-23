package main

import "fmt"

func main() {
	age := 25;
	name := "Jay";
	height := 5.215461;

	fmt.Println("Age:", age,"Name:", name, "Height:", height);
	fmt.Println("Hello World");

	fmt.Printf("Age is %d\n", age);
	fmt.Printf("Name is %s\n", name);
	fmt.Printf("Height is %f\n", height);

	fmt.Printf("Name: %s, Age: %d, Height: %.2f", name, age, height);
}