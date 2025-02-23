package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON in Go Lang")
	person := Person{Name: "Jaydeep", Age: 22, IsAdult: true}
	// fmt.Println("Person data is: ", person);

	// convert Persion into JSON Encoding(Marshalling in go lang)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshalling!")
		return
	}
	fmt.Println("Person data is: ", string(jsonData));

	// Decoding (UnMarshalling)
	var personData Person;
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error in Unmarshalling")
		return
	}
	fmt.Println("Person data is: ", personData)
}