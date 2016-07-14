package goutils

import (
	"testing"
	"github.com/brunetto/goutils"
	"fmt"
)

func TestLeftPad(*testing.T) {
	expected := "000LeftPad"
	actual := goutils.LeftPad("LeftPad", "0", 10)
	if actual != expected {
		fmt.Errorf("Test failed, expected: '%s', got '%s'\n", expected, actual)
	}
}

func TestRightPad(*testing.T) {
	expected := "RightPad000"
	actual := goutils.RightPad("RightPad", "0", 11)
	if actual != expected {
		fmt.Errorf("Test failed, expected: '%s', got '%s'\n", expected, actual)
	}
}
