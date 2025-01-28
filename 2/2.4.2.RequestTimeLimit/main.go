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
	wg.Add(1)

	client := http.Client{}
	answs := make([]*APIResponse, 1)

	i := 1
	//for i, url := range url {
	go func(idx int, url string) {
		defer wg.Done()
		//answ := url // APIResponse{URL: url}
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			//answ.Err = err
			mu.Lock()
			answs[idx] = &answ

			mu.Unlock()
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			//answ.Err = err
			mu.Lock()
			//answs[idx] = &answ

			mu.Unlock()
			return
		}
		defer resp.Body.Close()

		//answ.StatusCode = resp.StatusCode

		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			//answ.Err = err
			mu.Lock()
			//answs[idx] = &answ

			mu.Unlock()
			return
		}

		answ.Data = string(buf)
		mu.Lock()
		//answs[idx] = &answ

		mu.Unlock()
	}(i, url)
	//}

	wg.Wait()

	return buf, nil
	/*defer wg.Done()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// anonymous struct to pack and unpack data in the channel
	c := make(chan struct {
		r   *http.Response
		err error
	}, 1)

	req, _ := http.NewRequest("GET", url, nil)
	go func() {
		mu.Lock()
		resp, err := client.Do(req)
		fmt.Println("Doing http request is a hard job")
		pack := struct {
			r   *http.Response
			err error
		}{resp, err}
		c <- pack
		mu.Unlock()
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for client.Do
		fmt.Println("Cancel the context")
		return nil, ctx.Err()
	case ok := <-c:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error ", err)
			return nil, err
		}

		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Server Response: %s\n", out)
	}
	return nil, nil*/
}

func main() {
	ctx := context.Background()
	go func() {
		http.HandleFunc("/hello", helloHandler)

		//http.HandleFunc("/long", longHanlder)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("error when starting a server\n")
		}
	}()

	//name := "ok"
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
	client := &http.Client{}
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
