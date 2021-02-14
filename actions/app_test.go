package actions

import (
	"testing"

	"github.com/akshitj1/squadpark/util"
)

func TestVacantSlot(t *testing.T) {
	app, _ := NewSquadParkApp(2)
	app.Park("KA-01-HH-1234", 21)
	app.Park("KA-01-HH-1235", 21)
	app.Vacate(1)
	s, _ := app.emptySlots.GetEmpty()
	if !(s == util.ToZeroIndex(1)) {
		t.Error("Vacant slot malfunction")
	}
}

func TestParkOverflow(t *testing.T) {
	app, _ := NewSquadParkApp(2)
	app.Park("KA-01-HH-1234", 21)
	app.Park("KA-01-HH-1235", 21)
	e := app.Park("KA-01-HH-1236", 21)
	if e == nil {
		t.Error("Expecting Parking reject as lot is full")
	}
}

func EqualStringArray(got []string, exp []string) bool {
	if len(got) != len(got) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if !(got[i] == exp[i]) {
			return false
		}
	}
	return true
}

func EqualIntArray(got []int, exp []int) bool {
	if len(got) != len(got) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if !(got[i] == exp[i]) {
			return false
		}
	}
	return true
}

// we assume that query result is ordered by slot id
func TestQueryRegNumsByAge(t *testing.T) {
	app, _ := NewSquadParkApp(3)
	app.Park("KA-01-HH-1234", 21)
	app.Park("KA-01-HH-1235", 21)
	app.Park("KA-01-HH-1236", 22)
	got := app.QueryRegNumsByAge(21)
	exp := []string{"KA-01-HH-1234", "KA-01-HH-1235"}
	if !EqualStringArray(got, exp) {
		t.Errorf("exp: %v, got %v", exp, got)
	}
}

func TestQuerySlotsByAge(t *testing.T) {
	app, _ := NewSquadParkApp(3)
	app.Park("KA-01-HH-1234", 21)
	app.Park("KA-01-HH-1235", 21)
	app.Park("KA-01-HH-1236", 22)
	got := app.QuerySlotsByAge(21)
	exp := []int{1, 2}
	if !EqualIntArray(got, exp) {
		t.Errorf("exp: %v, got %v", exp, got)
	}
}

func TestQuerySlotByRegNum(t *testing.T) {
	app, _ := NewSquadParkApp(3)
	app.Park("KA-01-HH-1234", 21)
	app.Park("KA-01-HH-1235", 21)
	app.Park("KA-01-HH-1236", 22)
	got := app.QuerySlotByRegNum("KA-01-HH-1236")
	exp := 3
	if !(got == exp) {
		t.Errorf("exp: %v, got %v", exp, got)
	}
	got = app.QuerySlotByRegNum("KA-01-HH-1237")
	exp = 0
	if !(got == exp) {
		t.Errorf("exp: %v, got %v", exp, got)
	}
}
