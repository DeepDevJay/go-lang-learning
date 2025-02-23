package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",");
	fmt.Println(parts)

	str := "one two three four two two five";
	count := strings.Count(str, "two");
	fmt.Println("Count of two is:", count)

	str = "            Hello, Go!   ";
	trimmed := strings.TrimSpace(str);
	fmt.Println("trimmed: ", trimmed);

	str1 := "Jaydeep";
	str2 := "Prajapati";
	result := strings.Join([]string{str1, "Dipakbhai", str2}, " ");
	fmt.Println("Result is: ", result);

}