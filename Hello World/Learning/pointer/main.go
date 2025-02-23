package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20;
}

func main() {
	// var num int
	num := 2
	// num := "Jaydeep"

	// var ptr *int
	// ptr = &num

	ptr := &num

	// fmt.Println("Num has value:", num);
	fmt.Println("Pointer contains:", ptr);
	fmt.Println("Data containts through pointer:", *ptr);


	var pointer *int  // default pointer == nil
	if pointer == nil {
		fmt.Println("Pointer is not assigned!");
	}

	value := 10;
	modifyValueByReference(&value);
	fmt.Println("Value contains:", value);
}