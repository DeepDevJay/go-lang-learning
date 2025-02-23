package main

import "fmt"

func main() {

	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 3, 10, 11, 34, 43, 45, 53);
	// fmt.Println("Number :", numbers);
	// fmt.Printf("Number has data type : %T\n", numbers);
	// fmt.Println("Length:", len(numbers));

	names := []string{};
	fmt.Println("Slice:", names);
	fmt.Println("Length:", len(names));
	fmt.Println("Capacity:", cap(names));

	numbers := make([]int, 3, 5);
	numbers = append(numbers, 4);
	numbers = append(numbers, 9);
	// slice will double it's size goes over the capacity
	numbers = append(numbers, 20);
	numbers = append(numbers, 22);

	fmt.Println("Slice:", numbers);
	fmt.Println("Length:", len(numbers));
	fmt.Println("Capacity:", cap(numbers));

	
}
