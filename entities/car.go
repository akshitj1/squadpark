package entities

import "errors"

type Car struct {
	RegNum string
	Age    int
}

func NewCar(regNum string, age int) (*Car, error) {
	c := &Car{}
	e := c.setAge(age)
	if e != nil {
		return c, e
	}
	e = c.setRegistrationNumber(regNum)
	if e != nil {
		return c, e
	}
	return c, nil
}

func (c *Car) setRegistrationNumber(regNum string) error {
	if len(regNum) != 13 {
		return errors.New("Malformed registration number")
	}
	c.RegNum = regNum
	return nil
}

func (c *Car) setAge(age int) error {
	if age < 18 {
		return errors.New("Underage Driver")
	}
	c.Age = age
	return nil
}
