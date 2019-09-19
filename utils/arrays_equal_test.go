package utils

import (
	"testing"
)

type stringArraysEqualTestCase struct {
	arrayA         []string
	arrayB         []string
	expectedResult bool
}

var stringArraysEqualTestCaseSuit = []stringArraysEqualTestCase{
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
}

func TestStringArraysEqual(t *testing.T) {
	for _, testCase := range stringArraysEqualTestCaseSuit {
		arrayA := testCase.arrayA
		arrayB := testCase.arrayB
		expectedResult := testCase.expectedResult

		result := StringArraysEqual(arrayA, arrayB)

		if result != expectedResult {
			t.Errorf(
				"Expected %#v, got %#v when comparing %#v for equality with %#v",
				expectedResult, result, arrayA, arrayB,
			)
		}
	}
}
