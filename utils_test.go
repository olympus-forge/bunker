package bunker

import (
	"testing"
)

func TestIsEmptyString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"\t\n  \r", true},
		{"hello", false},
		{"  hello  ", false},
	}

	for _, tc := range tests {
		actual := IsEmptyString(tc.input)
		if actual != tc.expected {
			t.Errorf("IsEmptyString(%q) = %v; expected %v", tc.input, actual, tc.expected)
		}
	}
}

func TestLoggers(t *testing.T) {
	// Verify logger functions don't panic
	LogInfo("test info")
	LogWarn("test warn")
	PrintErr("test error")
}
