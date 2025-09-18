package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Introduce la operación a solucionar (presiona Enter para finalizar):")
	fmt.Println("Separado con un espacio ' ' Ej: 10 + 8:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSpace(input)

	calculatedValue := calculate(input)
	fmt.Printf("Result: %.2f\n", calculatedValue)

}

func separates(expression string) []string {
	return strings.Split(expression, " ")
}
func calculate(expression string) float64 {
	parts := separates(expression)
	if len(parts) != 3 {
		return 0
	}

	a, errA := strconv.ParseFloat(parts[0], 64)
	b, errB := strconv.ParseFloat(parts[2], 64)
	op := parts[1]
	if errA != nil || errB != nil {
		return 0
	}
	return operation(a, b, op)
}

func operation(a float64, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	case "%":
		return float64(int(a) % int(b))
	default:
		return 0
	}
}
