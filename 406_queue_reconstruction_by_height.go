/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 * 7月 13 2020
 */
package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return false
		} else if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return true
		}
	})

	people = sortKValue(people)

	return people
}

func sortKValue(people [][]int) [][]int {
	var peopleNew [][]int

	for _, item := range people {
		height := item[0]
		number := item[1]
		if len(peopleNew) < number {
			peopleNew = append(peopleNew, []int{height, number})
		} else {
			temp := append([][]int{}, peopleNew[number:]...)
			peopleNew = append(people[0:number], []int{height, number})
			peopleNew = append(peopleNew, temp...)
		}
	}
	return peopleNew
}
