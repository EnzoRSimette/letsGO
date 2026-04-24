package main

import (
	"fmt"
	"slices"
)

func analyze(nums []int) string {
	count := len(nums)
	if count == 0 {
		return "you didnt type any number"
	}
	var sum int
	for _, v := range nums {
		sum += v
	}
	avg := float64(sum) / float64(count)
	max := slices.Max(nums)
	min := slices.Min(nums)
	return fmt.Sprintf("count=%d sum=%d min=%d max=%d avg=%.2f", count, sum, min, max, avg)
}
