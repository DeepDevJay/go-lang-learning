package main

import "fmt";

func main()  {
	for i:=0; i<3; i++ {
		fmt.Println("Number is:", i)
	}

	counter := 0;
	for {
		fmt.Println("Infinite Loop");
		counter++;
		if counter == 1 {
			break;
		}
	}

	numbers := []int{12, 23, 33, 44, 43};
	for index, value := range numbers {
		fmt.Printf("Index: %d and value %d\n", index, value);
	}

	data := "Hello World!";
	for index, value := range data {
		fmt.Printf("Index: %d and value %c\n", index, value);
	}
}