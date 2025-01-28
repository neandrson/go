package main

import (
	"context"
	"fmt"
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
	var wg sync.WaitGroup
	var mu sync.Mutex
	var APRS []*APIResponse
	for _, i := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			nwCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			buffer := make([]byte, 4096)
			req, err := http.NewRequestWithContext(nwCtx, "GET", u, nil)
			var client = http.Client{}
			if err != nil {
				mu.Lock()
				defer mu.Unlock()
				APRS = append(APRS, &APIResponse{URL: u, Err: err})
				//mu.Unlock()
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				defer mu.Unlock()
				APRS = append(APRS, &APIResponse{URL: u, Data: "", StatusCode: resp.StatusCode, Err: err})
				//mu.Unlock()
				return
			}
			defer resp.Body.Close()
			n, err := resp.Body.Read(buffer)
			mu.Lock()
			APRS = append(APRS, &APIResponse{URL: u, Data: string(buffer[:n]), StatusCode: resp.StatusCode, Err: nil})
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return APRS
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
	timeout := 2 * time.Second

	client := &http.Client{}
	time.Sleep(500 * time.Millisecond)
	go FetchAPI(ctx, urls, timeout)
	req, err := http.NewRequest("GET", url, nil)
	// добавляем заголовки
	//req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
	//req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("resp, %w", resp)
	}
	defer resp.Body.Close()
}
