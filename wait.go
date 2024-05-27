package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup


func side() {
	time.Sleep(time.Second * 5)
	for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    wg.Done()
}
func main() {
	fmt.Println("Starting main")
	// Used to tell main to wait for the goroutines to finish.
	wg.Add(2)
	go side()
	go side()
	fmt.Println("Return to main")
	wg.Wait()
	fmt.Println("End of main")
}	
