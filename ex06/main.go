package ex06

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ContextFinish() {
	//аналогично с WithDeadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func(number int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("finished", number)
					return
				default:
					fmt.Println("here", number)
					time.Sleep(time.Second)
				}
			}
		}(i)
	}

	wg.Wait()
}

func SignalFinish() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for i := 1; i <= 5; i++ {
		go func(number int) {
			for {
				fmt.Println("here", number)
				time.Sleep(time.Second)
			}
		}(i)
	}

	<-sigs
}

func AfterFinish() {
	for i := 1; i <= 5; i++ {
		go func(number int) {
			for {
				fmt.Println("here", number)
				time.Sleep(time.Second)
			}
		}(i)
	}

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("finished")
	}
}
