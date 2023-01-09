package math

import "errors"

func Division(a, b float32)  (float32, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Operadores menor ou igual do que 0")
	}
	result := a / b
	return result, nil
}