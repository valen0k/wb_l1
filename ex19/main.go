package main

import (
	"fmt"
	"log"
)

func SwitchSymbol(str string) string {
	newRunes := []rune(str)

	for i, j := 0, len(newRunes)-1; i < j; i, j = i+1, j-1 {
		newRunes[i], newRunes[j] = newRunes[j], newRunes[i]
	}

	return string(newRunes)
}

func main() {
	var str string

	fmt.Print("➜ ")
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(SwitchSymbol(str))
}
