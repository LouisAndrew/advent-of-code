package day_nine

import (
	"regexp"
	"strconv"

	"github.com/louisandrew/advent-of-code/2024/utils"
)

const EMPTY_ID = -1

type block struct {
	id int
	r  utils.Range
}

type filedisk struct {
	f           []block
	e           []block
	initialSize int
}

// Part 1 solution. Keeping in mind that every
// block has a range of 0
// func (fd *filedisk) move() {
// 	left := 0
// 	right := len(*fd) - 1
//
// 	for left < right {
//
// 		for left < right && (*fd)[right].id == EMPTY_ID {
// 			right--
// 		}
//
// 		for left < right && (*fd)[left].id != EMPTY_ID {
// 			left++
// 		}
//
// 		if left < right && (*fd)[left].id == EMPTY_ID && (*fd)[right].id != EMPTY_ID {
// 			(*fd)[left], (*fd)[right] = (*fd)[right], (*fd)[left]
// 		}
//
// 		left++
// 		right--
// 	}
// }
// func (fd *filedisk) move() {
// 	right := len(*fd) - 1
//
// 	for right >= 0 {
// 		fmt.Println(fd.print())
//
// 		left := 0
// 		var fileId int
// 		swap := false
//
// 		for left < right && (*fd)[right].id == EMPTY_ID {
// 			right--
// 		}
//
// 		fs := utils.Range{Min: right, Max: right}
// 		fileId = (*fd)[right].id
// 		for left < right && (*fd)[right].id == fileId {
// 			fs.Min--
// 			right--
// 		}
//
// 		var es utils.Range
// 		for left < right {
// 			for left < right && (*fd)[left].id != EMPTY_ID {
// 				left++
// 			}
//
// 			es = utils.Range{Min: left, Max: left}
// 			for left < right && (*fd)[left].id == EMPTY_ID {
// 				es.Max++
// 				left++
// 			}
//
// 			if es.Length() >= fs.Length() {
// 				// swap
// 				swap = true
// 				break
// 			}
// 		}
//
// 		// if swap -> left = emptySpace.Max + 1
// 		if swap {
// 			for i := 0; i < fs.Length(); i++ {
// 				esIndex := es.Min + i
// 				fsIndex := fs.Min + i + 1
// 				(*fd)[esIndex], (*fd)[fsIndex] = (*fd)[fsIndex], (*fd)[esIndex]
// 			}
// 		}
// 	}
// }

func (fd *filedisk) swap(fIndex int, eIndex int) {
	size := (*fd).f[fIndex].r.Length()
	(*fd).f[fIndex].r.Min = (*fd).e[eIndex].r.Min
	(*fd).f[fIndex].r.Max = (*fd).e[eIndex].r.Min + size

	if (*fd).e[eIndex].r.Length() == size {
		(*fd).e = append((*fd).e[:eIndex], (*fd).e[eIndex+1:]...)
	} else {
		(*fd).e[eIndex].r.Min += size
	}
}
func (fd *filedisk) move() {
	for i := len((*fd).f) - 1; i > 0; i-- {
		size := (*fd).f[i].r.Length()
		for y, e := range (*fd).e {
			if e.r.Length() >= size && e.r.Min < (*fd).f[i].r.Min {
				fd.swap(i, y)
				break
			}
		}
	}
}
func parseLine(input string) filedisk {
	f := []block{}
	e := []block{}
	size := 0
	r, _ := regexp.Compile(`\d`)
	matches := r.FindAllString(input, -1)
	id := 0
	totalLen := 0

	for i, match := range matches {
		len, _ := strconv.Atoi(match)
		r := utils.Range{Min: totalLen, Max: totalLen + len}
		totalLen += len

		size += r.Length()
		if i%2 == 0 {
			f = append(f, block{
				id: id,
				r:  r,
			})
			id++
			continue
		}

		e = append(e, block{
			id: EMPTY_ID,
			r:  r,
		})
	}

	return filedisk{
		f: f,
		e: e,
	}
}

func Solution(input string) int {
	fd := parseLine(input)
	fd.initialSize = fd.f[len(fd.f)-1].r.Max
	fd.move()

	out := 0
	for _, b := range fd.f {
		for i := b.r.Min; i < b.r.Max; i++ {
			out += i * b.id
		}
	}

	return out
	// 8573601415879
}
