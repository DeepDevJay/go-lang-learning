package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1");
	if err != nil {
		fmt.Println("Error getting..", err)
		return
	}
	defer res.Body.Close();

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while getting response", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading ...", err);
	// 	return
	// }
	// fmt.Println("Data is:", string(data));

	var todo Todo;
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Todo data is: ", todo)

	fmt.Println("Title: ", todo.Title);
	fmt.Println("Completed:", todo.Completed)
}

func performPostRequest()  {
	todo := Todo {
		UserId: 7,
		Title: "Jaydeep Prajapati",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo);
	if err != nil {
		fmt.Println("Error marshelling..", err)
		return
	}

	// Convert JSON data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString);

	myURL := "https://jsonplaceholder.typicode.com/todos";

	// send POST request
	res, err := http.Post(myURL, "application/json", jsonReader);
	if err != nil {
		fmt.Println("Error while sending request...", err)
		return
	}
	defer res.Body.Close();

	// data, _ := io.ReadAll(res.Body);
	// fmt.Println("Response: ", string(data))

	fmt.Println("Response status: ", res.Status);
}

func performUpdateRequest() {
	todo := Todo {
		UserId: 165,
		Title: "Jaydeep Prajapati GoLang",
		Completed: false,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo);
	if err != nil {
		fmt.Println("Error marshelling..", err)
		return
	}

	// Convert JSON data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString);

	myURL := "https://jsonplaceholder.typicode.com/todos/1";

	// create PUT request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader);
	if err != nil {
		fmt.Println("Error creating PUT request...", err)
		return
	}
	req.Header.Set("Content-Type", "application/json");

	// Send the request
	client := http.Client{}
	res, err := client.Do(req);
	if err != nil {
		fmt.Println("Error sending request...", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body);
	fmt.Println("Response: ", string(data))

	fmt.Println("Response status: ", res.Status);
}

func performDeleteRequest() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1";	

	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("Error sending request...", err)
		return
	}
	req.Header.Set("Content-Type", "application/json");

	// Send the request
	client := http.Client{}
	res, err := client.Do(req);
	if err != nil {
		fmt.Println("Error sending request...", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status: ", res.Status);
}

func main() {
	fmt.Println("Learning CRUD...");
	// performGetRequest();
	// performPostRequest();
	// performUpdateRequest();
	performDeleteRequest()
}