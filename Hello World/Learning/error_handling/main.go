package main

import "fmt"

// func divide(a, b float64) (float64, error)  {
// 	if(b == 0) {
// 		return 0, fmt.Errorf("division value can't be zero")
// 	}
// 	return a / b, nil;
// }

func divide(a, b float64) (float64, string)  {
	if(b == 0) {
		return 0, "division value can't be zero"
	}
	return a / b, "nil";
}

func main()  {
	fmt.Println("Started error handling");

	// ans, err := divide(10, 0);
	// if err != nil {
	// 	fmt.Println("Handling error");
	// }
	ans, _ := divide(10, 0);

	fmt.Println("Divison of two numbers is:", ans);
}