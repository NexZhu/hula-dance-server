package main

func kill(monsters []int, kills []int) []int {
	set := make(map[int]bool)
	for _, v := range kills {
		set[v] = true
	}
	result := []int{}
	for key, value := range monsters {
		if !set[key] {
			result = append(result, value)
		}
	}
	return result
}
