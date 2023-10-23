package main

import (
	"fmt"
	"sort"
)

type Interview struct {
	StartTime int
	EndTime   int
}

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)

	interviews := make([]Interview, n)
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		interviews[i] = Interview{StartTime: start, EndTime: end}
	}
	sort.Slice(interviews, func(i, j int) bool {
		return interviews[i].StartTime < interviews[j].StartTime
	})

	needInterviewers := 1
	endTimes := make([]int, m)
	endTimes[0] = interviews[0].EndTime

	for i := 1; i < n; i++ {
		startTime := interviews[i].StartTime
		scheduled := false
		for j := 0; j < m; j++ {
			if endTimes[j] <= startTime {
				endTimes[j] = interviews[i].EndTime
				scheduled = true
				break
			}
		}
		if !scheduled {
			needInterviewers++
			endTimes[0] = interviews[i].EndTime
		}
		sort.Ints(endTimes)
	}
	fmt.Println(needInterviewers)
}
