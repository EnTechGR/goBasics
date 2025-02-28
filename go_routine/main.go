package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// func say(message string) {
// 	for _, word := range strings.Fields(message) {
// 		fmt.Printf("Simon says: %s...\n", word)
// 		dur := time.Duration(rand.Intn(100)) * time.Millisecond
// 		time.Sleep(dur)
// 	}
// }

func sayConv(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur  := time.Duration(rand.Intn(100)) * time.Microsecond
		time.Sleep(dur)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		sayConv(1, "Hello, World!")
	}()


	go func() {
		defer wg.Done()
		sayConv(2, "Hello, Go!")
	}()

}