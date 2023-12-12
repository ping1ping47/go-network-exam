package main

import "testing"

func TestExampleFunction(t *testing.T) {
    // TODO: Implement tests for the student's code
    // Example test (adjust according to the actual assignment)
    expected := "Hello"
    if got := HelloFunction(); got != expected {
        t.Errorf("HelloFunction() = %v, want %v", got, expected)
    }
}
