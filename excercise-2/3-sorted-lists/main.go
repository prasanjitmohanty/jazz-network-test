package main

import "fmt"

func main() {
	fmt.Println(mergeSortedLists([]int{1, 4, 6}, []int{2, 3, 5}))
}

func mergeSortedLists(l1 []int, l2 []int) []int {
	mlist := make([]int, len(l1)+len(l2))
	i, j, k := 0, 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			mlist[k] = l1[i]
			i++
			k++
		} else {
			mlist[k] = l2[j]
			j++
			k++
		}
	}

	for i < len(l1) {
		mlist[k] = l1[i]
		i++
		k++
	}

	for j < len(l2) {
		mlist[k] = l2[j]
		j++
		k++
	}

	return mlist
}
