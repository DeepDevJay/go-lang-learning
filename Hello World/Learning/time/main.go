package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now();
	fmt.Println("Current time is:", currentTime);
	fmt.Printf("Current time Type is: %T\n", currentTime);

	// formatted := currentTime.Format("dd-mm-yyyy, hh-mm-ss"); // doesn't work in go-lang due to fix format
	// formatted := currentTime.Format("02-01-2006, Monday 15:04:05");
	formatted := currentTime.Format("2006/01/02, Monday 3:04 PM");
	fmt.Println("Formatted time is:", formatted);

	layout_str := "2006/01/02";
	dateStr := "2024/11/25";
	formattedTime, _ := time.Parse(layout_str, dateStr);
	fmt.Println("FormattedTime is:", formattedTime)

	// add 1 more day in current time
	new_date := currentTime.Add(24 * time.Hour);
	fmt.Println("New date time is:", new_date);
	formatted_new_date := new_date.Format("2006/01/02, Monday 15:04:05");
	fmt.Println("Formatted new date is:", formatted_new_date);
}