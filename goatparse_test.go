package goatparse

import (
	"testing"
	"time"
)

func TestSimpleParse(t *testing.T) {
	input := "3 minutes - 4 days that i need to go to the store"
	result, _ := ParseDuration(input)
	expected := 3*time.Minute - 4*24*time.Hour
	if result.Duration != expected {
		t.Error("expected ", expected, " but received ", result)
	}
}

func TestBadParse(t *testing.T) {
	input := "3 dingdong"
	_, err := ParseDuration(input)
	if err == nil {
		t.Error("expected an error")
	}
}
