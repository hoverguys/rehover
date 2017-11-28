package main

type Color struct {
	R uint8
	G uint8
	B uint8
}

type LUT struct {
	Start        uint8
	End          uint8
	Intermediate []Color
}

// This is directly inspired by squish-gen and Simon Brown's excellent work
// on DXT1 compression
//
// Original squish-gen code (MIT license):
//    https://github.com/p3/regal/blob/master/src/squish/extra/squishgen.cpp
// Simon Brown's article on DXT1:
//    http://sjbrown.co.uk/2006/01/19/dxt-compression-techniques/

func genLookups(depth uint8) (lut []LUT) {
	maxcolors := 1 << depth
	lut = make([]LUT, maxcolors, maxcolors)

	for start := 0; start < maxcolors; start++ {

	}
	return
}
