package todoliste

import "sort"

type Task struct {
	Id          int
	Name        string
	Description string
	State       string
	Deadline    string
}

func GetLastId(tasks []Task) int {
	if len(tasks) == 0 {
		return 0
	}
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Id < tasks[j].Id
	})
	return tasks[len(tasks)-1].Id
}
