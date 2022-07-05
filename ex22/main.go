package main

import (
	"fmt"
	"log"
	"math/big"
)

type BigInt struct {
	a, b *big.Int
}

func NewBigInt(a, b int64) *BigInt {
	return &BigInt{a: big.NewInt(a), b: big.NewInt(b)}
}

func (i *BigInt) Addition() *big.Int {
	return new(big.Int).Add(i.a, i.b)
}

func (i *BigInt) Subtraction() *big.Int {
	return new(big.Int).Sub(i.a, i.b)
}

func (i *BigInt) Multiplication() *big.Int {
	return new(big.Int).Mul(i.a, i.b)
}

func (i *BigInt) Division() *big.Int {
	return new(big.Int).Div(i.a, i.b)
}

func main() {
	var a, b int64
	var sign string

	fmt.Print("➜ ")
	_, err := fmt.Scan(&a, &sign, &b)
	if err != nil {
		log.Fatalln(err)
	}

	bigInt := NewBigInt(a, b)

	switch sign {
	case "+":
		fmt.Println(bigInt.Addition())
	case "-":
		fmt.Println(bigInt.Subtraction())
	case "*":
		fmt.Println(bigInt.Multiplication())
	case "/":
		fmt.Println(bigInt.Division())
	default:
		log.Fatalln("wrong sign")
	}
}
