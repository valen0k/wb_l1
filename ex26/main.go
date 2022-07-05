package main

import (
	"fmt"
	"log"
	"strings"
)

func UniquenessCheck(str string) bool {
	newStr := []rune(strings.ToLower(str))
	buf := make(map[rune]struct{})

	for _, s := range newStr {
		if _, ok := buf[s]; ok {
			return false
		}
		buf[s] = struct{}{}
	}

	return true
}

func main() {
	var str string

	fmt.Print("➜ ")
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(UniquenessCheck(str))
}
