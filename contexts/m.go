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
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Println("ok from server")
	}))
	defer server.Close()

	for range n {
		go func(ctx context.Context) {
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil)
			if err != nil {
				panic(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("timeout")
					return
				}
				panic(err)
			}
			defer res.Body.Close()
			fmt.Println("ok from concurrent approach!")
		}(ctx)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func main() {
	// sequentiallyCallGoogle()
	concurrentlyCallGoogle()
}
