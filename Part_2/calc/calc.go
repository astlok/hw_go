package calc

import (
	"fmt"
	"strconv"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	if len(s) == 0 {
		return s, ""
	}

	l := len(s)
	return  s[:l-1], s[l-1]
}

func (s stack) Top() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s) - 1]
}

var OperatorPriority = map[string]int{
	"(": 1,
	")": 1,
	"+": 2,
	"-": 2,
	"*": 3,
	"/": 3,
}

func getRes(val1 float64, val2 float64, sign string) float64 {
	switch sign {
	case "+":
		return val1 + val2
	case "-":
		return val2 - val1
	case "*":
		return val1 * val2
	case "/":
		return val2 / val1
	}
	return 0
}

//возвращает след сущность (число или знак) и какое количество знаков она занимает
func getNextEssence(expr string) (string, int) {
	var result string
	for i, ch := range expr {
		if string(ch) == " " {
			continue
		}
		if OperatorPriority[string(ch)] > 0 && len(result) == 0 {
			return string(ch), i + 1
		} else if OperatorPriority[string(ch)] > 0 {
			return result, i
		} else if _, err := strconv.Atoi(string(ch)); err == nil || ch == '.'{
			result += string(ch)
		}
	}
	return result, len(expr)
}

func signHandler(stackSigns *stack, stackDigits *stack, next string, result *float64) error {
	calculate := func(next string) error {
		var val1String string
		var val2String string

		*stackDigits, val1String = stackDigits.Pop()
		*stackDigits, val2String = stackDigits.Pop()

		val1Int, err1 := strconv.ParseFloat(val1String, 64)
		val2Int, err2 := strconv.ParseFloat(val2String, 64)

		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}

		var sign string
		*stackSigns, sign = stackSigns.Pop()

		*result = getRes(val1Int, val2Int, sign)
		*stackDigits = stackDigits.Push(fmt.Sprintf("%f", *result))

		if next != "" && next != ")" {
			*stackSigns = stackSigns.Push(next)
		}
		return nil
	}
	if OperatorPriority[stackSigns.Top()] < OperatorPriority[next] || (stackSigns.Top() == "(" && next != ")") {
		*stackSigns = stackSigns.Push(next)
	} else if next == "(" {
		*stackSigns = stackSigns.Push(next)
	} else  if next != ")" && next != "(" {
		if err := calculate(next); err != nil {
			return err
		}
	} else {
		for stackSigns.Top() != "(" {
			if err := calculate(next); err != nil {
				return err
			}
		}
		*stackSigns, _ = stackSigns.Pop()
	}
	return nil
}

func Calc(expr string) (float64, error) {
	var result float64

	var stackDigits stack
	var stackSigns stack


	for len(expr) != 0 || len(stackSigns) > 0 {
		next, length := getNextEssence(expr)
		if _, err := strconv.ParseFloat(next, 64); err == nil {
			stackDigits = stackDigits.Push(next)
		} else {
			if err := signHandler(&stackSigns, &stackDigits, next, &result); err != nil {
				return 0, err
			}
		}
		expr = expr[length:]
	}
	return result, nil
}
