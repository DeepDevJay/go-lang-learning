package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done(); // Signal that this goroutine is done
	fmt.Printf("Worker %d started\n", i);
	// some work happening
	fmt.Printf("Worker %d completed\n", i);
}

func main() {
	// fmt.Println("Learnig Sync WaitGroup");

	var wg sync.WaitGroup;
	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1); // Increment the WaitGroup counter
		go worker(i, &wg);
	}

	// Wait for all worker to finish
	wg.Wait();
	fmt.Println("Worker task complete.")
}