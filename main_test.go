package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	t.Log("test executed")

	got := getMessage()
	want := "Hello World!"

	if got != want {
		t.Fatalf("Got: %v Want: %v", got, want)
	}
}
