package main

import rc "github.com/tobiaskrok/rivercrossing/rivercrossing"

func main() {

	rc.CreateWorld()

	rc.PrintWorld()
	rc.PutBoat("HS")
	rc.PutBoat("Kylling")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PopBoat("Kylling")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PutBoat("Korn")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PopBoat("Korn")
	rc.PutBoat("Kylling")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PopBoat("Kylling")
	rc.PutBoat("Rev")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PopBoat("Rev")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PutBoat("Kylling")
	rc.PrintWorld()

	rc.CrossRiver()
	rc.PrintWorld()

	rc.PopBoat("Kylling")
	rc.PopBoat("HS")
	rc.PrintWorld()
}
