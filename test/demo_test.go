package demo_test

import "testing"

type DemoTest struct {
	Name           string
	Param1, Param2 int
	ExpectedResult int
}

func TestSumNumber(t *testing.T) {
	testCases := []DemoTest{
		{
			Name:           "Success 1",
			Param1:         1,
			Param2:         1,
			ExpectedResult: 2,
		},
		{
			Name:           "Success 2",
			Param1:         2,
			Param2:         2,
			ExpectedResult: 4,
		},
		{
			Name:           "Success 3",
			Param1:         100,
			Param2:         100,
			ExpectedResult: 200,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			result := SumNumber(testCase.Param1, testCase.Param2)
			if result != testCase.ExpectedResult {
				t.Errorf("test case %s fail...", t.Name())
			}
		})
	}
}
