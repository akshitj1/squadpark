package entities

type Slot struct {
	// Pos is one-indexed. ie. actual label. not how its strored in lot ds
	Pos int
	Car *Car
}

func NewSlot(pos int) *Slot {
	return &Slot{pos, nil}
}

func (s *Slot) SetCar(car *Car) {
	s.Car = car
}

func (s *Slot) Empty() bool {
	return s.Car == nil
}
