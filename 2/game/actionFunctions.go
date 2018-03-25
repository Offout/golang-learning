package main

type actionFunction func(command playerCommand) string

func lookAround(_ playerCommand) (reply string) {
	if player.place.name == placeNameKitchen {
		reply = "ты находишься на кухне, на столе чай,"
	}

	if len(player.place.containing) > 0 {
		var length, i = len(player.place.containing), 0
		var table, chair = false, false

		for _, val := range player.place.containing {
			i++

			if val.container && !chair {
				chair = true
				reply += "на стуле - " + val.name
			} else if !val.container && !table {
				table = true
				reply += "на столе: " + val.name
			} else {
				reply += val.name
			}

			if i != length {
				reply += ", "
			}
		}
	} else if player.place.name == placeNameRoom {
		reply += "пустая комната"
	}

	if player.place.name == placeNameKitchen {
		reply += " надо"

		if player.quest.takeBackpack && player.quest.goToUniversity {
			reply += " собрать рюкзак и идти в универ"
		} else if player.quest.takeBackpack {
			reply += " собрать рюкзак"
		} else if player.quest.goToUniversity {
			reply += " идти в универ"
		}
	}

	reply += canMoveToSubstring(*player.place)

	return
}

func move(command playerCommand) (reply string) {

	var moveTo place
	for _, v := range player.place.neighboringPlaces {
		if command.arg1 == v.name {
			moveTo = *v
		}
	}

	if moveTo.name == "" {
		return "нет пути в " + command.arg1
	}

	for _, v := range doors {
		if (moveTo.name == v.to.name) && (player.place.name == v.from.name) && v.locked {
			return "дверь закрыта"
		}
	}

	if moveTo.name == placeNameOutside {
		return "на улице весна. можно пройти - домой"
	}

	if moveTo.name == placeNameKitchen {
		reply += "кухня, "
	}

	if len(moveTo.containing) == 0 && moveTo.name != placeNameOutside {
		reply += "ничего интересного"
	} else {
		if moveTo.name == placeNameRoom {
			reply += "ты в своей комнате"
		}
		if moveTo.name == placeNameKitchen {
			reply += "кухня"
		}
	}

	reply += canMoveToSubstring(moveTo)

	player.place = &moveTo

	return
}

func take(command playerCommand) (reply string) {
	var toTake, deleteIndex = findItem(*player.place, command.arg1)

	if toTake.name == "" {
		return "нет такого"
	}

	if toTake.container {
		reply += "нельзя взять эту вещь, только одеть"
	} else {
		successfullyTaken := false

		for idx, val := range player.equipped {
			if val.name != "" {
				successfullyTaken = true
				player.equipped[idx].containing = append(val.containing, toTake)
				player.place.containing = append(player.place.containing[:deleteIndex], player.place.containing[deleteIndex+1:]...)
			}
		}

		if !successfullyTaken {
			return "некуда класть"
		}

		reply += "предмет добавлен в инвентарь: " + toTake.name
	}

	return
}

func equip(command playerCommand) (reply string) {

	var toTake, deleteIndex = findItem(*player.place, command.arg1)

	if toTake.name == "" {
		return "нет такого"
	}

	if toTake.container {
		reply += "вы одели: " + toTake.name

		equippedItem := equippedItem{toTake.name, []item{}}
		player.equipped = append(player.equipped, equippedItem)

		player.place.containing = append(player.place.containing[:deleteIndex], player.place.containing[deleteIndex+1:]...)
		if toTake.name == itemNameBackpack {
			player.quest.takeBackpack = false
		}
	} else {
		reply += "нельзя одеть эту вещь, только взять"
	}

	return
}

func use(command playerCommand) (reply string) {
	usedItem := command.arg1
	object := command.arg2

	foundItem := item{"", false, false}

SearchItemLoop:
	for _, container := range player.equipped {
		if container.name != "" {
			for _, val := range container.containing {
				if val.name == usedItem {
					foundItem = val
					break SearchItemLoop
				}
			}
		}
	}

	if foundItem.name == "" {
		return "нет предмета в инвентаре - " + usedItem
	}

	if object != "дверь" {
		return "не к чему применить"
	}

	for idx, door := range doors {
		if door.from.name == player.place.name {
			doors[idx].locked = false
			return "дверь открыта"
		}
	}

	return
}

func canMoveToSubstring(place place) (reply string) {
	reply += ". можно пройти - "
	var length, i = len(place.neighboringPlaces), 0
	for _, val := range place.neighboringPlaces {
		i++
		reply += val.name
		if i != length {
			reply += ", "
		}
	}
	return
}

func findItem(place place, itemName string) (item item, deleteKey int) {
	if len(place.containing) > 0 {
		for key, val := range player.place.containing {
			if val.name == itemName {
				item = val
				deleteKey = key
			}
		}
	}
	return
}
