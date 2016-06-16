package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	t0 := time.Now()
	url := "http://localhost:10888"
	count := 1000
	goRoutineCount := 100

	for i := 0; i < count; i++ {
		var wg sync.WaitGroup
		for j := 1; j <= +goRoutineCount; j++ {
			fmt.Printf("\r%d", i*goRoutineCount+j)
			wg.Add(1)
			go run(&wg, "GET", url)
		}
		wg.Wait()
	}
	t1 := time.Now()
	fmt.Printf("\nВремя выполнения %v .\n", t1.Sub(t0))
	return
}

func run(wg *sync.WaitGroup, method, url string) {
	client := &http.Client{}
	defer wg.Done()
	req, err := http.NewRequest(method, url, nil)
	req.SetBasicAuth("admin", "123321")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
	}
	return
}
