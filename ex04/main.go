package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var works int

	fmt.Print("➜ ")
	_, err := fmt.Scan(&works)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	ch := make(chan interface{})

	wg := sync.WaitGroup{}
	wg.Add(works)

	for i := 0; i < works; i++ {
		go func(number int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("finished", number)
					return
				case value := <-ch:
					fmt.Printf("%d) info: %d\n", number, value)
				}
			}
		}(i)
	}

	n := uint(0)
	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			fmt.Println("main finished")
			return
		default:
			ch <- n
			n++
			time.Sleep(time.Second)
		}
	}
}
