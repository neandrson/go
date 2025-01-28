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

	//client := http.Client{}
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
					return nil, http.St
				default:
					req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
					if err != nil {
						return nil, http.Error("failed to create request with ctx:", err)
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

		//http.HandleFunc("/long", longHanlder)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("error when starting a server\n")
		}
	}()

	name := "ok"
	url := "http://localhost:8080/hello"
	timeout := 2 * time.Second
	want := APIResponse{
				Data:       `Hello, World!`,
				StatusCode: http.StatusOK,
			}
	/*cases := []struct {
		timeout time.Duration
		url     string
		want    APIResponse
		wantErr error
		name    string
	}{
		{
			name:    "ok",
			url:     "http://localhost:8080/hello",
			timeout: 2 * time.Second,
			want: APIResponse{
				Data:       `Hello, World!`,
				StatusCode: http.StatusOK,
			},
		},
		{
			name:    "timeout",
			url:     "http://localhost:8080/long",
			timeout: 10 * time.Millisecond,
			want: APIResponse{
				Data:    ``,
				StatusCode: http.StatusCode,
			},
		},
	}*/
	//client := &http.Client{}
    req, err := http.NewRequest("GET", url, want[Data]) 
    // добавляем заголовки
    //req.Header.Add("Accept", "text/html")   // добавляем заголовок Accept
    //req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent
  
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
}
