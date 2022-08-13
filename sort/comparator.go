package sort

import "sort"

type Student struct {
	ID   int
	Age  int
	Name string
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age // 按年龄正序
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func CompareStudent(studentSlice []Student) []Student {
	sort.Sort(StudentSlice(studentSlice))
	return studentSlice
}
