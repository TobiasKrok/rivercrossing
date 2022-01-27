package rivercrossing

import (
	"errors"
	"fmt"
)

// Oppsett av datastrukturer for verden
type Land struct {
	direction string
	occupants []Occupant
}
type Occupant struct {
	name     string
	predator bool // Definerer om beboer vil spise svakere beboere eller ikke
	strength int  // Definerer "næringskjeden". En beboer med høyest styrke spiser ikke de med lavest styrke
}

type Boat struct {
	position   *Land
	passengers []Occupant
}

type World struct {
	east *Land
	west *Land
	boat *Boat
}

// Verdensvariabler
var occupants []Occupant
var west = Land{}
var east = Land{}
var boat = Boat{}
var world = World{}

// Oppretter en ny verden
func CreateWorld() World {
	occupants = []Occupant{{"HS", false, 3}, {"Rev", true, 2}, {"Kylling", true, 1}, {"Korn", false, 0}}
	west = Land{"west", occupants}
	east = Land{"east", []Occupant{}}
	boat = Boat{&west, []Occupant{}}
	world = World{east: &east, west: &west, boat: &boat}
	return world
}

// Printer ut verdentilstanden
func PrintWorld() {
	fmt.Println(GetWorldStateString())
}

// Bygger tekststreng som viser verdenstilstanden
func GetWorldStateString() string {
	westLandStr := fmt.Sprintf("[%v ---W", getOccupantNameString(west.occupants))
	eastLandStr := fmt.Sprintf("E--- %v]", getOccupantNameString(east.occupants))
	boatStr := ""
	if boat.position == &west {
		boatStr = fmt.Sprintf(" \\_%v_/_____________ ", getOccupantNameString(boat.passengers))
	} else {
		boatStr = fmt.Sprintf(" _____________\\_%v_/ ", getOccupantNameString(boat.passengers))
	}

	worldStr := westLandStr + boatStr + eastLandStr
	return worldStr
}

// Flytter båten over elven
func CrossRiver() {
	if boat.position == &west {
		boat.position = &east
	} else {
		boat.position = &west
	}
}

// PutBoat puts occupant in boat
func PutBoat(occupantName string) {
	land := &boat.position
	occupant, _ := GetOccupantByName(occupantName, (*land).occupants)
	// Fjerner beboer fra land og legger den til i båten
	removeOccupant(occupant, &(*land).occupants)
	boat.passengers = append(boat.passengers, occupant)
}

// PopBoat henter beboer ut fra båten og setter den på land
func PopBoat(occupantName string) {
	land := &boat.position
	occupant, _ := GetOccupantByName(occupantName, boat.passengers)
	removeOccupant(occupant, &boat.passengers)
	(*land).occupants = append((*land).occupants, occupant)
}

// Fjerner beboer fra slice
func removeOccupant(occupant Occupant, occupants *[]Occupant) {
	for i, o := range *occupants {
		if o.name == occupant.name {
			*occupants = append((*occupants)[:i], (*occupants)[i+1:]...)
			return
		}
	}
}

// Lager en tekststreng ut av beboernavn
func getOccupantNameString(occupants []Occupant) string {
	str := ""
	for _, occupant := range occupants {
		str = str + occupant.name + " "
	}
	return str
}

// Finner beboerobjekt ut ifra navn
func GetOccupantByName(name string, occupants []Occupant) (Occupant, error) {
	for _, occupant := range occupants {
		if occupant.name == name {
			return occupant, nil
		}
	}
	return Occupant{}, errors.New("No occupant found with name:" + name)
}
