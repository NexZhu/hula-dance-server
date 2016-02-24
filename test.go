package main

import (
    "fmt"
//	"strings"
//	"errors"
    "errors"
    "strings"
)

const nColor = 6

func main() {
    colorToInt := map[byte]int{'!': 0, 'r': 1, 'y': 2, 'g': 3, 'b': 4, 'p': 5}
    intToColor := map[int]byte{}
    for key, value := range colorToInt {
        intToColor[value] = key
    }

    //	monsters := []int{2, 2, 1, 2}
    //	monsters := []int{2, 2, 1, 3, 2}
    //	monsters := []int{2, 2, 1, 2, 3}
    //	monsters := []int{3, 2, 2, 1, 2, 3, 3}
    //	monsters := []int{3, 4, 2, 2, 1, 2, 3, 3}
    //	monsters := []int{3, 4, 2, 2, 1, 2, 3, 3, 2, 2, 3, 2}
    //	monsters := []int{1, 2, 2, 4, 1, 3, 0, 0, 1, 0, 4, 1, 2, 4, 3, 4, 1, 0, 2, 1, 0, 1, 3, 3, 2}
    // Real data 1
    //	monsters := []int{1, 1, 2, 3, 1, 1, 2, 3, 3, 1, 1, 3, 1, 3, 4, 4, 2, 4, 3, 3, 3, 1, 4, 1, 3, 2, 4, 2, 4, 2, 1, 3, 2, 1, 5, 4, 2, 1, 5, 4, 4, 1, 4, 2, 1, 2, 2, 4, 2, 2, 1, 4, 5, 2, 2, 1, 5, 4, 4, 5}
    // Real data 2
    //	monsters := []int{2, 2, 5, 1, 2, 5, 4, 1, 4, 5, 1, 5, 5, 5, 4, 4, 1, 6, 1, 1, 2, 1, 2, 2, 2, 4, 1, 1, 2, 4, 2, 2, 6, 1, 4, 1, 3, 1, 5, 3, 4, 5, 6, 3, 4, 3, 3, 1, 4, 3, 5, 3, 5, 3, 3, 4, 4, 6, 5, 1, 5, 5, 4, 4, 4, 4, 3, 1, 3, 3, 5, 4, 3, 5, 5, 2, 5, 4, 3, 4, 2, 2, 5, 2, 3, 3, 3, 3, 3, 2, 5, 6, 5, 3, 4, 5, 5, 2, 2, 6, 2, 5, 3, 3, 2, 5, 5, 2, 3, 5, 2, 2, 2, 3, 2, 6, 2, 5, 2, 3, 1, 5, 3, 3, 6, 5, 5, 3, 3, 2, 1, 3, 2, 5, 1, 5, 1, 1, 4, 4, 1, 1, 1, 4}

    for {
        fmt.Print("Last right color: ")
        var lastRightStr string
        for err := errors.New(""); err != nil; {
            _, err = fmt.Scanln(&lastRightStr)
        }
        lastRight := colorToInt[lastRightStr[0]]

        fmt.Print("Start index: ")
        var start int
        for err := errors.New(""); err != nil; {
            _, err = fmt.Scanln(&start)
        }

        fmt.Println("Monsters: ")
        var monstersStr string
        for err := errors.New(""); err != nil; {
            _, err = fmt.Scanln(&monstersStr)
        }

        monstersColor := strings.Split(monstersStr, "")
        monsters := []int{}
        for _, value := range monstersColor {
            monsters = append(monsters, colorToInt[value[0]])
        }

        //lastRight := 0
        //start := 1
        //nMonster := 250
        //monsters := Generate(nColor, nMonster)

        fmt.Println()
        //	fmt.Println(monsters)
        solutions := solveHulaDance(monsters, nColor)

        max := -1
        maxSolutions := []solution{}
        //solution := solution{}
        for lc := 1; lc < nColor; lc++ {
            if (lc == lastRight) {
                continue
            }

            for rc := 1; rc < nColor; rc++ {
                if solutions[lc][rc].score > max {
                    maxSolutions = []solution{}
                    max = solutions[lc][rc].score
                }
                if solutions[lc][rc].score >= max {
                    maxSolutions = append(maxSolutions, solutions[lc][rc])
                }
            }
        }

        if len(maxSolutions) > 0 && maxSolutions[0].score == 0 {
            fmt.Println("no solution")
        } else {
            for k, s := range maxSolutions {
                fmt.Println("solution", k + 1)
                result := kill(monsters, s.kill)
                resultColor := []byte{}
                for key, value := range s.kill {
                    s.kill[key] = start + value
                }
                for key, value := range s.last {
                    s.last[key] = start + value
                }
                for _, value := range result {
                    resultColor = append(resultColor, intToColor[value])
                }
                resultStr := string(resultColor)

                fmt.Println("score:", max)
                //fmt.Println("left:", s.left)
                //fmt.Println("right:", s.right)
                fmt.Println("kill:", s.kill)
                fmt.Println("result:", resultStr)
                fmt.Println("last:", s.last)
                fmt.Println("")
            }
        }

        fmt.Println("************************************")
        fmt.Println("")
    }
}
