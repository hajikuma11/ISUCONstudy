package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pre, cha int

func main() {
	for p := 1; p < 100; p++ {
		fmt.Printf("%d\n", p)
	}
	n := 0
	for i := 0; i < 10; i++ { // 昇順
		pre, cha = 0, 0
		n += 10000
		data := make([]int, n)
		for j := 0; j < n; j++ {
			data[j] = j
		}
		if !sorted(merge(data)) {
			fmt.Printf("[%d]sort error!", i)
			fmt.Println(data)
			return
		}
		fmt.Printf("%d,%d,%d\n", n, pre, cha)
	}
	n = 0
	for i := 0; i < 10; i++ { // 降順
		pre, cha = 0, 0
		n += 10000
		data := make([]int, n)
		for j := 0; j < n; j++ {
			data[j] = n - j - 1
		}
		if !sorted(merge(data)) {
			fmt.Printf("[%d]sort error!", i)
			return
		}
		fmt.Printf("%d,%d,%d\n", n, pre, cha)
	}
	rand.Seed(time.Now().Unix())
	n = 0
	for i := 0; i < 10; i++ { // ランダム
		pre, cha = 0, 0
		n += 10000
		data := make([]int, n)
		for j := 0; j < n; j++ {
			data[j] = rand.Intn(10000)
		}
		if !sorted(merge(data)) {
			fmt.Printf("[%d]sort error!", i)
			return
		}
		fmt.Printf("%d,%d,%d\n", n, pre, cha)
	}
}
func merge(s []int) []int {
	var result []int
	if len(s) < 2 {
		return s
	}
	mid := int(len(s) / 2)
	r := merge(s[:mid])
	l := merge(s[mid:])
	i, j := 0, 0
	for i < len(r) && j < len(l) {
		pre++
		if r[i] > l[j] {
			cha++
			result = append(result, l[j])
			j++
		} else {
			cha++
			result = append(result, r[i])
			i++
		}
	}
	cha++
	result = append(result, r[i:]...)
	cha++
	result = append(result, l[j:]...)
	return result
}
func sorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
