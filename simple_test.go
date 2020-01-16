package main

import "testing"

func TestSimplego(t *testing.T) {
	if simple() != 1 {
		t.Error("Expected 1")
	}
}
