package math

func Soma(a, b int) int {
	return a + b
}

func SomaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}