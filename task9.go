package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
	Task 9. Реализовать калькулятор на горутинах (без пакета sync)
	На входе: строка вида "1+1,2 -1, 6 / 2, 3*2"
	На выходе: результат арифметических операций
*/

type Expression struct {
	operand1, operand2, result float64
	operator                   byte
	haveError                  bool
}

func sum(e *Expression, ch chan Expression) {
	e.result = e.operand1 + e.operand2
	ch <- *e
}

func multiplication(e *Expression, ch chan Expression) {
	e.result = e.operand1 * e.operand2
	ch <- *e
}

func difference(e *Expression, ch chan Expression) {
	e.result = e.operand1 - e.operand2
	ch <- *e
}

func division(e *Expression, ch chan Expression) {
	if e.operand2 == 0 {
		fmt.Printf("%f / %f - division by zero\n", e.operand1, e.operand2)
		e.haveError = true
	} else {
		e.result = e.operand1 / e.operand2
	}

	ch <- *e
}

func main() {
	input := `6*6,7-8,2 /  0,1+ 1,2 +4, 2* 2,9/  3,    2  *  8   , 2*f, a8, *, 5f*2,2//2,2++2,6-2`

	// regexp1 - для удаления всех пробелов из строки
	regexForSpaces, err := regexp.Compile("\\s+")
	if err != nil {
		fmt.Println("Regular expression regexForSpaces is not parsed!")
		return
	}
	// очистка пробелов
	input = regexForSpaces.ReplaceAllString(input, "")
	// разделение по запятым на слайс строк
	expressions := strings.Split(input, ",")

	// regexp 2 - проверка на арифметическое выражение
	regexForOperation, err := regexp.Compile(`^[0-9]*[.,]?[0-9]+(\s)*[+-/*](\s)*[0-9]*[.,]?[0-9]+$`)
	if err != nil {
		fmt.Println("Regular expression regexForOperation is not parsed!")
		return
	}

	// количество корректных выражений
	countOfExprs := 0
	ch := make(chan Expression)
	for _, expressionString := range expressions {
		// проверка регуляркой на соответствие арифметическому выражению
		if !regexForOperation.MatchString(expressionString) {
			fmt.Printf("String '%s' is not expression\n", expressionString)
			continue
		}

		operatorIndex := strings.IndexAny(expressionString, "+-/*")

		// Парсинг вещественных значений
		op1, err := strconv.ParseFloat(expressionString[:operatorIndex], 64)
		if err != nil {
			panic(err)
		}

		op2, err := strconv.ParseFloat(expressionString[operatorIndex+1:], 64)
		if err != nil {
			panic(err)
		}

		expression := Expression{operand1: op1, operator: expressionString[operatorIndex], operand2: op2}
		switch expression.operator {
		case 43:
			go sum(&expression, ch)
		case 42:
			go multiplication(&expression, ch)
		case 47:
			go division(&expression, ch)
		case 45:
			go difference(&expression, ch)
		default:
			fmt.Println("Operation not found")
		}
		countOfExprs++
	}

	// вывод
	for i := 0; i < countOfExprs; i++ {
		select {
		case expr := <-ch:
			if !expr.haveError {
				fmt.Printf("%f %c %f = %f\n", expr.operand1, expr.operator, expr.operand2, expr.result)
			}
		}
	}

}
