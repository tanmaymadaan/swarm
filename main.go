package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func hitLb() {

}

func main() {
	var wg sync.WaitGroup

	lbUrl := "http://localhost:8080/"
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		time.Sleep(2 * time.Millisecond)

		go func(i int) {
			defer wg.Done()
			if mod := i % 1000; mod == 0 {
				fmt.Printf("Goroutine %d started\n", i)
			}

			_, err := http.Get(lbUrl)
			if err != nil {
				fmt.Printf("Error while calling LB: %d %v\n", i, err)
				return
			}

			if mod := i % 100; mod == 0 {
				fmt.Printf("Successful call: %d\n", i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished")
}
