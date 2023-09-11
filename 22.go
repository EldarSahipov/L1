package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем переменные a и b с помощью типа big.Int и устанавливаем им значения > 2^20.
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значения переменных a и b.
	a.SetString("1044563866546446546325464568576", 10)
	b.SetString("2097486486464654654646443431152", 10)

	// Умножение: a * b
	mulResult := new(big.Int)
	mulResult.Mul(a, b)

	// Деление: a / b
	divResult := new(big.Int)
	divResult.Div(a, b)

	// Сложение: a + b
	addResult := new(big.Int)
	addResult.Add(a, b)

	// Вычитание: a - b
	subResult := new(big.Int)
	subResult.Sub(a, b)

	// Вывод результатов операций.
	fmt.Printf("Умножение: %s\n", mulResult.String())
	fmt.Printf("Деление: %s\n", divResult.String())
	fmt.Printf("Сложение: %s\n", addResult.String())
	fmt.Printf("Вычитание: %s\n", subResult.String())
}
