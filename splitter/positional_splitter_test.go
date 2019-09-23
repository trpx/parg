package splitter

import (
	"reflect"
	"testing"
)

type positionalSplitterTestCase struct {
	opt               Options
	args              []string
	expectedResult    Parts
	expectedRemainder []string
}

var positionalSplitterTestSuit = []positionalSplitterTestCase{
	{
		opt:  Options{prefix: "-"},
		args: []string{"a", "b"},
		expectedResult: Parts{
			Positional: []string{"a", "b"},
		},
		expectedRemainder: []string{},
	},
}

func TestPositionalSplitter(t *testing.T) {

	for idx, testCase := range positionalSplitterTestSuit {

		var testCaseNum = idx + 1

		args := testCase.args
		opt := testCase.opt
		expectedResult := testCase.expectedResult
		expectedRemainder := testCase.expectedRemainder

		splitter := PositionalSplitter{
			opt: opt,
		}

		result, remainder := splitter.Split(args)

		if !reflect.DeepEqual(expectedResult, result) {
			t.Errorf(
				"Test case %#v failed:"+
					"\nexpected result:\n%#v"+
					"\ngot:\n%#v"+
					"\nwhen splitting args: %#v"+
					"\nwith options: %#v",
				testCaseNum, expectedResult, result, args, opt,
			)
		}
		if !reflect.DeepEqual(expectedRemainder, remainder) {
			t.Errorf(
				"Test case %#v failed:"+
					"\nexpected remainder:\n%#v"+
					"\ngot:\n%#v"+
					"\nwhen splitting args: %#v"+
					"\nwith options: %#v",
				testCaseNum, expectedRemainder, remainder, args, opt,
			)
		}
	}
}
