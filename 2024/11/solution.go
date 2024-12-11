package day_eleven

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

type arrangement []int

func parse(input string) arrangement {
	r := regexp.MustCompile(`\d+`)
	arr := []int{}
	matches := r.FindAllString(input, -1)
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		arr = append(arr, num)
	}

	return arr
}

var mu sync.Mutex
var m = make(map[string]int)

func write(j, blinks int, value int) {
	mu.Lock()
	defer mu.Unlock()
	key := fmt.Sprintf("%d-%d", j, blinks)

	m[key] = value
}

func read(j, blinks int) int {
	mu.Lock()
	defer mu.Unlock()
	key := fmt.Sprintf("%d-%d", j, blinks)
	if _, o := m[key]; o {
		return m[key]
	}

	return 0
}

func blink(i int) arrangement {
	digits := strconv.Itoa(i)

	if i == 0 {
		return []int{1}
	}

	if len(digits)%2 == 0 {
		h := len(digits) / 2
		a, b := digits[:h], digits[h:]
		firstHalf, _ := strconv.Atoi(a)
		secondHalf, _ := strconv.Atoi(b)

		return []int{firstHalf, secondHalf}
	}

	return []int{i * 2024}
}
func blinkStones(stones []int) arrangement {
	out := []int{}
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, i := range stones {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer m.Unlock()
			m.Lock()
			out = append(out, blink(i)...)
		}(i)
	}

	wg.Wait()
	return out
}

func getCountAfterBlinks(num, blinks int) int {
	count := 0
	stones := []int{num}
	if blinks == 0 {
		return 1
	}

	if val := read(num, blinks); val != 0 {
		return val
	}

	stones = blinkStones(stones)
	for _, stone := range stones {
		count += getCountAfterBlinks(stone, blinks-1)
	}

	write(num, blinks, count)
	return count
}

func Solution(input string) int {
	arr := parse(input)
	out := 0
	blinks := 75
	for _, i := range arr {
		out += getCountAfterBlinks(i, blinks)
	}

	return out
}
