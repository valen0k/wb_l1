package main

import (
	"fmt"
	"strings"
)

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

1. глобальная переменная (доступ и возможность изменить)
2. someFunc функция ради функции
3. локация функции createHugeString, если она реализована
4. возможно память, тк срез ссылается на ту же память, что и строка

*/

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

func someFunc() {
	var sb strings.Builder
	v := createHugeString(1 << 10)
	sb.WriteString(v[:100])
	str := sb.String()
	fmt.Printf("New string: %s\n", str)
}

func main() {
	someFunc()
}
