package patient

import "testing"

func TestValidateName(t *testing.T) {
	tests := []struct {
		desc     string
		name     string
		expected bool
	}{
		{
			desc:     "Case1",
			name:     "aakanksha",
			expected: true,
		},
		{
			desc:     "Case2",
			name:     "",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isValid := validatename(test.name)
			if isValid != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, isValid)
			}
		})
	}
}

func TestValidateId(t *testing.T) {
	tests := []struct {
		desc   string
		input  int
		output bool
	}{
		{
			desc:   "Case1",
			input:  1,
			output: true,
		},
		{
			desc:   "Case2",
			input:  -1,
			output: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isValid := validId(test.input)
			if isValid != test.output {
				t.Errorf("Expected: %v, Got: %v", test.output, isValid)
			}
		})
	}
}

