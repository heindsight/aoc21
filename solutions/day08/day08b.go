package day08

import (
	"fmt"
	"sort"
	"unicode/utf8"

	"github.com/heindsight/aoc21/registry"
	"github.com/heindsight/aoc21/utils/set"
)

func solveDay08b() {
	sum := 0
	for display := range readDisplaySignals() {
		signal_map := decodeSignals(display)
		sum += decodeMessage(display, signal_map)
	}
	fmt.Println(sum)
}

type signalInfo struct {
	signal   string
	segments set.Set
}

func makeSignalInfo(signal string) signalInfo {
	segments := set.NewSet()
	for _, segment := range signal {
		segments.Add(segment)
	}
	return signalInfo{signal: sortString(signal), segments: segments}
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func decodeSignals(disp displaySignal) map[string]int {
	digitSignals := make(map[int]signalInfo)
	fiveSegment := make([]signalInfo, 0, 3)
	sixSegment := make([]signalInfo, 0, 3)

	for _, signal := range disp.uniques {
		digit, unique := uniqueDigits[utf8.RuneCountInString(signal)]
		if unique {
			digitSignals[digit] = makeSignalInfo(signal)
		} else if len(signal) == 5 {
			fiveSegment = append(fiveSegment, makeSignalInfo(signal))
		} else if len(signal) == 6 {
			sixSegment = append(sixSegment, makeSignalInfo(signal))
		}
	}

	for _, signal := range sixSegment {
		if digitSignals[4].segments.IsSubset(signal.segments) {
			digitSignals[9] = signal
		} else if digitSignals[1].segments.IsSubset(signal.segments) {
			digitSignals[0] = signal
		} else {
			digitSignals[6] = signal
		}
	}

	for _, signal := range fiveSegment {
		if digitSignals[1].segments.IsSubset(signal.segments) {
			digitSignals[3] = signal
		} else if signal.segments.IsSubset(digitSignals[6].segments) {
			digitSignals[5] = signal
		} else {
			digitSignals[2] = signal
		}
	}

	decodeMap := make(map[string]int)
	for digit, signal := range digitSignals {
		decodeMap[signal.signal] = digit
	}
	return decodeMap
}

func decodeMessage(disp displaySignal, signal_map map[string]int) int {
	value := 0
	for _, digit := range disp.output {
		value *= 10
		value += signal_map[sortString(digit)]
	}
	return value
}

func init() {
	if err := registry.RegisterSolution("day08b", solveDay08b); err != nil {
		panic(err)
	}
}
