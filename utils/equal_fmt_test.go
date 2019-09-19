package utils

import (
	"testing"
)

type equalFmtTestCase struct {
	a              interface{}
	b              interface{}
	expectedResult bool
}

var equalFmtTestSuit = []equalFmtTestCase{
	{
		[]string{"a", "b"},
		[]string{"a", "b"},
		true,
	},
	{
		[]string{"a", "b"},
		[]string{"a", "c"},
		false,
	},
	{
		[]string{"a", "b"},
		[]string{"a"},
		false,
	},
	{
		[]string{},
		[]string{},
		true,
	},
	{
		[]string{"a", "b"},
		[]string{},
		false,
	},
	{
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 2, 3},
		},
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 2, 3},
		},
		true,
	},
	{
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 2, 3},
		},
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 4, 3},
		},
		false,
	},
	{
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 2, 3},
		},
		[]string{},
		false,
	},
	{
		map[string][]int{
			"a": {1, 2, 3},
			"b": {1, 2, 3},
		},
		false,
		false,
	},
	{
		1,
		1,
		true,
	},
	{
		1,
		2,
		false,
	},
	{
		1.0,
		1,
		false,
	},
	{
		1.001,
		1.001,
		true,
	},
	{
		1.002,
		1.001,
		false,
	},
}

func TestEqualFmt(t *testing.T) {
	for _, testCase := range equalFmtTestSuit {
		a := testCase.a
		b := testCase.b
		expectedResult := testCase.expectedResult

		result := EqualFmt(a, b)

		if result != expectedResult {
			t.Errorf(
				"Expected %#v, got %#v when comparing %#v for equality with %#v",
				expectedResult, result, a, b,
			)
		}
	}
}
