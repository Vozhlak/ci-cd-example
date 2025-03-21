package main

import (
	"os"
	"testing"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}
