// Package lsproduct calculates the largest product in a series
package lsproduct

import (
	"fmt"
	"strconv"
)

var testVersion = 4

func parse(s string) ([]int, error) {
	numbers := make([]int, 0)
	for x := range s {
		n, err := strconv.Atoi(s[x : x+1])
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func window(series []int, width int) <-chan []int {
	ch := make(chan []int)
	go func() {
		mx := len(series) - width
		for x := 0; x <= mx; x++ {
			ch <- series[x : x+width]
		}
		close(ch)
	}()
	return ch
}

func product(series []int) int {
	total := 1
	for _, x := range series {
		total *= x
	}
	return total
}

// LargestSeriesProduct returns the product of a subseries of s with the largest value
func LargestSeriesProduct(s string, width int) (int, error) {
	if width < 0 || width > len(s) {
		return 0, fmt.Errorf("Width must be 0 <= width <= length of string")
	}

	series, err := parse(s)
	if err != nil {
		return 0, err
	}

	result := 0
	for w := range window(series, width) {
		p := product(w)

		if p > result {
			result = p
		}
	}
	return result, nil
}
