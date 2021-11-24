package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SplitIntoGroups(heroes *[]Hero) map[int]int{
	result := make(map[int]int)
	contains := make([]bool, len(*heroes), len(*heroes))
	rand.Seed(time.Now().UnixNano())
	for len(result) != len(*heroes) / 2 {
		firstFighter := rand.Intn(len(*heroes))
		if contains[firstFighter] {
			for contains[firstFighter] {
				firstFighter = rand.Intn(len(*heroes))
			}
			contains[firstFighter] = true
		} else {
			contains[firstFighter] = true
		}
		secondFighter := rand.Intn(len(*heroes))
		if contains[secondFighter] {
			for contains[secondFighter] {
				secondFighter = rand.Intn(len(*heroes))
			}
			contains[secondFighter] = true
		}
		result[firstFighter] = secondFighter
	}
	return result
}

func Print(heroes []Hero) {
	for _, hero := range heroes {
		fmt.Println("ID:", hero.getId(), " NAME:", hero.getName(),
			"HEALTH:", hero.getHealth(), "DAMAGE:", hero.getDamage())
	}
}

func Fight() {
	numberOfHeroes := 12
	heroes:= make([]Hero, numberOfHeroes)
	FillHeroesArray(numberOfHeroes, &heroes)
	Print(heroes)
	groups := SplitIntoGroups(&heroes)
	fmt.Println(groups)

	//firstFighter := make(chan Hero)
	//secondFighter := make(chan Hero)
	//winner := make(chan Hero)

}

