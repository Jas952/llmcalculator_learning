package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(x int, y int) int { return x + y }
func sub(x int, y int) int { return x - y }
func mul(x int, y int) int { return x * y }
func div(x int, y int) int { return x / y }

type opMapFunc func(int, int) int // объявление фнукциональных типов. Проще присвоить, если дальше func(int, int)int будет использоваться неоднократно

var opMap = map[string]opMapFunc{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func condition(expression []string) (result int, err error) {

	if len(expression) < 3 {
		fmt.Println("Invalid expression")
		return 0, err
	}

	p1, err := strconv.Atoi(expression[0])
	if err != nil {
		fmt.Println("Первая строка не является числом")
		return 0, err
	}

	op := expression[1]
	check, ok := opMap[op]
	if !ok {
		fmt.Println("Невалидный знак")
		return 0, err
	}

	p2, err := strconv.Atoi(expression[2])
	if err != nil {
		fmt.Println("Второе значение не является числом")
		return 0, err
	}

	result = check(p1, p2)
	return result, err
}

func main() {
	reader := bufio.NewReader(os.Stdin) // даем команду на сбор инфы

	fmt.Print("Введите выражение, которое нужно посчитать: ")
	calc, _ := reader.ReadString('\n')
	calc = strings.TrimSpace(calc) // очищаем выражение

	for _, v := range calc {
		symbol := string(v)
		if opMap[symbol] != nil {
			calc = strings.ReplaceAll(calc, symbol, " "+symbol+" ") // идет перезапись содержимого сalc = ... и создаем пробелы между знаком
		}
	}
	sliceData := strings.Fields(calc) // создаем срез из элементов
	x, err := condition(sliceData)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(calc, "=", x)
	}

}
