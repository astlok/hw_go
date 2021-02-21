package uniq

import (
	"strings"
	"unicode/utf8"
)

type F struct {
	Exists bool
	NumFields int
}

type S struct {
	Exists bool
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

func Uniq(lines []string, options Options) []string {
	// перевод всех строк в нижний регистр
	if options.I {
		for _, line := range lines {
			line = strings.ToLower(line)
		}
	}
	// вырезаем numFields полей из массива строк
	if options.F.Exists {
		for _, line := range lines {
			fieldsInLine := len(strings.Split(line, " "))
			if  fieldsInLine < options.F.NumFields {
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
			if  charsInLine < options.S.NumChars {
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
	if options.C {
		//TODO: подсчитать количество встречаний строки во входных данных.
		// Вывести это число перед строкой отделив пробелом.
	} else if options.D {
		//TODO: вывести только те строки, которые повторились во входных данных.
	} else if options.U {
		//TODO: вывести только те строки, которые не повторились во входных данных.
	} else {
		//TODO: вывод уникальных строк из входных данных.
	}
	return lines //[]string{"TODO return"}
}