package main

import "testing"

func TestGetSum(t *testing.T) {
	got := GetSum(1, 2)
	if got != 3 {
		t.Errorf("GetSum(1, 2) = %d; want 3", got)
	}
}
