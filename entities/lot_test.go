package entities

import "testing"

func TestNonPositiveLot(t *testing.T) {
	_, eZero := NewLot(0)
	_, eNeg := NewLot(-1)
	if !(eZero != nil && eNeg != nil) {
		t.Errorf("Lot should not be created for non-positive slots")
	}
}

func TestNumSlots(t *testing.T) {
	l, _ := NewLot(1)
	if l.NumSlots() != 1 {
		t.Error("NumSlots malfunction")
	}
}

func TestGetSlot(t *testing.T) {
	l, _ := NewLot(1)
	_, e := l.GetSlot(1)
	if e == nil {
		t.Error("Shoutld raise error for invalid slot index")
	}
}

func TestSetSlot(t *testing.T) {
	l, _ := NewLot(1)
	e := l.SetSlot(1, nil)
	if e == nil {
		t.Error("Shoutld raise error for invalid slot index")
	}
}
