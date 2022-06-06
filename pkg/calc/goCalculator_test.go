package calc

import (
	"testing"
)

func TestFindMax(t *testing.T) {
	m := FindMax([]string{"5", "100", "50", "100.1", "2", "99"})
	if m != 100.1 {
		t.Errorf("Expected the maximum value to be 100.1 but got %v", m)
	}
	m = FindMax([]string{"5.5", "5.9", "5.8", "4.4"})
	if m != 5.9 {
		t.Errorf("Expected the maximum value to be 5.9 but got %v", m)
	}
}

func TestFindMin(t *testing.T) {
	m := FindMin([]string{"5", "100", "50", "100.1", "2", "99"})
	if m != 2 {
		t.Errorf("Expected the minimum value to be 2 but got %v", m)
	}
	m = FindMin([]string{"5.5", "5.9", "5.8", "4.4"})
	if m != 4.4 {
		t.Errorf("Expected the minimum value to be 4.4 but got %v", m)
	}
}

func TestFindMean(t *testing.T) {
	m := FindMean([]string{"5", "100", "50", "100", "2", "99"})
	if m != 59.33 {
		t.Errorf("Expected the maximum value to be 100 but got %v", m)
	}
	m = FindMean([]string{"5.5", "5.9", "5.8", "4.4"})
	if m != 5.4 {
		t.Errorf("Expected the mean value to be 5.4 but got %v", m)
	}
}
