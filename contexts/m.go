// What is parallelism?
// You can run multiple actions at the same time. Imagine a market with several cash registers.
// Parallelism allows you to serve multiple customers at the same time, one customer per register.

// What is concurrency?
// Concurrency is about dealing with multiple tasks at once, but not necessarily simultaneously.
// Imagine a single cashier who can handle multiple customers by switching between them, attending to one customer for a moment, then another, and so on.
// The cashier is not serving multiple customers at the exact same time, but is making progress on multiple tasks by managing them concurrently.

// If your code runs concurrently, it is automatically running in parallel.

// When dealing with concurrency, we need to worry about race conditions.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sequentiallyCallGoogle() {
	start := time.Now()
	for range 10 {
		res, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		// We need to defer the response body Close() to ensure that the response body is properly closed
		// after we're done with it, even if an error occurs later in the function. This prevents
		// resource leaks and allows the connection to be reused for future requests.
		defer res.Body.Close()
		fmt.Println("ok from sequential approach!")
	}
	fmt.Println(time.Since(start))
}

func concurrentlyCallGoogle() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	for range n {
		go func() {
			defer wg.Done()

			res, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			fmt.Println("ok from concurrent approach!")
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func main() {
	sequentiallyCallGoogle()
	concurrentlyCallGoogle()
}
