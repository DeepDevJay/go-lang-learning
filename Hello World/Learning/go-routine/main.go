package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello,");
	// time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHello function ended successfully");
}

func sayHi() {
	fmt.Println("Hii, Jaydeep :)");
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("sayHi function ended successfully");
}

func main() {
	fmt.Println("Learning GoRoutine ...");

	go sayHello();
	go sayHi();

	// wait for a moment to allow the goroutines to finish
	time.Sleep(800 * time.Millisecond)
}