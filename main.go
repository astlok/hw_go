package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"main/uniq"
	"os"
	"strings"
)

func ReadLines(r io.Reader) []string {
	var line string
	data := make([]byte, 64)
	for {
		n, err := r.Read(data)
		if err == io.EOF {
			break
		}
		line += string(data[:n])
	}
	return strings.Split(line, "\n")
}

func WriteLines(w io.Writer, lines []string) error {
	for i, line := range lines {
		if i != len(lines) - 1 {
			line += "\n"
		}

		buf := bytes.NewBufferString(line)
		_, err := w.Write(buf.Bytes())

		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	c := flag.Bool("c", false, "подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	d := flag.Bool("d", false, "вывести только те строки, которые повторились во входных данных.")
	u := flag.Bool("u", false, "вывести только те строки, которые не повторились во входных данных.")
	i := flag.Bool("i", false, "вывести только те строки, которые не повторились во входных данных.")

	f := flag.Int("f", 0, "вывести только те строки, которые не повторились во входных данных.")
	s := flag.Int("s", 0, "вывести только те строки, которые не повторились во входных данных.")
	flag.Parse()

	if (*c && *d) || (*d && *u) || (*c && *u) {
		fmt.Println("Incorrect arguments!\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}

	options := uniq.Options{
		C: *c,
		D: *d,
		U: *u,
		I: *i,
		F: uniq.F{
			Exists:    *f > 0,
			NumFields: *f,
		},
		S: uniq.S{
			Exists:   *s > 0,
			NumChars: *s,
		},
	}

	var r io.Reader
	var w io.Writer

	switch len(flag.Args()) {
	case 0:
		r = os.Stdin
		w = os.Stdout
	case 1:
		readFile, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Println("Incorrect arguments!\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		}
		r = readFile
		w = os.Stdout
	case 2:
		readFile, errRead := os.Open(flag.Args()[0])
		writeFile, errWrite := os.OpenFile(flag.Args()[1], os.O_WRONLY, 0666)
		if errRead != nil || errWrite != nil {
			fmt.Println("Incorrect arguments!\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		}
		r = readFile
		w = writeFile
	}

	lines := ReadLines(r)

	result := uniq.Uniq(lines, options)

	err := WriteLines(w, result)

	if err != nil {
		fmt.Println("Can't write to file")
		return
	}
}
