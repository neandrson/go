package main

import (
	"io"
	"testing"
)

func TestUpper(t *testing.T) {
	upper := UpperWriter{
		UpperString: "any",
	}

	w := io.Writer(&upper)

	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedLen    uint
		expectedError  bool
	}{
		{
			name:           "simple",
			input:          "data",
			expectedOutput: "DATA",
			expectedLen:    4,
			expectedError:  false,
		},
		{
			name:           "same",
			input:          "DDD",
			expectedOutput: "DDD",
			expectedLen:    3,
			expectedError:  false,
		},
		{
			name:           "empty",
			input:          "",
			expectedOutput: "",
			expectedLen:    0,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			written, err := w.Write([]byte(tt.input))
			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected write error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if written != int(tt.expectedLen) {
				t.Errorf("Expected write error %d, but got none%d", written, tt.expectedLen)
			}
		})
	}
}
