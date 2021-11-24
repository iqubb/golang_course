package main

import (
	"math/rand"
	"strconv"
	"time"
)

func FillHeroesArray(numberOfHeroes int, heroes *[]Hero) {
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	numberOfWarriors, numberOfWizards, numberOfArchers := 0, 0, 0
	for i := 0; i < numberOfHeroes; i++ {
		heroType := number.Intn(3);
		if heroType == 0 {
			(*heroes)[i] = &Wizard{id: uint64(i + 1), name: "wizard(" + strconv.Itoa(numberOfWizards) + ")", health: 250, damage: 70, shield: 30}
			numberOfWizards++
		} else if heroType == 1 {
			(*heroes)[i] = &Warrior{id: uint64(i + 1), name: "warrior(" + strconv.Itoa(numberOfWarriors) + ")",health: 500, damage: 30, shield: 70}
			numberOfWarriors++
		} else {
			(*heroes)[i] = &Archer{id: uint64(i + 1), name: "archer(" +strconv.Itoa(numberOfArchers) + ")", health: 350, damage: 50, shield: 50}
			numberOfArchers++
		}
	}
}
