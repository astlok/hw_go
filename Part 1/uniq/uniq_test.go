package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var UniqTests = []struct {
	name string
	options Options
	in      []string
	out     []string
}{
	{ 	"Без опций",
		Options{
			C: false,
			D: false,
			U: false,
			I: false,
			F: F{
				Exists:    false,
				NumFields: 0,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		[]string{
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{ 	"Флаг -с",
		Options{
			C: true,
			D: false,
			U: false,
			I: false,
			F: F{
				Exists:    false,
				NumFields: 0,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		[]string{
			"3 I love music.",
			"1 \n",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{ 	"Флаг -d",
		Options{
			C: false,
			D: true,
			U: false,
			I: false,
			F: F{
				Exists:    false,
				NumFields: 0,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		[]string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	{ 	"Флаг -u",
		Options{
			C: false,
			D: false,
			U: true,
			I: false,
			F: F{
				Exists:    false,
				NumFields: 0,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		[]string{
			"\n",
			"Thanks.",
		},
	},
	{ 	"Флаг -i",
		Options{
			C: false,
			D: false,
			U: false,
			I: true,
			F: F{
				Exists:    false,
				NumFields: 0,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"\n",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
			"I love music of kartik.",
			"I love MuSIC of Kartik.",
		},
		[]string{
			"I LOVE MUSIC.",
			"\n",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of kartik.",
		},
	},
	{ 	"Флаг -f num",
		Options{
			C: false,
			D: false,
			U: false,
			I: false,
			F: F{
				Exists:    true,
				NumFields: 1,
			},
			S: S{
				Exists:   false,
				NumChars: 0,
			},
		},
		[]string{
			"We love music.",
			"I love music.",
			"They love music.",
			"\n",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		[]string{
			"We love music.",
			"\n",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{ 	"Флаг -s num",
		Options{
			C: false,
			D: false,
			U: false,
			I: false,
			F: F{
				Exists: false,
				NumFields: 0,
			},
			S: S{
				Exists: true,
				NumChars: 1,
			},
		},
		[]string{
			"I love music.",
			"A love music.",
			"C love music.",
			"\n",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		[]string{
			"I love music.",
			"\n",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
}

func TestUniq(t *testing.T) {
	for _, tt := range UniqTests {
		t.Run(tt.name, func(t *testing.T) {
			lines := Uniq(tt.in, tt.options)
			require.Equal(t, tt.out, lines)
		})
	}
}
