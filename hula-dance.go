package main

type solution struct {
	score       int
	single      bool
	left, right int
	kill, last  []int
}

type palindrome struct {
	flag   bool
	length int
	kill   []int
	mid    int
}

func solveHulaDance(monsters []int, nColor int) [][]solution {
	p := findPalindromes(monsters, nColor)

	length := len(monsters)
	s := make([][][][]solution, length)
	for i := 0; i < length; i++ {
		s[i] = make([][][]solution, length)
		for j := 0; j < length; j++ {
			s[i][j] = make([][]solution, nColor)
			for c := 1; c < nColor; c++ {
				s[i][j][c] = make([]solution, nColor)
				if p[i][j][c].flag {
					s[i][j][c][c] = solution{p[i][j][c].length - 1, true, c, c, p[i][j][c].kill, []int{p[i][j][c].mid}}
				}
			}
		}
	}

	for i := 4; i <= length; i++ {
		for l := 0; l <= length-i; l++ {
			r := l + i - 1
			for m := l + 3; m < r-3; m++ {
				for llc := 1; llc < nColor; llc++ {
					for lrc := 1; lrc < nColor; lrc++ {
						for rc := 1; rc < nColor; rc++ {
							if rc == lrc {
								continue
							}
							curr := &s[l][r][llc][rc]
							ls := s[l][m][llc][lrc]
							rs := s[m+1][r][rc][rc]
							if !(ls.score > 0 && rs.single && rs.score > 0 && ls.score+rs.score > curr.score) {
								continue
							}

							newKill, newLast := []int{}, []int{}
							newKill, newLast = append(newKill, ls.kill...), append(newLast, ls.last...)
							newKill, newLast = append(newKill, rs.kill...), append(newLast, rs.last...)
							*curr = solution{ls.score + rs.score, false, llc, rc, newKill, newLast}
						}
					}
				}
			}
		}
	}

	return s[0][length-1]
}

func findPalindromes(monsters []int, nColor int) [][][]palindrome {
	length := len(monsters)
	p := make([][][]palindrome, length)
	for i := 0; i < length; i++ {
		p[i] = make([][]palindrome, length)
		for j := 0; j < length; j++ {
			p[i][j] = make([]palindrome, nColor)
		}
		p[i][i][monsters[i]] = palindrome{true, 1, []int{}, i}
	}
	for i, count := 1, 1; i <= length; i++ {
		if i < length && monsters[i] == monsters[i-1] {
			count++
		} else {
			if count >= 4 {
				p[i-count][i-1][monsters[i-1]] = palindrome{true, 2, []int{}, i - 1 - count/2}
			}
			count = 1
		}
	}

	for i := 4; i <= length; i++ {
		for l := 0; l <= length-i; l++ {
			r := l + i - 1
			for ml := l + 1; ml < r; ml++ {
				for mr := ml; mr < r; mr++ {
					midPalindromes := p[ml][mr]

					for c := 1; c < nColor; c++ {
						curr := &p[l][r][c]

						midPalindrome := &midPalindromes[c]
						if midPalindrome.flag && (!curr.flag || curr.length < midPalindrome.length) {
							newKill := appendKill(l, r, ml, mr, -1, -1, -1, midPalindrome.kill)
							*curr = palindrome{true, midPalindrome.length, newKill, midPalindrome.mid}
						}

						mcMax := -1
						max := 0
						for mc := 0; mc < nColor; mc++ {
							if mc != c && midPalindromes[mc].flag && midPalindromes[mc].length > max {
								mcMax = mc
								max = midPalindromes[mc].length
							}
						}
						if mcMax == -1 || (curr.flag && max < curr.length) {
							continue
						}

						midPalindrome = &midPalindromes[mcMax]
						enough, c1, c2, c3 := countWrap(monsters, c, l, ml-1, mr+1, r)
						if !enough {
							continue
						}
						newKill := appendKill(l, r, ml, mr, c1, c2, c3, midPalindrome.kill)
						*curr = palindrome{true, max + 1, newKill, midPalindrome.mid}
					}
				}
			}
		}
	}

	return p
}

func countWrap(monsters []int, c, ll, lr, rl, rr int) (bool, int, int, int) {
	lc, l1, l2 := count(monsters[ll:lr+1], c)
	l1 += ll
	l2 += ll
	rc, r1, r2 := count(monsters[rl:rr+1], c)
	r1 += rl
	r2 += rl
	if lc > 1 && rc > 0 {
		return true, l1, l2, r1
	} else if lc > 0 && rc > 1 {
		return true, l1, r1, r2
	} else {
		return false, -1, -1, -1
	}
}

func count(monsters []int, color int) (count, c1, c2 int) {
	count, c1, c2 = 0, -1, -1
	for i, c := range monsters {
		if c == color {
			count++
			if count == 1 {
				c1 = i
			} else if count == 2 {
				c2 = i
			}
		}
	}
	return count, c1, c2
}

func appendKill(l, r, ml, mr, c1, c2, c3 int, midKill []int) []int {
	newKill := []int{}
	for i := l; i < ml; i++ {
		if i != c1 && i != c2 {
			newKill = append(newKill, i)
		}
	}
	newKill = append(newKill, midKill...)
	for i := mr + 1; i <= r; i++ {
		if i != c2 && i != c3 {
			newKill = append(newKill, i)
		}
	}
	return newKill
}
