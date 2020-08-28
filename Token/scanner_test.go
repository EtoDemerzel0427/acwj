package token

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

type testScanner struct {
	input string
	result []*Token
}

var (
	testCases = []testScanner{
		{
			input: "2 + 3 * 5 - 8 / 3",
			result: []*Token{
				{token: Integer, intValue: 2},
				{token: Plus},
				{token: Integer, intValue: 3},
				{token: Aster},
				{token: Integer, intValue: 5},
				{token: Minus},
				{token: Integer, intValue: 8},
				{token: Slash},
				{token: Integer, intValue: 3},
			},
		},
		{
			input: "13 -6+  4*\n5\n       +\n08 / 3",
			result: []*Token{
				{token: Integer, intValue: 13},
				{token: Minus},
				{token: Integer, intValue: 6},
				{token: Plus},
				{token: Integer, intValue: 4},
				{token: Aster},
				{token: Integer, intValue: 5},
				{token: Plus},
				{token: Integer, intValue: 8},
				{token: Slash},
				{token: Integer, intValue: 3},
			},
		},
		{
			input: "12 34 + -56 * / - - 8 + * 2",
			result: []*Token{
				{token: Integer, intValue: 12},
				{token: Integer, intValue: 34},
				{token: Plus},
				{token: Minus},
				{token: Integer, intValue: 56},
				{token: Aster},
				{token: Slash},
				{token: Minus},
				{token: Minus},
				{token: Integer, intValue: 8},
				{token: Plus},
				{token: Aster},
				{token: Integer, intValue: 2},
			},
		},
		{
			input: "23 +\n18 -\n45.6 * 2\n/ 18",
			result: []*Token{
				{token: Integer, intValue: 23},
				{token: Plus},
				{token: Integer, intValue: 18},
				{token: Minus},
				{token: Integer, intValue: 45}, // after this token, an error occurs.
			},
		},
	}
)

func TestNewScanner(t *testing.T) {
	ts := NewScanner(strings.NewReader("123"))
	if ts.Tok.token != -2 {
		t.Error("Init fail.\n")
	}
}

func TestTokenScanner_Scan(t *testing.T) {
	for i, test := range testCases {
		result := testScan(test.input)

		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Case %d:\n expected: %+v\n result: %+v\n", i, test.result, result)
		}

	}
}

func testScan(input string) (result []*Token)  {
	ts := NewScanner(strings.NewReader(input))

	for {
		num, err := ts.Scan()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}
		if num != 1 || err != nil {
			break
		}
		fmt.Printf("Token %s", ts.Tok.String())

		fmt.Print("\n")
		result = append(result, ts.Tok)
	}

	return
}

