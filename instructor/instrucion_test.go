package instructor

import (
	"errors"
	"reflect"
	"testing"
)

// Test CompileBf test cases
func TestCompileBf(t *testing.T) {
	var testsCases = []struct {
		input          string
		expectedStruct []Instruction
		expect         string
	}{
		{"+[----->+++<]>+.---.+++++++..+++.++++++++.--------.+++.------.--------.",
			[]Instruction{{0, 0}, {6, 12},
				{1, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {3, 0},
				{0, 0}, {0, 0}, {0, 0},
				{2, 0}, {7, 1}, {3, 0},
				{0, 0}, {4, 0}, {1, 0},
				{1, 0}, {1, 0}, {4, 0},
				{0, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {0, 0},
				{0, 0}, {4, 0}, {4, 0},
				{0, 0}, {0, 0}, {0, 0},
				{4, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {0, 0},
				{4, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {1, 0},
				{4, 0}, {0, 0}, {0, 0},
				{0, 0}, {4, 0}, {1, 0},
				{1, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {4, 0},
				{1, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {1, 0},
				{1, 0}, {1, 0}, {4, 0}}, "晨晥晬晬景晷景晲晬晤"},

		{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.",
			[]Instruction{{0, 0}, {0, 0},
				{0, 0}, {0, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {6, 48},
				{3, 0}, {0, 0}, {0, 0},
				{0, 0}, {0, 0}, {6, 33},
				{3, 0}, {0, 0}, {0, 0},
				{3, 0}, {0, 0}, {0, 0},
				{0, 0}, {3, 0}, {0, 0},
				{0, 0}, {0, 0}, {3, 0},
				{0, 0}, {2, 0}, {2, 0},
				{2, 0}, {2, 0}, {1, 0},
				{7, 14}, {3, 0}, {0, 0},
				{3, 0}, {0, 0}, {3, 0},
				{1, 0}, {3, 0}, {3, 0},
				{0, 0}, {6, 45}, {2, 0},
				{7, 43}, {2, 0}, {1, 0},
				{7, 8}, {3, 0}, {3, 0},
				{4, 0}}, "H"}}
	var instruction = Instructor{}

	for _, test := range testsCases {
		instruction.CompileBf(test.input)
		if !reflect.DeepEqual(instruction.Instruct, test.expectedStruct) {
			t.Errorf("got: %v, want: %v.", instruction.Instruct, test.expectedStruct)
		}
		instruction.Instruct = nil
	}

	var testCasesError = []struct {
		input  string
		expect string
	}{{">+++[++].]", errors.New("missing '['").Error()},
		{"++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.[][][.", errors.New("missing ']'").Error()},
		{"", errors.New("input is empty").Error()},
	}

	for _, test := range testCasesError {
		if err := instruction.CompileBf(test.input); err.Error() != test.expect {
			t.Errorf("got: %s, want: %s.", err, test.expect)
		}
	}

}

//Test ExecuteBf test cases
func TestExecuteBf(t *testing.T) {

	var instructor = Instructor{}
	var testCases = []struct {
		input  []Instruction
		expect error
	}{
		{[]Instruction{{0, 0},
			{6, 12}, {1, 0},
			{1, 0}, {1, 0}}, nil},

		{[]Instruction{{0, 0},
			{9, 12}, {1, 0},
			{1, 0}, {1, 0}}, errors.New("unknown operator")},
	}

	instructor.Instruct = testCases[0].input
	if err := instructor.ExecuteBf(); err != nil {
		t.Errorf("%v got: %v want: %v", instructor.Instruct[0].Operate, err, testCases[0].expect)
	}

	instructor.Instruct = testCases[1].input
	if err := instructor.ExecuteBf(); err == nil {
		t.Errorf("%v got: %v want: %v", instructor.Instruct[1].Operate, err, testCases[1].expect)
	}

}
