package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	sum := 0
	//delta := make(chan int, 1000)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//delta <- 1
			//fmt.Println("Goroutine start", i)
			sum++
			defer wg.Done()
			//fmt.Println("Goroutine ends:", i)
		}()
	}

	/*for i := 0; i < 1000; i++ {
		sum += <-delta
	}*/
	wg.Wait()
	fmt.Println(sum)
}
