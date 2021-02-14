package entities

import "testing"

func TestBadRegistration(t *testing.T) {
	_, e := NewCar("foo", 20)
	if e == nil {
		t.Errorf("Car generation should fail for invalid reg. no. \"foo\"\n")
	}
}

func TestUnderAge(t *testing.T) {
	_, e := NewCar("KA-01-HH-1234", 17)
	if e == nil {
		t.Errorf("Car generation should fail for underage driver\n")
	}
}
