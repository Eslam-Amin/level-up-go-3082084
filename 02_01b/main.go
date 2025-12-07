package main

import (
	"flag"
	"fmt"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	ch := make(chan string)
	for i := 1; i < n+1; i++ {
		go func(i int) {
			ch <- fmt.Sprintf("[G%d]:%s\n", i, message)
		}(i)
	}

	for i := 1; i < n+1; i++ {
		log.Printf(<-ch)
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
