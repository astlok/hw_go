package main

import (
	"fmt"
	"main/uniq"
)

func main() {
	options := uniq.Options{
		C: false,
		D: false,
		U: false,
		I: false,
		F: uniq.F {
			Exists: true,
			NumFields: 10,
		},
		S: uniq.S {
			Exists: true,
			NumChars: 5,
		},
	}

	lines := []string{
		"I want to tell u about yourself",
		"im a student from Moscow",
		"My name is Oleg",
		"I love programming and snowboarding",
	}

	lines = uniq.Uniq(lines, options)

	fmt.Println(lines)
}
