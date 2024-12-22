package handler

import "testing"

func TestIsUUID(t *testing.T) {
	testcases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "123e4567-e89b-12d3-a456-426655440000",
			expected: true,
		},
		{
			input:    "123e4567-e89b-12d3-a456-42665544000",
			expected: false,
		},
		{
			input:    "abcdef67-daer-12d3-a456-42665544000",
			expected: false,
		},
	}

	for _, tc := range testcases {
		result := IsUUID(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestIsEmail(t *testing.T) {
	testcases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "trplezeor@example.com",
			expected: true,
		},
		{
			input:    "usertrplezeor@example.com",
			expected: true,
		},
		{
			input:    "abc",
			expected: false,
		},
	}

	for _, tc := range testcases {
		result := IsEmail(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
