package main

import (
	"context"
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
				//defer mu.Unlock()
				APRS = append(APRS, &APIResponse{URL: u, Err: err})
				mu.Unlock()
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				mu.Lock()
				APRS = append(APRS, &APIResponse{URL: u, Data: "", StatusCode: resp.StatusCode, Err: err})
				mu.Unlock()
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
