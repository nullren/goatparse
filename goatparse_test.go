package goatparse

import (
	"testing"
	"time"
)

func TestSimpleParse(t *testing.T) {
	input := "3 minutes - 4 days that i need to go to the store"
	result, _ := ParseDuration(input)
	expected := time.Now().Add(3*time.Minute - 4*24*time.Hour)
	if expected.Sub(result.Time).Seconds() > 1 {
		t.Error("expected ", expected, " but received ", result, " and difference is ", expected.Sub(result.Time).Seconds())
	}
}

func TestBadParse(t *testing.T) {
	input := "3 dingdong"
	_, err := ParseDuration(input)
	if err == nil {
		t.Error("expected an error")
	}
}
