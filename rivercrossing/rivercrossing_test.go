package rivercrossing

import (
	"sort"
	"testing"
)

func TestWorldIllegalState(t *testing.T) {

	world := CreateWorld()

	PutBoat("Rev")
	PutBoat("HS")
	CrossRiver()
	PopBoat("Rev")
	CrossRiver()
	PutBoat("Korn")
	CrossRiver()
	PopBoat("Korn")

	state, legal := isWorldStateLegal(world)

	if !legal {
		t.Errorf("World state was illegal! \nState error returned: %v \n \n World State:  %v", state, GetWorldStateString())
	}
}

func TestWorldLegalState(t *testing.T) {
	world := CreateWorld()
	PutBoat("Kylling")
	PutBoat("HS")
	CrossRiver()
	PopBoat("Kylling")
	CrossRiver()
	PutBoat("Korn")
	CrossRiver()
	PopBoat("Korn")
	_, legal := isWorldStateLegal(world)

	if legal {
		t.Errorf("Illegal state was deemed legal! \n \n World State:  %v", GetWorldStateString())
	}
}

func TestPutBoat(t *testing.T) {
	world = CreateWorld()
	PutBoat("Rev")

	if !contains("Rev", world.boat.passengers) {
		t.Errorf("Rev was not found in boat! \n \n World State: %v", GetWorldStateString())
	}
}

func TestPopBoat(t *testing.T) {
	//default posisjon for båten er vest, så vi antar at kyllingen vil være på øst
	world := CreateWorld()

	PutBoat("Kylling")
	CrossRiver()
	PopBoat("Kylling")

	if !contains("Kylling", world.east.occupants) {
		t.Errorf("Kylling was not found in the east! \n \n World State: %v", GetWorldStateString())
	}
}

func TestAll(t *testing.T) {
	t.Run("WORLD_ILLEGAL_STATE", TestWorldIllegalState)
	t.Run("WORLD_LEGAL_STATE", TestWorldLegalState)
	t.Run("BOAT_PUT", TestPutBoat)
	t.Run("BOAT_POP", TestPopBoat)
}

// Tetser på om landstatus er ulovlig, og returnerer status
func isWorldStateLegal(world World) (string, bool) {
	westState, westLegal := isLandStateLegal(*world.west)
	eastState, eastLegal := isLandStateLegal(*world.east)
	boatState, boatLegal := isLandStateLegal(Land{"boat", world.boat.passengers}) // veldig hacky vei å sjekke om båt state er lovlig

	if !westLegal {
		return westState, westLegal
	} else if !eastLegal {
		return eastState, eastLegal
	} else if !boatLegal {
		return boatState, boatLegal
	} else {
		return "", true
	}
}

// Sjekker om landtilstanden er lovlig
func isLandStateLegal(land Land) (string, bool) {
	worldLand := land.occupants
	// Om HS er på land så betyr det at andre beboere ikke kan skade hverandre, dermed sjekker vi om HS ikke er på land
	if !contains("HS", worldLand) && len(worldLand) > 1 {
		// sorter beboere slik at de med høyest styrke kommer først i slice
		sort.Slice(worldLand, func(i, j int) bool {
			return worldLand[j].strength < worldLand[i].strength
		})
	}
	if len(worldLand) == 1 {
		return "", true
	}
	for i, occupant := range worldLand {
		// Passe på at vi ikke prøver å hente index som er utenfor rekkevidde av slice nb7jhuhujin
		if i < len(worldLand) {
			// Om beboer er rovdyr og beboer sin styrke er større en neste beboer i array, og at beboer sin styrke er høyere enn to (styrke 3 kan ikke spise styrke 1)
			if occupant.predator && occupant.strength > worldLand[i+1].strength && (occupant.strength-worldLand[i+1].strength) == 1 {
				return worldLand[0].name + " can eat " + worldLand[i+1].name, false
			}
		}

	}
	return "", true
}

// Sjekker om slice inneholder en beboer
func contains(name string, occupants []Occupant) bool {
	_, err := GetOccupantByName(name, occupants)
	if err != nil {
		return false
	}
	return true
}
