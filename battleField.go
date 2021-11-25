package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SplitIntoGroups(heroes *[]Hero) map[int]int {
	result := make(map[int]int)
	contains := make([]bool, len(*heroes), len(*heroes))
	rand.Seed(time.Now().UnixNano())

	for len(result) != len(*heroes)/2 {
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
		} else {
			contains[secondFighter] = true
		}

		result[firstFighter] = secondFighter
	}
	return result
}

func CreatePair(firstFighter, secondFighter int, heroes []Hero, first, second chan Hero) {
	first <- heroes[firstFighter]
	second <- heroes[secondFighter]
}

func Print(heroes []Hero) {
	for _, hero := range heroes {
		fmt.Println("ID:", hero.getId(), " NAME:", hero.getName(),
			"HEALTH:", hero.getHealth(), "DAMAGE:", hero.getDamage())
	}
}

func Fight() {
	numberOfHeroes := 12
	heroes := make([]Hero, numberOfHeroes)
	survivingHeroes := make([]Hero, numberOfHeroes)
	FillHeroesArray(numberOfHeroes, &heroes)
	Print(heroes)

	groups := SplitIntoGroups(&heroes)

	fmt.Println(groups)
	firstFighter := make(chan Hero)
	secondFighter := make(chan Hero)
	winner := make(chan Hero)

	for len(groups) >= 1 {
		survivingHeroes = nil
		for key, value := range groups {
			go CreatePair(key, value, heroes, firstFighter, secondFighter)
			go Duel(firstFighter, secondFighter, winner)
		}
		numberOfHeroes /= 2
		for i := 0; i < numberOfHeroes; i++ {
			survivingHeroes = append(survivingHeroes, <-winner)
		}
		heroes = survivingHeroes
		groups = SplitIntoGroups(&heroes)
		fmt.Println(groups)
		Print(heroes)
	}
	fmt.Println("TOURNAMENT WINNER: ", heroes[0].getName())
}

func Duel(firstFighter, secondFighter, winner chan Hero) {
	for first := range firstFighter {
		for second := range secondFighter {
			fmt.Println("Fight between ", first.getName(), "and ", second.getName())

			for first.getHealth() >= 0 && second.getHealth() >= 0 {
				first.attack(second)
				if second.getHealth() <= 0 {
					winner <- first
					fmt.Println("Duel Winner: ", first.getName())
					return
				}
				second.attack(first)
			}
			winner <- second
			fmt.Println("Duel Winner: ", second.getName())
		}
	}
}
