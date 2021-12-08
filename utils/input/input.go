package input

import (
	"fmt"
	"strconv"
	"strings"
)

type InputItem struct {
	Index int
	Value string
}

type InputItemInt struct {
	Index int
	Value int
}

func ReadCommaSepLine() chan InputItem {
	out := make(chan InputItem)
	go func() {
		var line string
		_, err := fmt.Scanf("%s\n", &line)
		if err != nil {
			panic(err)
		}

		for index, val := range strings.Split(line, ",") {
			out <- InputItem{Index: index, Value: val}
		}
		close(out)
	} ()
	return out
}

func ReadCommaSepLineInts() chan InputItemInt {
	out := make(chan InputItemInt)
	go func() {
		for item := range ReadCommaSepLine() {
			value, err := strconv.Atoi(item.Value)
			if err != nil {
				panic(err)
			}
			out <- InputItemInt{Index: item.Index, Value: value}
		}
		close(out)
	} ()
	return out
}
