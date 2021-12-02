package main

import (
	"testing"
)

func TestDay02aSolution(t *testing.T) {
	if Day02aSolution() != 150 {
		t.Errorf("Wrong solution. Expected 150, got %v\n", Day02aSolution())
	}

}
