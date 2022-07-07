package ex03

import (
	"sync"
)

func SumSquareNumber(arr []int) int {
	//Использование sync.WaitGroup для ожидания завершения нескольких goroutine
	wg := sync.WaitGroup{}
	//канал для записи сумм корней чисел со всех goroutine
	sum := make(chan int, 1)

	sum <- 0
	//установить количество goroutine, которых необходимо ожидать
	wg.Add(len(arr))
	for _, value := range arr {
		go func(n int) {
			//goroutine запускается и вызывает Done, когда завершается
			defer wg.Done()
			sum <- n*n + <-sum
		}(value)
	}

	//Wait может быть использован, чтобы блокировать, пока все goroutine не завершились
	wg.Wait()
	return <-sum
}
