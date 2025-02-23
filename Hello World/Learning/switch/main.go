package main

import "fmt"

func main() {
	
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday");
	case 2:
		fmt.Println("Tuesday");
	case 3:
		fmt.Println("Wednesday");
	default:
			fmt.Println("Unknown day");
	}

	month := "August";
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter");
	case "April", "May", "June":
		fmt.Println("Spring");
	default:
		fmt.Println("Other Season");
	}

	temprature := 25;
	switch {
	case temprature < 0:
		fmt.Println("Freezing");
	case temprature >= 0 && temprature < 10:
		fmt.Println("Cold");
	case temprature >= 10 && temprature < 20:
		fmt.Println("Cool");
	case temprature >= 20 && temprature < 30:
		fmt.Println("Warm");
	default:
		fmt.Println("Hot");
	}
}