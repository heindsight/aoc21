package day21

import (
	"strconv"
	"strings"
)

func readStartPosition(reader chan string) int {
	line := <- reader
	parts := strings.Split(line, ": ")
	position, _ := strconv.Atoi(parts[1])
	return position
}
