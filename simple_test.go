package main

import "testing"

func TestSimplego(t *testing.T) {
	if Simple() != 1 {
		t.Error("Expected 1")
	}
}
