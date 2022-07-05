package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SwitchWords(str string) string {
	//разбиваем строку на массив слов
	split := strings.Split(str, " ")

	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	//собираем строку из массива слов обратно
	return strings.Join(split, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("➜ ")
	scanner.Scan()

	fmt.Println(SwitchWords(scanner.Text()))
}
