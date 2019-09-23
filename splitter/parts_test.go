package splitter

import (
	"reflect"
	"testing"
)

type partsTestCase struct {
	addBoolTrue    []string
	addBoolFalse   []string
	addString      []string
	expectedResult Parts
}

var partsTestSuit = []partsTestCase{
	{
		addBoolTrue:  []string{"a", "b", "c", "a", "a"},
		addBoolFalse: []string{"a", "b", "c"},
		addString:    []string{"a", "b", "c", "d", "a", "b", "a", "e"},
		expectedResult: Parts{
			NamedString: map[string][]string{
				"a": {"b", "b", "e"},
				"c": {"d"},
			},
			NamedBool: map[string][]bool{
				"a": {true, true, true, false},
				"b": {true, false},
				"c": {true, false},
			},
		},
	},
}

func TestParts(t *testing.T) {

	for idx, testCase := range partsTestSuit {

		var testCaseNum = idx + 1

		parts := NewParts()

		for _, s := range testCase.addBoolTrue {
			parts.addNamedBool(s, true)
		}
		for _, s := range testCase.addBoolFalse {
			parts.addNamedBool(s, false)
		}
		for idx, s := range testCase.addString {
			if idx%2 == 0 {
				parts.addNamedString(s, testCase.addString[idx+1])
			}
		}

		expected := testCase.expectedResult

		if !reflect.DeepEqual(expected, parts) {
			t.Errorf(
				"Test case %#v failed:"+
					"\nexpected:\n%#v"+
					"\ngot:\n%#v"+
					"\nwhen adding bool true: %#v"+
					"\nbool false: %#v"+
					"\nstrings ({arg, value, arg, value...}):\n%#v",
				testCaseNum, expected, parts, testCase.addBoolTrue, testCase.addBoolFalse, testCase.addString,
			)
		}
	}
}
