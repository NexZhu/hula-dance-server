package main

import (
    "math/rand"
)

func Generate(nColor, nMonster int) []int {
    monsters := []int{}
    for i := 0; i < nMonster; i++ {
        monsters = append(monsters, rand.Intn(nColor))
    }
    return monsters
}
