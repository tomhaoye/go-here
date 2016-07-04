package main

import "time"
import "fmt"

func worker(start chan bool) {
    heartbeat := time.Tick(1 * time.Second)
    for {
	    select {
            case <- heartbeat:
            fmt.Printf("A")
        }	
    }
}

func main() {
	start := make(chan bool)
	worker(start)
}