package main

import (
	"fmt"
)

type solution struct {
	single      bool
	left, right int
	kill, last  []int
}

type fragment struct {
	score    int
	solution []solution
	flag     map[int]map[int]bool // Optimization
}

func dp(monsters []int, nColor int) fragment {
	length := len(monsters)
	f := make([][]fragment, length)
	for i := 0; i < length; i++ {
		f[i] = make([]fragment, length)
		for j := 0; j < length; j++ {
			f[i][j].score = 0
			f[i][j].solution = []solution{}
			// Optimization {
			f[i][j].flag = map[int]map[int]bool{}
			for c := 0; c < nColor; c++ {
				f[i][j].flag[c] = map[int]bool{}
			}
			// }
		}
		f[i][i].solution = []solution{{true, monsters[i], monsters[i], []int{}, []int{i}}}
		f[i][i].flag[monsters[i]][monsters[i]] = true // Optimization
	}

	for i := 4; i <= length; i++ {
		for l := 0; l <= length - i; l++ {
			r := l + i - 1
			// Split
			for m := l + 3; m < r - 3; m++ {
				if f[l][m].score > 0 && f[m + 1][r].score > 0 && f[l][m].score + f[m + 1][r].score > f[l][r].score {
					tmp := fragment{f[l][m].score + f[m + 1][r].score, []solution{}, makeNewFlag(nColor)} // Optimization
					for ls := 0; ls < len(f[l][m].solution); ls++ {
						for rs := 0; rs < len(f[m + 1][r].solution); rs++ {
							if f[l][m].solution[ls].right == f[m + 1][r].solution[rs].left ||
									tmp.flag[f[l][m].solution[ls].left][f[m + 1][r].solution[rs].right] {  // Optimization
								continue
							}
							tmpKill, tmpLast := []int{}, []int{}
							tmpKill, tmpLast = append(tmpKill, f[l][m].solution[ls].kill...), append(tmpLast, f[l][m].solution[ls].last...)
							tmpKill, tmpLast = append(tmpKill, f[m + 1][r].solution[rs].kill...), append(tmpLast, f[m + 1][r].solution[rs].last...)
							tmp.solution = append(tmp.solution, solution{
								false,
								f[l][m].solution[ls].left,
								f[m + 1][r].solution[rs].right,
								tmpKill,
								tmpLast})
							tmp.flag[f[l][m].solution[ls].left][f[m + 1][r].solution[rs].right] = true // Optimization
						}
					}
					if len(tmp.solution) > 0 {
						f[l][r] = tmp
					}
				}
			}

			// Sandwitch
			for ml := l + 1; ml < r; ml++ {
				for mr := ml; mr < r; mr++ {
					if f[ml][mr].score < f[l][r].score - 1 {
						continue
					}
					tmpFragment := fragment{f[ml][mr].score + 1, []solution{}, makeNewFlag(nColor)}
					for c := 1; c < nColor; c++ {
						ms := 0
						for ; ms < len(f[ml][mr].solution) && !(f[ml][mr].solution[ms].single && f[ml][mr].solution[ms].left != c); ms++ {
						}
						if ms == len(f[ml][mr].solution) {
							continue
						}
						lc, l1, l2 := count(monsters[l:ml], c)
						l1 += l
						l2 += l
						rc, r1, r2 := count(monsters[mr + 1:r + 1], c)
						r1 += mr + 1
						r2 += mr + 1
						var first, second, third int
						if lc > 1 && rc > 0 {
							first, second, third = l1, l2, r1
						} else if lc > 0 && rc > 1 {
							first, second, third = l1, r1, r2
						} else {
							continue
						}
						tmpKill, tmpLast := []int{}, make([]int, 1)
						copy(tmpLast, f[ml][mr].solution[ms].last)
						for i := l; i < ml; i++ {
							if i != first && i != second {
								tmpKill = append(tmpKill, i)
							}
						}
						tmpKill = append(tmpKill, f[ml][mr].solution[ms].kill...)
						for i := mr + 1; i <= r; i++ {
							if i != second && i != third {
								tmpKill = append(tmpKill, i)
							}
						}
						tmpSolution := solution{true, c, c, tmpKill, tmpLast}
						if f[ml][mr].score == f[l][r].score - 1 {
							f[l][r].solution = append(f[l][r].solution, tmpSolution)
							f[l][r].flag[tmpSolution.left][tmpSolution.right] = true // Optimization
						} else {
							tmpFragment.solution = append(tmpFragment.solution, tmpSolution)
							tmpFragment.flag[tmpSolution.left][tmpSolution.right] = true // Optimization
						}
					}
					if len(tmpFragment.solution) > 0 && f[ml][mr].score >= f[l][r].score {
						f[l][r] = tmpFragment
					}
				}
			}

			//fmt.Println("************************************")
			//fmt.Println("l: ", l, ", r: ", r)
			//fmt.Println("************************************")
			//print(monsters, f, l, r)
			//fmt.Println("************************************")
		}
	}

	return f[0][length - 1]
}

func makeNewFlag(nColor int) map[int]map[int]bool {
	flag := map[int]map[int]bool{}
	for c := 0; c < nColor; c++ {
		flag[c] = map[int]bool{}
	}
	return flag
}

func count(monsters []int, color int) (count, first, second int) {
	count, first, second = 0, -1, -1
	for i, c := range monsters {
		if c == color {
			count++
			if count == 1 {
				first = i
			} else if count == 2 {
				second = i
			}
		}
	}
	return count, first, second
}

func print(monsters []int, f [][]fragment, i, j int) {
	fragment := f[i][j]
	fmt.Println("score:", fragment.score)
	for key, solution := range fragment.solution {
		fmt.Println("solution", key)
		fmt.Println("left:", solution.left)
		fmt.Println("right:", solution.right)
		fmt.Println("kill:", solution.kill)
		copyKill := []int{}
		for _, value := range solution.kill {
			copyKill = append(copyKill, value - i)
		}
		fmt.Println("result:", kill(monsters[i : j + 1], copyKill))

	}
}
