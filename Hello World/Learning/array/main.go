package main

import "fmt";

func main()  {
	fmt.Println("Array's in Go Lang");
	
	var name[5]string;
	name[0] = "Jaydeep";
	name[1] = "Prajapati";

	fmt.Println("Names of person is:", name);

	var numbers = [5]int{1, 2, 3, 4, 5};
	fmt.Println("Numbers are:", numbers);

	fmt.Println("Length of numbers is:", len(numbers));
	fmt.Println("Value of name on index 2 is:", name[1]);
}