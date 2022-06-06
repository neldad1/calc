package calc

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	s, _ := Addition([]string{"1", "2", "3", "4", "5"})
	if reflect.TypeOf(s).Kind() != reflect.Float64 {
		t.Errorf("Expected a return tyoe of float64 but got %T", s)
	}
	if s != 15 {
		t.Errorf("Expected a value of 15  but got %v", s)
	}
	s, _ = Addition([]string{"1.6", "2", "3.9", "4", "5"})
	if s != 16.5 {
		t.Errorf("Expected a value of 16.5  but got %v", s)
	}
}

func TestSubtraction(t *testing.T) {
	d, _ := Subtraction("10", "2")
	if reflect.TypeOf(d).Kind() != reflect.Float64 {
		t.Errorf("Expected a return tyoe of float64 but got %T", d)
	}
	if d != 8 {
		t.Errorf("Expected a value of 8  but got %v", d)
	}
	dif, err := Subtraction("a", "b")
	if dif != 0 {
		t.Errorf("Expected a value of 0  but got %v", d)
	}
	if err == nil {
		t.Errorf("Expected to get an error but got nil")
	}
}

func TestMultiplication(t *testing.T) {
	d, _ := Multiplication("10", "2")
	if reflect.TypeOf(d).Kind() != reflect.Float64 {
		t.Errorf("Expected a return tyoe of float64 but got %T", d)
	}
	if d != 20 {
		t.Errorf("Expected a value of 20  but got %v", d)
	}
}

func TestDivision(t *testing.T) {
	d, _ := Division("10.5", "2")
	if reflect.TypeOf(d).Kind() != reflect.Float64 {
		t.Errorf("Expected a return tyoe of float64 but got %T", d)
	}
	if d != 5.25 {
		t.Errorf("Expected a value of 5 but got %v", d)
	}
}
