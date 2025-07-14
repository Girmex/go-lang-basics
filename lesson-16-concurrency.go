package main

import (
	"fmt"
	"time"
)
func task(id int) {
    fmt.Printf("Task %d starting\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Task %d done\n", id)
}

func Concurency() {
    for i := 1; i <= 3; i++ {
        go task(i) // concurrent tasks
    }
    time.Sleep(2 * time.Second) // wait for goroutines to finish
}
