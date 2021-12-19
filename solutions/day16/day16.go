package day16

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/heindsight/aoc21/registry"
)

type parsedPacket struct {
	Version  int64
	Type     int64
	Value    int64
	Children []parsedPacket
}

type Day16 struct {
	Evaluate func(parsedPacket) int64
}

func (d *Day16) solve() {
	var hexPacket string
	fmt.Scanln(&hexPacket)

	binPacket := hexToBin(hexPacket)

	parsed, _ := parse(binPacket)
	value := d.Evaluate(parsed)
	fmt.Println(value)
}

func sumVersions(packet parsedPacket) int64 {
	sum := int64(packet.Version)
	for _, child := range packet.Children {
		sum += sumVersions(child)
	}
	return sum
}

func evaluate(packet parsedPacket) int64 {
	var value int64
	switch packet.Type {
	case 0:
		for _, child := range packet.Children {
			value += evaluate(child)
		}
	case 1:
		value = 1
		for _, child := range packet.Children {
			value *= evaluate(child)
		}
	case 2:
		value = math.MaxInt64
		for _, child := range packet.Children {
			val := evaluate(child)
			if val < value {
				value = val
			}
		}
	case 3:
		value = 0
		for _, child := range packet.Children {
			val := evaluate(child)
			if val > value {
				value = val
			}
		}
	case 4:
		value = packet.Value
	case 5:
		if evaluate(packet.Children[0]) > evaluate(packet.Children[1]) {
			value = 1
		}
	case 6:
		if evaluate(packet.Children[0]) < evaluate(packet.Children[1]) {
			value = 1
		}
	case 7:
		if evaluate(packet.Children[0]) == evaluate(packet.Children[1]) {
			value = 1
		}
	}
	return value
}

func hexToBin(h string) string {
	if len(h)%2 == 1 {
		h = h + "0"
	}
	decoded, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	var builder strings.Builder
	for _, b := range decoded {
		fmt.Fprintf(&builder, "%08b", b)
	}
	return builder.String()
}

func parse(bits string) (parsed parsedPacket, length int) {
	parsed.Version , _ = strconv.ParseInt(bits[0:3], 2, 8)
	parsed.Type , _ = strconv.ParseInt(bits[3:6], 2, 8)

	if parsed.Type == 4 {
		parsed.Value, length = parseLiteral(bits)
	} else {
		parsed.Children, length = parseOperator(bits)
	}
	return parsed, length
}

func parseLiteral(bits string) (int64, int) {
	var builder strings.Builder
	offset := 6

	for {
		offset += 5
		fmt.Fprintf(&builder, "%s", bits[offset-4:offset])
		if bits[offset-5] == '0' {
			break
		}
	}

	value, _ := strconv.ParseInt(builder.String(), 2, 64)
	return value, offset
}

func parseOperator(bits string) ([]parsedPacket, int) {
	var children []parsedPacket
	var offset int

	length_type := bits[6]

	if length_type == '0' {
		operand_bits, _ := strconv.ParseInt(bits[7:22], 2, 16)
		children = make([]parsedPacket, 0, 2)

		for offset = 22; offset < 22+int(operand_bits); {
			child, length := parse(bits[offset:])
			children = append(children, child)
			offset += length
		}
	} else {
		num_children, _ := strconv.ParseInt(bits[7:18], 2, 32)
		children = make([]parsedPacket, num_children)
		offset = 18

		for i := 0; i < int(num_children); i++ {
			child, length := parse(bits[offset:])
			children[i] = child
			offset += length
		}
	}

	return children, offset
}

func init() {
	day16a := Day16{Evaluate: sumVersions}
	if err := registry.RegisterSolution("day16a", day16a.solve); err != nil {
		panic(err)
	}
	day16b := Day16{Evaluate: evaluate}
	if err := registry.RegisterSolution("day16b", day16b.solve); err != nil {
		panic(err)
	}
}
