package day22

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/heindsight/aoc21/utils/numeric"
)

type Interval struct {
	Upper int
	Lower int
}

func (iv *Interval) Intersection(other *Interval) (Interval, bool) {
	intersects := iv.Lower < other.Upper && other.Lower < iv.Upper

	return Interval{
		Lower: numeric.Max(iv.Lower, other.Lower),
		Upper: numeric.Min(iv.Upper, other.Upper),
	}, intersects
}

func (iv *Interval) Length() uint64 {
	return uint64(iv.Upper - iv.Lower)
}

type Cuboid struct {
	X, Y, Z Interval
}

func (cu *Cuboid) Volume() uint64 {
	return cu.X.Length() * cu.Y.Length() * cu.Z.Length()
}

func (cu *Cuboid) Intersection(other *Cuboid) (Cuboid, bool) {
	x, x_intersects := cu.X.Intersection(&other.X)
	y, y_intersects := cu.Y.Intersection(&other.Y)
	z, z_intersects := cu.Z.Intersection(&other.Z)

	return Cuboid{X: x, Y: y, Z: z}, x_intersects && y_intersects && z_intersects
}

type BootStep struct {
	On bool
	Region Cuboid
}

func parseBootStep(instruction string) BootStep {
	bs := BootStep{}

	fields := strings.Fields(instruction)
	bs.On = fields[0] == "on"

	for _, interval := range strings.Split(fields[1], ",") {
		axis, ival := parseInterval(interval)
		switch axis {
		case 'x':
			bs.Region.X = ival
		case 'y':
			bs.Region.Y = ival
		case 'z':
			bs.Region.Z = ival
		}
	}
	return bs
}

func parseInterval(desc string) (rune, Interval) {
	parts := strings.Split(desc, "=")
	axis, _ := utf8.DecodeRuneInString(parts[0])
	endpoints := strings.Split(parts[1], "..")
	ival := Interval{}
	ival.Lower, _ = strconv.Atoi(endpoints[0])
	ival.Upper, _ = strconv.Atoi(endpoints[1])
	ival.Upper++
	return axis, ival
}
