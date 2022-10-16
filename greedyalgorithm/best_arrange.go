package greedyalgorithm

import (
	"math"
	"math/rand"
	"sort"
)

type Program struct {
	start int // 会议开始时间
	end   int // 会议结束时间
}

type Programs []Program

func (p Programs) Len() int {
	return len(p)
}

func (p Programs) Less(i, j int) bool {
	return p[i].end < p[j].end
}

func (p Programs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func BestArrange(programs []Program) int {
	sort.Sort(Programs(programs))
	curEndTime := 0
	arrangeNum := 0
	for _, p := range programs {
		if p.start >= curEndTime {
			arrangeNum++
			curEndTime = p.end
		}
	}
	return arrangeNum
}

// 穷举法，应用于对数器
func BestArrangeByExhaustive(programs []Program) int {
	return process(programs, 0, 0)
}

func process(programs []Program, done int, timeLine int) int {
	if len(programs) == 0 {
		return done
	}
	max := done
	for index, p := range programs {
		if p.start >= timeLine {
			next := copyButExcept(programs, index)
			max = int(math.Max(float64(max), float64(process(next, done+1, p.end))))
		}
	}
	return max
}

func copyButExcept(programs []Program, i int) []Program {
	var results []Program
	for index := range programs {
		if index != i {
			results = append(results, programs[index])
		}
	}
	return results
}

func generatePrograms(size int, timeMax int) []Program {
	result := make([]Program, rand.Intn(size+1))
	for index := range result {
		time1 := rand.Intn(timeMax + 1)
		time2 := rand.Intn(timeMax + 1)
		if time1 < time2 {
			result[index] = Program{
				start: time1,
				end:   time2,
			}
		} else if time2 < time1 {
			result[index] = Program{
				start: time2,
				end:   time1,
			}
		} else {
			result[index] = Program{
				start: time1,
				end:   time1 + 1,
			}
		}
	}
	return result
}
