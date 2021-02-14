package entities

import "testing"

func TestEmptySlot(t *testing.T) {
	s := NewSlot(1)
	if !s.Empty() {
		t.Error("Expecting empty")
	}
	c, _ := NewCar("KA-01-HH-1234", 18)
	s.SetCar(c)
	if s.Empty() {
		t.Error("Expecting Not empty")
	}
}
