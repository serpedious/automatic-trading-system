package main

import (
	"testing"
)

func TestPrivateFuncSuccess(t *testing.T) {
	output := hello()
	if output != 1 {
		t.Fatal("failed test")
	}
}