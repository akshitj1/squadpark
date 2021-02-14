package entities

import (
	"errors"

	"github.com/akshitj1/squadpark/util"
)

type Lot []*Slot

func NewLot(numParkingSlots int) (*Lot, error) {
	if numParkingSlots <= 0 {
		return nil, errors.New("Malformed parking slots")
	}
	var slots []*Slot
	for slotIdx := 0; slotIdx < numParkingSlots; slotIdx++ {
		slotPos := util.ToOneIndex(slotIdx)
		slots = append(slots, NewSlot(slotPos))
	}
	lot := Lot(slots)
	return &lot, nil
}

// Empty car ie. nil is valid choice
func (l *Lot) SetSlot(slotIdx int, car *Car) error {
	if !l.ValidSlot(slotIdx) {
		return errors.New("Bad slot")
	}
	(*l)[slotIdx].SetCar(car)
	return nil
}

func (l *Lot) ValidSlot(slotIdx int) bool {
	return slotIdx >= 0 && slotIdx < l.NumSlots()
}

func (l *Lot) GetSlot(slotIdx int) (*Slot, error) {
	if !l.ValidSlot(slotIdx) {
		return nil, errors.New("Bad slot")
	}
	return (*l)[slotIdx], nil
}

func (l *Lot) NumSlots() int {
	return len(*l)
}
