package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 40
	fmt.Println("Number is:", num);
	fmt.Printf("Type of num is: %T\n", num)

	// num = num + 1.25 // error for data conversion due to int to float

	var data float64 = float64(num);
	data = data + 1.25
	fmt.Println("Data is:", data);
	fmt.Printf("Type of data is: %T\n", data);

	num = 123;
	str := strconv.Itoa(num);
	fmt.Println("Number is:", str);
	fmt.Printf("Type of num is: %T\n", str)

	num_string := "465";
	num_int, _ := strconv.Atoi(num_string);
	num_int = num_int + 50;
	fmt.Println("Number is:", num_int);
	fmt.Printf("Type of num is: %T\n", num_int)

	num_string = "3.25";
	num_float, _ := strconv.ParseFloat(num_string, 64);
	fmt.Println("Number is:", num_float);
	fmt.Printf("Type of num is: %T\n", num_float)
}