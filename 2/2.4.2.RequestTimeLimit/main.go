package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	//URL        string
	Data       string
	StatusCode int
	//Err        error
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(url))

	client := http.Client{}
	answs := make([]*APIResponse, len(url))

	for i, url_ := range url {
		go func(idx int, url string) {
			defer wg.Done()
			answ := &APIResponse{}
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			for {
				select {
				case <-ctx.Done():
					return nil, fmt.Printf("failed to create request with ctx: %w", ctx)
				default:
					req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://localhost:8080/", nil)
					if err != nil {
						return nil, http.Errorf("failed to create request with ctx: %w", err)
					}

					//res, err := http.DefaultClient.Do(req)
					res, err := io.Reader(req)
					if err != nil {
						return nil, fmt.Errorf("failed to perform http request: %w", err)
					}

					return res, nil
					defer resp.Body.Close()

					answ.StatusCode = resp.StatusCode

					buf, err := io.ReadAll(resp.Body)
					if err != nil {
						answ.Err = err
						mu.Lock()
						answs[idx] = &answ
						mu.Unlock()
						return
					}

					answ.Data = string(buf)
					mu.Lock()
					answs[idx] = &answ
					mu.Unlock()
				}
			}
		}(i, url_)
	}

	wg.Wait()

	return answs
}

func main() {
	go func() {
		http.HandleFunc("/hello", helloHandler)

		http.HandleFunc("/long", longHanlder)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("error when starting a server\n")
		}
	}()
}
