package main

import (
	"fmt"
	"reflect"
)

//определить тип можно вручную
func CheckType1(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("type: int")
	case string:
		fmt.Println("type: string")
	case bool:
		fmt.Println("type: bool")
	case chan bool:
		fmt.Println("type: chan bool")
	case chan string:
		fmt.Println("type: chan string")
	case chan int:
		fmt.Println("type: chan int")
	default:
		fmt.Println("type: something")
	}
}

//определение типа с помощью reflect
func CheckType2(data interface{}) {
	fmt.Println("Type:", reflect.TypeOf(data))
}

//определение типа с помощью fmt
func CheckType3(data interface{}) {
	fmt.Printf("Type: %T\n", data)
}

func main() {
	i := 123
	s := "123"
	b := true
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan bool)
	CheckType1(i)
	CheckType2(i)
	CheckType3(i)
	CheckType1(s)
	CheckType2(s)
	CheckType3(s)
	CheckType1(b)
	CheckType2(b)
	CheckType3(b)
	CheckType1(c1)
	CheckType2(c1)
	CheckType3(c1)
	CheckType1(c2)
	CheckType2(c2)
	CheckType3(c2)
	CheckType1(c3)
	CheckType2(c3)
	CheckType3(c3)
}
