package main

import (
	"fmt"
	"sort"
)

type CriticalPoints [][]int

func (cp CriticalPoints) Len() int {
	return len(cp)
}

func (cp CriticalPoints) Less(i, j int) bool {
	return cp[i][0] < cp[j][0]
}

func (cp CriticalPoints) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
}

func getSkyline(buildings [][]int) [][]int {
	ret := getAllCriticalPoints(buildings)
	for i := 0; i < len(buildings); i++ {
		for j := 0; j < len(ret); j++ {
			if ret[j][0] >= buildings[i][0] && ret[j][0] < buildings[i][1] {
				ret[j][1] = max(buildings[i][2], ret[j][1])
			}
		}
	}
	for i := len(ret) - 1; i >= 1; i-- {
		if ret[i][1] == ret[i-1][1] {
			ret = append(ret[:i], ret[i+1:]...)
		}
	}
	return ret
}

func getSkyline2(buildings [][]int) [][]int {
	cps := getAllCriticalPoints(buildings)
	for _, c := range cps {
		for _, b := range buildings {
			if b[2] < c[1] {
				continue
			}
			if c[0] >= b[0] && c[0] < b[1] {
				c[1] = max(b[2], c[1])
			}
		}
	}
	for i := len(cps) - 1; i >= 1; i-- {
		if cps[i][1] == cps[i-1][1] {
			cps = append(cps[:i], cps[i+1:]...)
		}
	}
	return cps
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getAllCriticalPoints(buildings [][]int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < len(buildings); i++ {
		ret = append(ret, []int{buildings[i][0], buildings[i][2]})
		// Right edge of a building cannot be treated as part of the building
		ret = append(ret, []int{buildings[i][1], 0})
	}
	sort.Sort(CriticalPoints(ret))
	return ret
}

func main() {
	buildings := [][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	}
	fmt.Printf("%v\n", getSkyline(buildings))
	fmt.Printf("%v\n", getSkyline2(buildings))
}
