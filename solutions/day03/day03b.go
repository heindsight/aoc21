package day03

import (
	"fmt"
	"io"
	"strconv"

	"github.com/heindsight/aoc21/registry"
)

type bitString []rune

func (bits bitString) toInt() int64 {
	bitstr := string(bits)
	val, err := strconv.ParseInt(bitstr, 2, 0)
	if err != nil {
		panic(err)
	}
	return val
}

type bitStringList []*bitString

func solveDay03b() {
	bit_strings := bitStringList{}

	for {
		var bitstring string
		_, err := fmt.Scan(&bitstring)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		bits := bitString(bitstring)
		bit_strings = append(bit_strings, &bits)
	}

	oxygen_rating, co2_rating := filter_bits(bit_strings)
	fmt.Println(oxygen_rating * co2_rating)
}

func filter_bits(bitstrings bitStringList) (int64, int64) {
	most_common, least_common := split_by_bit(bitstrings, 0)

	for pos := 1; len(most_common) > 1; pos += 1 {
		most_common, _ = split_by_bit(most_common, pos)
	}
	for pos := 1; len(least_common) > 1; pos += 1 {
		_, least_common = split_by_bit(least_common, pos)
	}

	return most_common[0].toInt(), least_common[0].toInt()
}

func split_by_bit(bitstrings bitStringList, pos int) (bitStringList, bitStringList) {
	ones := bitStringList{}
	zeros := bitStringList{}

	for _, bits := range bitstrings {
		switch (*bits)[pos] {
		case '1':
			ones = append(ones, bits)
		case '0':
			zeros = append(zeros, bits)
		}
	}

	if len(ones) >= len(zeros) {
		return ones, zeros
	} else {
		return zeros, ones
	}
}

func init() {
	if err := registry.RegisterSolution("day03b", solveDay03b); err != nil {
		panic(err)
	}
}
