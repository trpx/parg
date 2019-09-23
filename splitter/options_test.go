package splitter

import (
	"testing"
)

type isPositionalArgTestCase struct {
	opt            Options
	arg            string
	expectedResult bool
}

var isPositionalArgTestSuit = []isPositionalArgTestCase{
	{
		opt: Options{
			prefix: "-",
		},
		arg:            "--arg",
		expectedResult: false,
	},
	{
		opt: Options{
			prefix:          "-",
			subCommandNames: []string{"a", "b"},
		},
		arg:            "b",
		expectedResult: false,
	},
	{
		opt: Options{
			prefix:          "-",
			subCommandNames: []string{"a", "b"},
		},
		arg:            "c",
		expectedResult: true,
	},
	{
		opt: Options{
			prefix: "-",
		},
		arg:            "a",
		expectedResult: true,
	},
}

func TestOptions_isPositionalArg(t *testing.T) {

	for idx, testCase := range isPositionalArgTestSuit {
		testCaseNum := idx + 1

		opt := testCase.opt
		arg := testCase.arg
		expectedResult := testCase.expectedResult

		result := opt.isPositionalArg(arg)

		if expectedResult != result {
			t.Errorf(
				"Test case %#v failed:\nexpected %#v, got %#v"+
					" when testing arg %#v on being positional with options:\n%#v.",
				testCaseNum, expectedResult, result, arg, opt,
			)
		}
	}
}

type isSubCommandArgTestCase struct {
	opt            Options
	arg            string
	expectedResult bool
}

var isSubCommandArgTestSuit = []isSubCommandArgTestCase{
	{
		opt: Options{
			prefix: "-",
		},
		arg:            "a",
		expectedResult: false,
	},
	{
		opt: Options{
			prefix:          "-",
			subCommandNames: []string{"a", "b"},
		},
		arg:            "b",
		expectedResult: true,
	},
	{
		opt: Options{
			prefix:          "-",
			subCommandNames: []string{"a", "b"},
		},
		arg:            "c",
		expectedResult: false,
	},
}

func TestOptions_isSubCommandArg(t *testing.T) {

	for idx, testCase := range isSubCommandArgTestSuit {
		testCaseNum := idx + 1

		opt := testCase.opt
		arg := testCase.arg
		expectedResult := testCase.expectedResult

		result := opt.isSubCommandArg(arg)

		if expectedResult != result {
			t.Errorf(
				"Test case %#v failed:\nexpected %#v, got %#v"+
					" when testing arg %#v on being positional with options:\n%#v.",
				testCaseNum, expectedResult, result, arg, opt,
			)
		}
	}
}
