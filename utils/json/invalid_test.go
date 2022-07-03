package json_test

import "testing"

func TestInvalid(t *testing.T) {
	tests := []testCase{
		{
			Json:  `'`,
			Error: true,
		},
		{
			Json:  `"`,
			Error: true,
		},
		{
			Json:  `(`,
			Error: true,
		},
		{
			Json:  `[`,
			Error: true,
		},
		{
			Json:  `]`,
			Error: true,
		},
		{
			Json:  `{`,
			Error: true,
		},
		{
			Json:  `}`,
			Error: true,
		},

		{
			Json:  `{'}`,
			Error: true,
		},
		{
			Json:  `{"}`,
			Error: true,
		},
		{
			Json:  `{(}`,
			Error: true,
		},

		{
			Json:  `[']`,
			Error: true,
		},
		{
			Json:  `["]`,
			Error: true,
		},
		{
			Json:  `[(]`,
			Error: true,
		},

		{
			Json:  `{[1],'}`,
			Error: true,
		},
		{
			Json:  `{[1],"}`,
			Error: true,
		},
		{
			Json:  `{[1],(}`,
			Error: true,
		},

		{
			Json:  `{"foo": '}`,
			Error: true,
		},
		{
			Json:  `{"foo": "}`,
			Error: true,
		},
		{
			Json:  `{"foo": (}`,
			Error: true,
		},

		{
			Json:  `{"1": 1 "2": 2}`,
			Error: true,
		},
		{
			Json:  `{"1": 1, "2", 2}`,
			Error: true,
		},
		{
			Json:  `{"1", 1, "2", 2}`,
			Error: true,
		},
		{
			Json:  `{"1": 1: "2": 2}`,
			Error: true,
		},
		{
			Json:  `{"1": 1. "2": 2}`,
			Error: true,
		},

		{
			Json:  `{"1": true. }`,
			Error: true,
		},
		{
			Json:  `{"1": tru }`,
			Error: true,
		},
		{
			Json:  `{"1": True }`,
			Error: true,
		},
		{
			Json:  `{"1": TRUE }`,
			Error: true,
		},
		{
			Json:  `{"1": false. }`,
			Error: true,
		},
		{
			Json:  `{"1": fals }`,
			Error: true,
		},
		{
			Json:  `{"1": False }`,
			Error: true,
		},
		{
			Json:  `{"1": FALSE }`,
			Error: true,
		},

		{
			Json:  `[1 1]`,
			Error: true,
		},
		{
			Json:  `[1.0 1.0]`,
			Error: true,
		},
		{
			Json:  `["1" "1"]`,
			Error: true,
		},
		{
			Json:  `[true true]`,
			Error: true,
		},
		{
			Json:  `[false false]`,
			Error: true,
		},
		{
			Json:  `[eurt eurt]`,
			Error: true,
		},

		{
			Json:  `{1 1}`,
			Error: true,
		},
		{
			Json:  `{1.0 1.0}`,
			Error: true,
		},
		{
			Json:  `{"1" "1"}`,
			Error: true,
		},
		{
			Json:  `{true true}`,
			Error: true,
		},
		{
			Json:  `{false false}`,
			Error: true,
		},
		{
			Json:  `{eurt eurt}`,
			Error: true,
		},

		{
			Json:  `{"foo": "bar"}}`,
			Error: true,
		},
		{
			Json:  `# foo bar`,
			Error: true,
		},
	}

	runTestCases(t, tests)
}
