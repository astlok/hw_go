package uniq

import (
	"strconv"
	"strings"
	"unicode/utf8"
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
	Line  string
	Count int
}

//проверяет входит ли элемент в массив и возвращает его последнее вхождение или -1
func ContainsLastNum(slice []LineOccursCount, line string, sensCase bool) int {
	lastNum := -1
	for i, elem := range slice {
		if line == elem.Line || (sensCase && strings.ToLower(line) == strings.ToLower(elem.Line)) {
			lastNum = i
		}
	}
	return lastNum
}

func Uniq(lines []string, options Options) []string {

	if options.F.Exists {
		for _, line := range lines {
			fieldsInLine := len(strings.Split(line, " "))
			if fieldsInLine < options.F.NumFields {
				lines = lines[1:]
				options.F.NumFields -= fieldsInLine
			} else if fieldsInLine == options.F.NumFields {
				lines = lines[1:]
				break
			} else if fieldsInLine > options.F.NumFields {
				lineArray := strings.Split(line, " ")
				lineArray = lineArray[options.F.NumFields:]
				lines[0] = strings.Join(lineArray, " ")
				break
			}
		}
	}
	//вырезаем numChars символов из массива строк
	if options.S.Exists {
		for _, line := range lines {
			charsInLine := utf8.RuneCountInString(line)
			if charsInLine < options.S.NumChars {
				lines = lines[1:]
				options.S.NumChars -= charsInLine
			} else if charsInLine == options.S.NumChars {
				lines = lines[1:]
				break
			} else if charsInLine > options.S.NumChars {
				lines[0] = line[options.S.NumChars:]
				break
			}
		}
	}

	//переписываем все уникальные строки в массив, подсчитывая сколько раз они встречались
	lineOccursCount := []LineOccursCount{}
	for i := 0; i < len(lines); i++ {
		if num := ContainsLastNum(lineOccursCount, lines[i], options.I); num != -1 && num == len(lineOccursCount)-1 {
			lineOccursCount[num].Count++
		} else {
			lineOccursCount = append(lineOccursCount, LineOccursCount{
				Line:  lines[i],
				Count: 1,
			})
		}
	}

	//формируем рузультирующий массив
	result := []string{}
	for _, lineOccurs := range lineOccursCount {
		if options.C {
			//подсчитывает количество встречаний строки во входных данных.
			result = append(result, strconv.Itoa(lineOccurs.Count)+" "+lineOccurs.Line)
		} else if options.D && lineOccurs.Count > 1 {
			//строки, которые повторились во входных данных
			result = append(result, lineOccurs.Line)
		} else if options.U && lineOccurs.Count == 1 {
			//строки, которые не повторились во входных данных.
			result = append(result, lineOccurs.Line)
		} else if !options.C && !options.D && !options.U {
			result = append(result, lineOccurs.Line)
		}
	}
	return result
}
