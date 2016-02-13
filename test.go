package main

import (
	"fmt"
//	"strings"
//	"errors"
	"errors"
	"strings"
)

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

		//lastRight := 5
		//start := 1
		//nColor := 5
		//nMonster := 250
		//monsters := Generate(nColor, nMonster)

		fmt.Println("")
		//	fmt.Println(monsters)
		fragment := dp(monsters, 6)
		fmt.Println("score:", fragment.score)
		fmt.Println("")

		minKill, minLast := 52, 52
		minKillSolution, minLastSolution := -1, -1

		for key, solution := range fragment.solution {
			if (solution.left != lastRight && len(solution.kill) < minKill) {
				minKill = len(solution.kill)
				minKillSolution = key
			}
			if (solution.left != lastRight && len(solution.last) < minLast) {
				minLast = len(solution.last)
				minLastSolution = key
			}
		}

		if (minKillSolution != -1) {
			solution := fragment.solution[minKillSolution]
			result := kill(monsters, solution.kill)
			resultColor := []byte{}
			for key, value := range solution.kill {
				solution.kill[key] = start + value
			}
			for key, value := range solution.last {
				solution.last[key] = start + value
			}
			for _, value := range result {
				resultColor = append(resultColor, intToColor[value])
			}
			resultStr := string(resultColor)

			fmt.Println("Min kill solution:")
			//			fmt.Println("left:", solution.left)
			//			fmt.Println("right:", solution.right)
			fmt.Println("kill:", solution.kill)
			fmt.Println("result:", resultStr)
			fmt.Println("last:", solution.last)
			fmt.Println()
		}

		if (minLastSolution != -1 && minLastSolution != minKillSolution) {
			solution := fragment.solution[minLastSolution]
			result := kill(monsters, solution.kill)
			resultColor := []byte{}
			for key, value := range solution.kill {
				solution.kill[key] = start + value
			}
			for key, value := range solution.last {
				solution.last[key] = start + value
			}
			for _, value := range result {
				resultColor = append(resultColor, intToColor[value])
			}
			resultStr := string(resultColor)

			fmt.Println("Min last solution")
			//			fmt.Println("left:", solution.left)
			//			fmt.Println("right:", solution.right)
			fmt.Println("kill:", solution.kill)
			fmt.Println("result:", resultStr)
			fmt.Println("last:", solution.last)
			fmt.Println()
		}

		//for key, solution := range fragment.solution {
		//	result := kill(monsters, solution.kill)
		//	resultColor := []byte{}
		//	for key, value := range solution.kill {
		//		solution.kill[key] = start + value
		//	}
		//	for key, value := range solution.last {
		//		solution.last[key] = start + value
		//	}
		//	for _, value := range result {
		//		resultColor = append(resultColor, intToColor[value])
		//	}
		//	resultStr := string(resultColor)
		//
		//	fmt.Println("Solution", key)
		//	//fmt.Println("left:", solution.left)
		//	//fmt.Println("right:", solution.right)
		//	fmt.Println("kill:", solution.kill)
		//	fmt.Println("result:", resultStr)
		//	fmt.Println("last:", solution.last)
		//	fmt.Println()
		//}

		fmt.Println("************************************")
		fmt.Println("")
	}
}
