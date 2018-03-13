package main

import "strings"

const placeNameKitchen = "кухня"
const placeNameOutside = "улица"
const placeNameCorridor = "коридор"
const placeNameRoom = "комната"

const itemNameBackpack = "рюкзак"
const itemNameKeys = "ключи"
const itemNameSynopses = "конспекты"

const actionNameLookAround = "осмотреться"
const actionNameMove = "идти"
const actionNameEquip = "одеть"
const actionNameTake = "взять"
const actionNameUse = "применить"

type playerStruct struct {
	place    *place
	equipped []equippedItem
	quest    quest
}

type quest struct {
	takeBackpack   bool
	goToUniversity bool
}

type equippedItem struct {
	name       string
	containing []item
}

type item struct {
	name      string
	equipable bool
	container bool
}

type place struct {
	name              string
	containing        []item
	neighboringPlaces []*place
}

type playerCommand struct {
	action string
	arg1   string
	arg2   string
}

var player playerStruct

var actions map[string]actionFunction

func initGame() {
	itemBackpack := item{itemNameBackpack, true, true}
	itemKeys := item{itemNameKeys, false, false}
	itemSynopses := item{itemNameSynopses, false, false}

	placeCorridor := place{placeNameCorridor, []item{}, []*place{}}
	placeKitchen := place{placeNameKitchen, []item{}, []*place{}}
	placeOutside := place{placeNameOutside, []item{}, []*place{}}
	placeRoom := place{placeNameRoom, []item{itemKeys, itemSynopses, itemBackpack}, []*place{}}

	placeCorridor.neighboringPlaces = []*place{&placeKitchen, &placeRoom, &placeOutside}
	placeKitchen.neighboringPlaces = []*place{&placeCorridor}
	placeOutside.neighboringPlaces = []*place{&placeCorridor}
	placeRoom.neighboringPlaces = []*place{&placeCorridor}

	player = playerStruct{
		&placeKitchen,
		[]equippedItem{},
		quest{true, true},
	}

	actions = map[string]actionFunction{
		actionNameLookAround: lookAround,
		actionNameMove:       move,
		actionNameEquip:      equip,
		actionNameTake:       take,
		actionNameUse:        use,
	}
}

func handleCommand(command string) string {
	var parsedCommand = parseCommand(command)

	actionFunction, ok := actions[parsedCommand.action]

	if !ok {
		return "неизвестная команда"
	}

	return actionFunction(parsedCommand)
}

func parseCommand(inputString string) (parsedCommand playerCommand) {
	var args = strings.Split(inputString, " ")
	for name := range actions {
		if name == args[0] {
			parsedCommand.action = args[0]
			if len(args) > 1 {
				parsedCommand.arg1 = args[1]
			}
			if len(args) > 2 {
				parsedCommand.arg2 = args[2]
			}
		}
	}
	return
}
func main() {
	initGame()
	println(handleCommand("осмотреться"))
	println(handleCommand("идти коридор"))
	println(handleCommand("идти комната"))
}
