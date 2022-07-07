package ex23

import "errors"

//если порядок неважен
//реализация с заменой последнего и i-ого элемента
//с последующим возвращением среза без последнего элемента
func RemoveItem1[T any](s []T, i int) ([]T, error) {
	if i < 0 || i >= len(s) {
		return s, errors.New("index is out of range")
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1], nil
}

//если порядок важен
//вазращается два среза без i-ого элемента
func RemoveItem2[T any](s []T, i int) ([]T, error) {
	if i < 0 || i >= len(s) {
		return s, errors.New("index is out of range")
	}
	return append(s[:i], s[i+1:]...), nil
}
