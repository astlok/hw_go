package calc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var CalcTests = []struct {
	name string
	in      string
	out    	float64
}{
	{
		name: "test1",
		in: "1+2+3",
		out: 6,
	},
	{
		name: "test2",
		in: "1+ (2 *3)",
		out: 7,
	},
	{
		name: "test3",
		in: "2* 5 /10 + (3* (3 - 1) -1 ) + 6",
		out: 12,
	},
	{
		name: "test4",
		in: "22 + (30/10) - 15 + (2 + (3 * (4 - (4 /2) -1) + 4))",
		out: 19,
	},
	{
		name: "test5",
		in: "22 + 38 * ( (((2 + 3))) - (5 * ((4 -2) / 2) - 5) + (3*5))",
		out: 782,
	},
	{
		name: "test6",
		in: "2.2 +3.3",
		out: 5.5,
	},
	{
		name: "test7",
		in: "2.2 + 3.8 * ( (((2.15 + 37.25))) - (5.5 * ((4.1 -2) / 2) - 5.7) + (3*5))",
		out: 208.635,
	},

}

func TestCalc(t *testing.T) {
	for _, tt := range CalcTests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Calc(tt.in)
			require.Equal(t, tt.out, result)
		})
	}
}