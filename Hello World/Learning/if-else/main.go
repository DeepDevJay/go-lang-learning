package main

import "fmt"

func main() {
	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10");
	} else if x > 5 {
		fmt.Println("x is greater than 5 but less than 10");
	} else {
		fmt.Println("x is smaller than 5");
	}

	y := 10;
	z := 20;
	if x > 5 && (y > 5 || z < 20) {
		fmt.Println("How are you ?");
	} else {
		fmt.Println("You need to learn programming!")
	}
}