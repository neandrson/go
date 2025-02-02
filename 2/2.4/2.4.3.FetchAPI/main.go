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
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Microsecond * 100)
	fmt.Fprintf(w, "Hi, World!")
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(urls))

	client := http.Client{}
	answs := make([]*APIResponse, len(urls))

	for i, url := range urls {
		go func(idx int, url string) {
			defer wg.Done()
			answ := APIResponse{URL: url}
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				answ.Err = err
				mu.Lock()
				answs[idx] = &answ
				mu.Unlock()
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				answ.Err = err
				mu.Lock()
				answs[idx] = &answ
				mu.Unlock()
				return
			}
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
		}(i, url)
	}

	wg.Wait()

	return answs
}

func main() {
	ctx := context.Background()
	go func() {
		http.HandleFunc("/hello", helloHandler)

		http.HandleFunc("/long", longHanlder)

		http.HandleFunc("/hi", hiHandler)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("error when starting a server\n")
		}
	}()

	urls := []string{"http://localhost:8080/hello", "http://localhost:8080/hi", "http://localhost:8080/long"}
	timeout := 1 * time.Second

	client := &http.Client{}
	time.Sleep(500 * time.Millisecond)

	go FetchAPI(ctx, urls, timeout)

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		// добавляем заголовки
		//req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
		//req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("resp, %w\n", resp)
		}
		defer resp.Body.Close()
	}
}
