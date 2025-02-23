package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learnig web request");
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1");
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close();
	fmt.Printf("Type of res is: %T\n", res);
	// fmt.Println("Response: ", res);

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while reading response", err)
		return
	}
	fmt.Println("Response: ", string(data));
}