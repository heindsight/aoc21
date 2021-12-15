package day13

import (
	"fmt"
	"io"

	"github.com/heindsight/aoc21/utils/grid"
	"github.com/heindsight/aoc21/utils/set"
)

type foldInfo struct {
	direction rune
	position int
}

func readDots() set.Set {
	dots := set.NewSet()
	for {
		var dot grid.Point
		scanned, _ := fmt.Scanf("%d,%d", &dot.X, &dot.Y)
		if scanned < 2 {
			break
		}

		dots.Add(dot)
	}
	return dots
}

func readFolds() chan foldInfo {
	out := make(chan foldInfo)
	go func() {
		for {
			var (
				fold foldInfo
				f string
				a string
			)
			_, err := fmt.Scanf("%s %s %c=%d", &f, &a, &fold.direction, &fold.position)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			out <- fold
		}
		close(out)
	}()
	return out
}

func doFold(page set.Set, how foldInfo) {
	for dot := range page.Iter() {
		coord := dot.(grid.Point)
		if how.direction == 'x' {
			if coord.X > how.position {
				page.Delete(coord)
				coord.X = calc_fold(coord.X, how.position)
				page.Add(coord)
			}
		} else if how.direction == 'y' {
			if coord.Y > how.position {
				page.Delete(coord)
				coord.Y = calc_fold(coord.Y, how.position)
				page.Add(coord)
			}
		}
	}
}

func calc_fold(coord int, fold_at int) int {
	return 2*fold_at - coord
}
