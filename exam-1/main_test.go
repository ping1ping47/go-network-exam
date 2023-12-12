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

func TestAdd(t *testing.T) {
    expected := 5
    if got := Add(2, 3); got != expected {
        t.Errorf("Add(2, 3) = %v, want %v", got, expected)
    }
}

func TestMinus(t *testing.T) {
    expected := 1
    if got := Minus(3, 2); got != expected {
        t.Errorf("Minus(3, 2) = %v, want %v", got, expected)
    }
}

func TestUserLogin(t *testing.T) {
    expected := true
    if got := UserLogin("admin", "password"); got != expected {
        t.Errorf("UserLogin(\"admin\", \"admin\") = %v, want %v", got, expected)
    }
}