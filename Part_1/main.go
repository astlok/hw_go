package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"main/uniq"
	"os"
	"strings"
	"syscall"
)

func ReadLines(r io.Reader) []string {
	var lines string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}
	return strings.Split(lines, "\n")
}

func ReadFromStream(args []string) ([]string, error) {
	var r io.Reader
	switch len(args) {
	case 0:
		r = os.Stdin
	case 1, 2:
		readFile, err := os.Open(flag.Args()[0])
		if err != nil {
			return []string{}, err
		} else {
			defer readFile.Close()
		}
		r = readFile
	}
	return ReadLines(r), nil
}

func WriteLines(w io.Writer, lines []string) error {
	for i, line := range lines {
		if i != len(lines) - 1 {
			line += "\n"
		}

		buf := bytes.NewBufferString(line)

		if _, err := w.Write(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func WriteToStream(args []string, toWrite []string) error {
	var w io.Writer
	switch len(args) {
	case 0,1:
		w = os.Stdout
	case 2:
		writeFile, errWrite := os.OpenFile(args[1], os.O_WRONLY, 0666)

		if errWrite != nil {
			return errWrite
		} else {
			defer writeFile.Close()
		}
		w = writeFile
	}
	err := WriteLines(w, toWrite)
	return err
}

func ArgHandle() (uniq.Options, error) {
	options := uniq.Options{}
	options.C = *flag.Bool("c", false, "подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	options.D = *flag.Bool("d", false, "вывести только те строки, которые повторились во входных данных.")
	options.U = *flag.Bool("u", false, "вывести только те строки, которые не повторились во входных данных.")
	options.I = *flag.Bool("i", false, "вывести только те строки, которые не повторились во входных данных.")

	options.F.NumFields = *flag.Int("f", 0, "вывести только те строки, которые не повторились во входных данных.")
	options.F.Exists = options.F.NumFields > 0
	options.S.NumChars = *flag.Int("s", 0, "вывести только те строки, которые не повторились во входных данных.")
	options.S.Exists = options.S.NumChars  > 0
	flag.Parse()

	if (options.C && options.D) || (options.D && options.U) || (options.C && options.U) {
		return uniq.Options{}, errors.New("Incorrect arguments!\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}
	return options, nil
}

func main() {
	//чтение аргументов
	options, err := ArgHandle()
	if err != nil {
		fmt.Println(err.Error())
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}

	//чтение данных
	lines, err := ReadFromStream(flag.Args())
	if err != nil {
		fmt.Println(err.Error())
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}

	result := uniq.Uniq(lines, options)

	//запись результата
	if err = WriteToStream(flag.Args(), result); err != nil {
		fmt.Println(err.Error())
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}
}
