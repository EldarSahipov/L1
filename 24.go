package main

import (
	"fmt"
	"math"
)

// Point Определение структуры Point с полями x и y для представления точек на плоскости.
type Point struct {
	x, y float64
}

// NewPoint Конструктор NewPoint создает и возвращает новую точку с заданными координатами x и y.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод distance вычисляет расстояние между текущей точкой (p) и переданной точкой (a) на плоскости.
// Он использует формулу расстояния между двуми точками: sqrt((x2-x1)^2 + (y2-y1)^2).
func (p *Point) distance(a Point) float64 {
	ab := math.Sqrt(math.Pow(p.x-a.x, 2) + math.Pow(p.y-a.y, 2))
	return ab
}

func main() {
	// Создаем две точки A и B с заданными координатами.
	A, B := NewPoint(5, 5), NewPoint(0, 0)

	// Вызываем метод distance для точки A, передавая точку B в качестве аргумента,
	// чтобы вычислить расстояние между ними.
	distance := A.distance(B)

	// Выводим расстояние между точками A и B на экран.
	fmt.Println("Расстояние между точкой A и точкой B:", distance)
}
