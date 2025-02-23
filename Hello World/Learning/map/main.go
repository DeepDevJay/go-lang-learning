package main

import "fmt"

func main() {
	studentsGrades := make(map[string]int)

	studentsGrades["Prince"] = 50
	studentsGrades["Yash"] = 60
	studentsGrades["Jay"] = 70
	studentsGrades["Kalu"] = 65
	studentsGrades["Kamal"] = 80

	fmt.Println("Marks of Jay:", studentsGrades["Jay"]);

	studentsGrades["Jay"] = 100;
	fmt.Println("New Marks of Jay:", studentsGrades["Jay"]);

	delete(studentsGrades, "Kalu");
	fmt.Println("Marks of Kalu:", studentsGrades["Kalu"]);

	// checking for existance
	grades, exists := studentsGrades["Don"]
	fmt.Println("Grades of Don:", grades);
	fmt.Println("Don exists:", exists);

	fmt.Println("Marks of Don:", studentsGrades["Don"]);

	for index, value := range studentsGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value);
	}

	person := map[string]int {
		"Jay": 90,
		"Yash": 80,
		"Harsh": 70,
	}

	for index, value := range person {
		fmt.Printf("Person is %s and marks is %d\n", index, value);
	}
}