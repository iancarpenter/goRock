package main

import (
	"strings"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	reader := strings.NewReader("rock\n")
	result := getUserInput(reader)
	if result != "rock" {
		t.Errorf("Expected 'rock', but got %s", result)
	}
}

func TestGetComputerInput(t *testing.T) {
	result := getComputerInput()
	choices := []string{"rock", "paper", "scissors"}

	if !strings.Contains(strings.Join(choices, ","), result) {
		t.Errorf("Unexpected result: %s", result)
	}
}

func TestCompareInput(t *testing.T) {
	testCases := []struct {
		user     string
		computer string
		expected string
	}{
		{"rock", "rock", "draw"},
		{"rock", "scissors", "win"},
		{"rock", "paper", "lose"},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := compareInput(tc.user, tc.computer)
		if result != tc.expected {
			t.Errorf("Expected %s, but got %s", tc.expected, result)
		}
	}
}
