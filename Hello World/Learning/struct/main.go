package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	HouseNo int
	Area string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var jay Person
	// fmt.Println("Jay Person :", jay)

	jay.FirstName = "Jay";
	jay.LastName = "Patel";
	jay.Age = 22;
	// fmt.Println("Person Jay :", jay);

	yash := Person {
		FirstName: "Yash",
		LastName: "Vaderiya",
		Age: 23,
	}
	fmt.Println("Person Yash:", yash);


	// new keyword
	var Amit = new(Person);
	Amit.FirstName = "Amit";
	Amit.LastName = "Patel";
	Amit.Age = 23;

	// fmt.Println("Person Amit:", Amit);

	employee1 := Employee {
		Person_Details: Person{
			FirstName: "Jaydeep",
			LastName: "Prajapati",
			Age: 22,
		},
		Person_Contact: Contact{
			Email: "jaydeep@gmail.com",
			Phone: "9283848474",
		},
		Person_Address: Address{
			HouseNo: 121,
			Area: "Adajan, Surat",
			State: "Gujarat",
		},
	}

	fmt.Println("Employee 1 :", employee1);
	fmt.Println("Employee Details :", employee1.Person_Details);
	fmt.Println("Employee Contact :", employee1.Person_Contact);
	fmt.Println("Employee Address :", employee1.Person_Address);
}