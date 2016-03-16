package goatparse

import (
	"testing"
	"fmt"
	"strings"
)

func TestSimpleParse(t *testing.T) {
	input := "3 minutes - 4 days that i need to go to the store"
	result, _ := ParseDurationFromNow(input)
	fmt.Printf("%v: the rest: %v\n", result, strings.TrimSpace(input[result.Offset:]))
}

func TestBadParse(t *testing.T) {
	input := "3 dingdong"
	_, err := ParseDurationFromNow(input)
	if err == nil {
		t.Error("expected an error")
	}
}
