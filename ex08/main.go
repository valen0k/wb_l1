package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var i int
	var num int64

	fmt.Print("➜ ")
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatalln(err)
	}

	num ^= 1 << i

	fmt.Println(strconv.FormatInt(num, 2))
}
