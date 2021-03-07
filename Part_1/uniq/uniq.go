package uniq

import (
	"strconv"
	"strings"
)

type F struct {
	Exists    bool
	NumFields int
}

type S struct {
	Exists   bool
	NumChars int
}

type Options struct {
	C bool
	D bool
	U bool
	I bool
	F F
	S S
}

type LineOccursCount struct {
	CompLine  string //для сравнения
	ActualLine string //для вывода результата
	Count int
}

//проверяет входит ли элемент в массив и возвращает его последнее вхождение или -1
func ContainsLastNum(slice []LineOccursCount, line string, sensCase bool) int {
	lastNum := -1
	for i, elem := range slice {
		if line == elem.CompLine || (sensCase && strings.ToLower(line) == strings.ToLower(elem.CompLine)) {
			lastNum = i
		}
	}
	return lastNum
}

func Uniq(lines []string, options Options) []string {
	//переписываем все уникальные строки в массив, подсчитывая сколько раз они встречались
	lineOccursCount := []LineOccursCount{}
	for i, line := range lines {
		if options.F.Exists {
			lineArray := strings.Split(line, " ")
			lineArray = lineArray[options.F.NumFields:]
			line = strings.Join(lineArray, " ")
		}
		if num := ContainsLastNum(lineOccursCount, line[options.S.NumChars:], options.I); num != -1 && num == len(lineOccursCount)-1 {
			lineOccursCount[num].Count++
		} else {
			lineOccursCount = append(lineOccursCount, LineOccursCount{
				CompLine: line[options.S.NumChars:],
				ActualLine:  lines[i],
				Count: 1,
			})
		}
	}

	//формируем рузультирующий массив
	result := make([]string, len(lineOccursCount))
	for _, lineOccurs := range lineOccursCount {
		switch {
		case options.C:
			//подсчитывает количество встречаний строки во входных данных.
			result = append(result, strconv.Itoa(lineOccurs.Count)+" "+lineOccurs.ActualLine)
		case options.D && lineOccurs.Count > 1:
			//строки, которые повторились во входных данных
			result = append(result, lineOccurs.ActualLine)
		case options.U && lineOccurs.Count == 1:
			//строки, которые не повторились во входных данных.
			result = append(result, lineOccurs.ActualLine)
		case !options.C && !options.D && !options.U:
			result = append(result, lineOccurs.ActualLine)
		}
	}
	return result
}
