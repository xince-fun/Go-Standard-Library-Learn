package main

import (
	"sort"
)

type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = StuScores(nil)

func main() {
	/*
		stus := StuScores{
			{"alan", 95},
			{"hikerell", 91},
			{"acmfly", 96},
			{"leao", 90},
		}

		fmt.Println("Default:\n\t", stus)
		sort.Sort(sort.Reverse(stus))
		fmt.Println("Is Sorted?\n\t", sort.IsSorted(stus))
		fmt.Println("Sorted:\n\t", stus)
	*/

	/*
		x := 11
		s := []int{3, 6, 8, 11, 45} // 已经升序排序
		pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
		if pos < len(s) && s[pos] == x {
			fmt.Println(x, " 在 s 中的位置为：", pos)
		} else {
			fmt.Println("s 不包含元素 ", x)
		}
	*/

	/*
		s := []int{5, 2, 6, 3, 1, 4}
		sort.Ints(s)
		fmt.Println(s)
	*/

	/*
		s := []int{5, 2, 6, 3, 1, 4}
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
		fmt.Println(s)
	*/

	/*
		people := []struct {
			Name string
			Age  int
		}{
			{"Gopher", 7},
			{"Alice", 55},
			{"Vera", 24},
			{"Bob", 75},
		}

		sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 按年龄升序排序
		fmt.Println("Sort by age:", people)
	*/
}
