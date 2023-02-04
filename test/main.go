package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	tm := time.Now()
	wg := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			url := "http://localhost:9000/api/"
			method := "GET"

			client := &http.Client{}
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				fmt.Println(err)
				return
			}

			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(res.Body)

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(i, time.Since(tm), string(body)[:20])
		}(i)
	}
	wg.Wait()
}
