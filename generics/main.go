package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T Number] (number1 T, number2 T) T {
	return number1 + number2
}

func Comparable[T comparable] (number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func Max[T constraints.Ordered] (number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer] (vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2 
	z = 3

	fmt.Println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1.1, "b": 2.5, "c": 3.5}))

	fmt.Println(Soma(3, 5))

	fmt.Println(Comparable(3, 5))

	fmt.Println(Max(7.9, 92.1))

	fmt.Println(concat([]MyString{"N", "i", "k", "o", "l", "l", "a", "s"}))
}