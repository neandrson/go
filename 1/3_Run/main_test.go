package main

import (
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	type Test struct {
		args       []string
		expectFail bool
	}

	tests := []Test{
		{
			args:       []string{"10", "20", "30"},
			expectFail: false,
		},
		{
			args:       []string{"-1", "-3", "-6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "-3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"-1", "3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "3", "-6"},
			expectFail: true,
		},
		{
			args:       []string{"0", "3", "6"},
			expectFail: true,
		},
		{
			args:       []string{"1", "0", "6"},
			expectFail: true,
		},
		{
			args:       []string{"44", "3", "0"},
			expectFail: true,
		},
		{
			args:       []string{"5", "5", "lala"},
			expectFail: true,
		},
		{
			args:       []string{"laa", "5", "10"},
			expectFail: true,
		},
		{
			args:       []string{"5", "lala", "10"},
			expectFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, ""), func(t *testing.T) {
			os.Args = []string{"cmd"}
			os.Args = append(os.Args, tt.args...)
			err := run()
			if tt.expectFail && err == nil {
				t.Errorf("Expected run() to fail, but it didn't. args: %v", tt.args)
			}
			if !tt.expectFail && err != nil {
				t.Errorf("Expected run() to succeed, but it failed. args: %v, error: %v", tt.args, err)
			}
		})
	}
}
