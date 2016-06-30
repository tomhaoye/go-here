package main

import "fmt"
import "time"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for {
			time.Sleep(1*time.Second)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
