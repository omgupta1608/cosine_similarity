package cosine_similarity

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func nearlyEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestCosineSimilarity(t *testing.T) {
	tests := []struct {
		name       string
		first      string
		second     string
		ignoreCase bool
		expected   float64
		expectErr  error
	}{
		{
			name:      "Identical strings",
			first:     "hello world",
			second:    "hello world",
			expected:  1.0,
			expectErr: nil,
		},
		{
			name:      "Completely different strings",
			first:     "hello world",
			second:    "goodbye moon",
			expected:  0.0,
			expectErr: nil,
		},
		{
			name:      "Partial overlap",
			first:     "hello world",
			second:    "hello everyone",
			expected:  0.5,
			expectErr: nil,
		},
		{
			name:       "Case sensitivity",
			first:      "Hello World",
			second:     "hello world",
			expected:   1.0,
			expectErr:  nil,
			ignoreCase: true,
		},
		{
			name:       "Case sensitivity - ignore case",
			first:      "Hello World",
			second:     "hello world",
			expected:   0.0,
			expectErr:  nil,
			ignoreCase: false,
		},
		{
			name:       "Case sensitivity - ignore case - partial match",
			first:      "Hello World",
			second:     "Hello world",
			expected:   0.5,
			expectErr:  nil,
			ignoreCase: false,
		},
		{
			name:      "Empty strings",
			first:     "",
			second:    "",
			expected:  -1.0,
			expectErr: ErrInvalidInput,
		},
		{
			name:      "One empty string",
			first:     "hello world",
			second:    "",
			expected:  -1.0,
			expectErr: ErrInvalidInput,
		},
		{
			name:      "Same words in different order",
			first:     "world hello",
			second:    "hello world",
			expected:  1.0,
			expectErr: nil,
		},
		{
			name:      "Repetitive words",
			first:     "hello hello world",
			second:    "hello world world",
			expected:  0.8,
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Cosine(tt.first, tt.second, tt.ignoreCase)
			fmt.Printf("%s: %.2f\n\n", tt.name, result)
			if !nearlyEqual(result, tt.expected, 0.01) || !errors.Is(err, tt.expectErr) {
				t.Errorf("%s: expected %.2f with error %v, got %.2f with error %v", tt.name, tt.expected, tt.expectErr, result, err)
			}
		})
	}
}
