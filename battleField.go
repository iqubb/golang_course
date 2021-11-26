package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FindFighter(contains *[]bool, heroes *[]Hero) int {
	fighter := rand.Intn(len(*heroes))
	if (*contains)[fighter] {
		for (*contains)[fighter] {
			fighter = rand.Intn(len(*heroes))
		}
		(*contains)[fighter] = true
	} else {
		(*contains)[fighter] = true
	}
	return fighter
}

func SplitIntoGroups(heroes *[]Hero) map[int]int {
	result := make(map[int]int)
	contains := make([]bool, len(*heroes), len(*heroes))
	rand.Seed(time.Now().UnixNano())
	for len(result) != len(*heroes)/2 {
		firstFighter := FindFighter(&contains, heroes)
		secondFighter := FindFighter(&contains, heroes)
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

	}
	fmt.Println("TOURNAMENT WINNER: ", heroes[0].getName())
}

func Duel(firstFighter, secondFighter, winner chan Hero) {
	for first := range firstFighter {
		for second := range secondFighter {
			fmt.Println("FIGHT BETWEEN ", first.getName(), "AND ", second.getName())
			for first.getHealth() >= 0 && second.getHealth() >= 0 {
				first.attack(second)
				fmt.Println("HEALTH ", second.getName(), " AFTER ATTACK ", first.getName(), " EQUAL ", second.getHealth())
				if second.getHealth() <= 0 {
					winner <- first
					fmt.Println("IN DUEL BETWEEN ", first.getName(), " AND ", second.getName(), " WON ", first.getName())
					return
				}
				second.attack(first)
				fmt.Println("HEALTH ", first.getName(), " AFTER ATTACK ", second.getName(), " EQUAL ", first.getHealth())
			}
			winner <- second
			fmt.Println("IN DUEL BETWEEN ", first.getName(), " AND ", second.getName(), " WON ", second.getName())
		}
	}
}
