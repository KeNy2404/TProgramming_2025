package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Герасимов Александр Евгеньевич")
	fmt.Println("Вариант №6")
	fmt.Println("x1 = ", calculate(0.1))
	fmt.Println("x2 = ", calculate(0.9))
	fmt.Println("x3 = ", calculate(1.2))
	fmt.Println("x4 = ", calculate(1.5))
	fmt.Println("x5 = ", calculate(2.3))

	fmt.Println("Решение задачи А")

	for i := 0.2; i <= 2.2; i += 0.1 {
		fmt.Println("x", i, " = ", calculate(i))
	}

	fmt.Println("Решение задачи Б")
	x := []float64{0.1, 0.9, 1.2, 1.5, 2.3}
	for i, value := range x {
		fmt.Println("x", i, " = ", calculate(value))
	}
}

func calculate(x float64) float64 {
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}
