package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//заполняем массив рандомной величины рандомными значениями
	n := rand.Int() % 100
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int() % 10000
	}

	//создаем первый канал
	//и записываем в него значения из массива в отдельной горутине,
	//в конце записи закрываем канал
	chanOne := make(chan int)
	go func() {
		defer close(chanOne)
		for _, value := range arr {
			chanOne <- value
		}
	}()

	//создаем второй канал,
	//в который будут записываться числа из первого возведенные в квадрат,
	//в конце чтения канал закрывается
	chanTwo := make(chan int)
	go func() {
		defer close(chanTwo)
		for value := range chanOne {
			chanTwo <- value * value
		}
	}()

	//в основной горутине будет происходить чтение из второго канала и данные выводиться в stdout
	for value := range chanTwo {
		fmt.Println(value)
	}
}
