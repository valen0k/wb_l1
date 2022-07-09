package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var timeToDie uint

	fmt.Print("➜ ")
	_, err := fmt.Scan(&timeToDie)
	if err != nil {
		log.Fatalln(err)
	}
	//создаем context.WithTimeout (аналогично можно с context.WithDeadline)
	//для завершения программы через timeToDie секунд
	timeout, cancel := context.WithTimeout(context.Background(), time.Duration(timeToDie)*time.Second)
	defer cancel()

	ch := make(chan uint)
	wg := sync.WaitGroup{}

	//горутина для последовательно отправлять значения в канал
	go func() {
		wg.Add(1)
		defer wg.Done()
		n := uint(0)
		for {
			select {
			case <-timeout.Done():
				fmt.Println("goroutine finish")
				close(ch)
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()

	//основная горутина, в которой происходит чтение из канала
	for {
		select {
		case <-timeout.Done():
			wg.Wait()
			fmt.Println("main finish")
			return
		case tmp := <-ch:
			fmt.Println(tmp)
		}
	}
}
