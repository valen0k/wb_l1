package main

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

func Swap[N Number](a, b *N) {
	*a, *b = *b, *a
}
