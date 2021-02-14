package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/akshitj1/squadpark/actions"
)

// func main() {
// 	app, _ := actions.NewSquadParkApp(6)
// 	app.Park("KA-01-HH-1234", 21)
// 	app.Park("PB-01-HH-1234", 21)
// 	app.QuerySlotsByAge(21)
// 	app.QueryRegNumsByAge(21)
// 	app.Park("PB-01-TG-2341", 40)
// 	app.QuerySlotByRegNum("PB-01-HH-1234")
// 	app.Vacate(2)
// 	app.Park("HR-29-TG-3098", 39)
// 	app.QueryRegNumsByAge(18)
// }

func ParseCreateParkingLot(cmd string) *actions.SquadParkApp {
	action := strings.Split(cmd, " ")
	if !(len(action) == 2 && action[0] == "Create_parking_lot") {
		log.Fatalf("Bad command: %v", cmd)
	}
	numSlots, e := strconv.Atoi(action[1])
	if e != nil {
		log.Fatal(e)
	}
	app, e := actions.NewSquadParkApp(numSlots)
	if e != nil {
		log.Fatal(e)
	}
	return app
}

func ParsePark(action []string, app *actions.SquadParkApp) {
	if len(action) != 4 {
		log.Fatalf("Bad command: %v", action)
	}
	regNum := action[1]
	age, e := strconv.Atoi(action[3])
	if e != nil {
		log.Fatal(e)
	}
	e = app.Park(regNum, age)
	if e != nil {
		log.Fatal(e)
	}
}

func ParseLeave(action []string, app *actions.SquadParkApp) {
	if len(action) != 2 {
		log.Fatalf("Bad command: %v", action)
	}
	slot, e := strconv.Atoi(action[1])
	if e != nil {
		log.Fatal(e)
	}
	e = app.Vacate(slot)
	if e != nil {
		log.Fatal(e)
	}
}

func ParseSlotByAge(action []string, app *actions.SquadParkApp) {
	if len(action) != 2 {
		log.Fatalf("Bad command: %v", action)
	}
	age, e := strconv.Atoi(action[1])
	if e != nil {
		log.Fatal(e)
	}
	app.QuerySlotsByAge(age)
}

func ParseRegByAge(action []string, app *actions.SquadParkApp) {
	if len(action) != 2 {
		log.Fatalf("Bad command: %v", action)
	}
	age, e := strconv.Atoi(action[1])
	if e != nil {
		log.Fatal(e)
	}
	app.QueryRegNumsByAge(age)
}

func ParseSlotByReg(action []string, app *actions.SquadParkApp) {
	if len(action) != 2 {
		log.Fatalf("Bad command: %v", action)
	}
	regNum := action[1]
	app.QuerySlotByRegNum(regNum)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: exec input_file")
	}
	fName := os.Args[1]
	file, e := os.Open(fName)
	if e != nil {
		log.Fatalf("Failed to open file: %v, error: %v", fName, e)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	exists := scanner.Scan()
	if !exists {
		log.Fatal("no commands")
	}
	app := ParseCreateParkingLot(scanner.Text())
	for scanner.Scan() {
		action := strings.Split(scanner.Text(), " ")
		if len(action) < 1 {
			continue
		}
		switch action[0] {
		case "Park":
			ParsePark(action, app)
		case "Leave":
			ParseLeave(action, app)
		case "Slot_numbers_for_driver_of_age":
			ParseSlotByAge(action, app)
		case "Vehicle_registration_number_for_driver_of_age":
			ParseRegByAge(action, app)
		case "Slot_number_for_car_with_number":
			ParseSlotByReg(action, app)
		}
	}
}
