package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero interface {
	attack(Hero)
	defense() bool
	criticalHit() bool
	setId(uint64)
	setName(string)
	setHealth(float64)
	setDamage(float64)
	getId() uint64
	getName() string
	getHealth() float64
	getDamage() float64
}

type Warrior struct {
	id     uint64
	name   string
	health float64
	damage float64
	shield float64
}

type Wizard struct {
	id     uint64
	name   string
	health float64
	damage float64
	shield float64
}

type Archer struct {
	id     uint64
	name   string
	health float64
	damage float64
	shield float64
}

func (warrior *Warrior) attack(hero Hero) {
	if hero.defense() {
		fmt.Println("Block")
	} else if !hero.criticalHit() {
		fmt.Println("Damage: ", warrior.damage)
		hero.setHealth(hero.getHealth() - warrior.damage)
	} else {
		fmt.Println("Damage: ", warrior.damage * 2)
		hero.setHealth(hero.getHealth() - warrior.damage * 2)
	}
}

func (warrior *Warrior) defense() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 3 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}

func (warrior *Warrior) criticalHit() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 2 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}


func (wizard *Wizard) attack(hero Hero) {
	if hero.defense() {
		fmt.Println("Block")
	} else if !hero.criticalHit() {
		fmt.Println("Damage: ", wizard.damage)
		hero.setHealth(hero.getHealth() - wizard.damage)
	} else {
		fmt.Println("Damage: ", wizard.damage * 2)
		hero.setHealth(hero.getHealth() - wizard.damage * 2)
	}
}

func (wizard *Wizard) defense() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 1 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}

func(wizard *Wizard) criticalHit() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 7 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}

func (archer *Archer) attack(hero Hero) {
	if hero.defense() {
		fmt.Println("Block")
	} else if !hero.criticalHit() {
		fmt.Println("Damage: ", archer.damage)
		hero.setHealth(hero.getHealth() - archer.damage)
	} else {
		fmt.Println("Damage: ", archer.damage * 2)
		hero.setHealth(hero.getHealth() - archer.damage * 2)
	}
}

func (archer *Archer) defense() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 2 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}

func (archer *Archer) criticalHit() bool {
	var array [10]bool
	number := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		if i < 4 {
			array[i] = true
		} else {
			array[i] = false
		}
	}
	index := number.Intn(9)
	return array[index]
}

func (warrior *Warrior) setId(id uint64) {
	warrior.id = id
}

func (warrior *Warrior) getId() uint64 {
	return warrior.id
}

func (wizard *Wizard) setId(id uint64) {
	wizard.id = id
}

func (wizard *Wizard) getId() uint64 {
	return wizard.id
}

func (archer *Archer) setId(id uint64) {
	archer.id = id
}

func (archer *Archer) getId() uint64 {
	return archer.id
}

func (warrior *Warrior) setName(name string) {
	warrior.name = name
}

func (warrior *Warrior) getName() string {
	return warrior.name
}

func (wizard *Wizard) setName(name string) {
	wizard.name = name
}

func (wizard *Wizard) getName() string {
	return wizard.name
}

func (archer *Archer) setName(name string) {
	archer.name = name
}

func (archer *Archer) getName() string {
	return archer.name
}

func (warrior *Warrior) setHealth(health float64) {
	warrior.health = health
}

func (warrior *Warrior) getHealth() float64 {
	return warrior.health
}

func (wizard *Wizard) setHealth(health float64) {
	wizard.health = health
}

func (wizard *Wizard) getHealth() float64 {
	return wizard.health
}

func (archer *Archer) setHealth(health float64) {
	archer.health = health
}

func (archer *Archer) getHealth() float64 {
	return archer.health
}

func (warrior *Warrior) setDamage(damage float64) {
	warrior.damage = damage
}

func (warrior *Warrior) getDamage() float64 {
	return warrior.damage
}

func (wizard *Wizard) setDamage(damage float64) {
	wizard.damage = damage
}

func (wizard *Wizard) getDamage() float64 {
	return wizard.damage
}

func (archer *Archer) setDamage(damage float64) {
	archer.damage = damage
}

func (archer *Archer) getDamage() float64 {
	return archer.damage
}

