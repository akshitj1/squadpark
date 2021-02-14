package actions

import (
	"errors"
	"fmt"

	"github.com/akshitj1/squadpark/entities"
	"github.com/akshitj1/squadpark/util"
)

type SquadParkApp struct {
	emptySlots *EmptySlots
	lot        *entities.Lot
}

func NewSquadParkApp(numParkingSlots int) (*SquadParkApp, error) {
	lot, e := entities.NewLot(numParkingSlots)
	if e != nil {
		return nil, e
	}
	s := &SquadParkApp{NewEmptySlots(numParkingSlots), lot}
	fmt.Printf("Created parking with %v slots\n", numParkingSlots)
	return s, nil
}

func (app *SquadParkApp) Park(regNum string, age int) error {
	car, e := entities.NewCar(regNum, age)
	if e != nil {
		return e
	}
	slotIdx, e := app.emptySlots.GetEmpty()
	if e != nil {
		return e
	}
	app.lot.SetSlot(slotIdx, car)
	slot, _ := app.lot.GetSlot(slotIdx)
	fmt.Printf("Car with vehicle registration number \"%v\" has been parked at slot number %v\n", slot.Car.RegNum, slot.Pos)
	return nil
}

func (app *SquadParkApp) Vacate(slot int) error {
	slot = util.ToZeroIndex(slot)
	if !app.lot.ValidSlot(slot) {
		return errors.New("malformed slot")
	}
	availSlot, e := app.lot.GetSlot(slot)
	if e != nil {
		return e
	}
	if availSlot.Empty() {
		return errors.New("trying to vacate empty slot")
	}
	car := availSlot.Car
	app.lot.SetSlot(slot, nil)
	app.emptySlots.SetEmpty(slot)
	fmt.Printf("Slot number %v vacated, the car with vehicle registration number \"%v\" left the space, the driver of the car was of age %v\n", availSlot.Pos, car.RegNum, car.Age)
	return nil
}

func (app *SquadParkApp) filterSlotsByAge(age int) []*entities.Slot {
	filSlots := make([]*entities.Slot, 0)
	for slotIdx := 0; slotIdx < app.lot.NumSlots(); slotIdx++ {
		slot, e := app.lot.GetSlot(slotIdx)
		if e != nil {
			panic(e)
		}
		if slot.Empty() {
			continue
		}
		if slot.Car.Age == age {
			filSlots = append(filSlots, slot)
		}
	}
	return filSlots
}

func (app *SquadParkApp) filterSlotByRegNum(regNum string) *entities.Slot {
	for slotIdx := 0; slotIdx < app.lot.NumSlots(); slotIdx++ {
		slot, e := app.lot.GetSlot(slotIdx)
		if e != nil {
			panic(e)
		}
		if slot.Empty() {
			continue
		}
		if slot.Car.RegNum == regNum {
			return slot
		}
	}
	return nil
}

func (app *SquadParkApp) QuerySlotsByAge(age int) []int {
	filSlots := app.filterSlotsByAge(age)
	var resSlots []int
	for _, slot := range filSlots {
		resSlots = append(resSlots, slot.Pos)
	}
	fmt.Println(resSlots)
	return resSlots
}

func (app *SquadParkApp) QueryRegNumsByAge(age int) []string {
	filSlots := app.filterSlotsByAge(age)
	var resRegNums []string
	for _, slot := range filSlots {
		resRegNums = append(resRegNums, slot.Car.RegNum)
	}
	fmt.Println(resRegNums)
	return resRegNums
}

// will return zero age for non-existent reg num
func (app *SquadParkApp) QuerySlotByRegNum(regNum string) int {
	slot := app.filterSlotByRegNum(regNum)
	if slot == nil {
		fmt.Printf("Vehicle with registration number \"%v\", not in lot\n", regNum)
		return 0
	}
	fmt.Printf("Vehicle with registration number \"%v\", is parked in slot %v\n", regNum, slot.Pos)
	return slot.Pos
}
