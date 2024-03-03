package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}
