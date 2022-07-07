package ex02

import (
	"fmt"
	"sync"
)

//последовательность вывода не гарантируется
func SquareNumber(arr []int) {
	//Использование sync.WaitGroup для ожидания завершения нескольких goroutine
	wg := sync.WaitGroup{}

	//установить количество goroutine, которых необходимо ожидать
	wg.Add(len(arr))
	for _, value := range arr {
		go func(n int) {
			//goroutine запускается и вызывает Done, когда завершается
			defer wg.Done()
			fmt.Println(n * n)
		}(value)
	}

	//Wait может быть использован, чтобы блокировать, пока все goroutine не завершились
	wg.Wait()
}
