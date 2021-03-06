// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"math"
	"testing"
)

// For GO386=387, make sure fucomi* opcodes are not used
// for comparison operations.
// Note that this test will fail only on a Pentium MMX
// processor (with GOARCH=386 GO386=387), as it just runs
// some code and looks for an unimplemented instruction fault.

//go:noinline
func compare1(a, b float64) bool {
	return a < b
}

//go:noinline
func compare2(a, b float32) bool {
	return a < b
}

func TestFloatCompare(t *testing.T) {
	if !compare1(3, 5) {
		t.Errorf("compare1 returned false")
	}
	if !compare2(3, 5) {
		t.Errorf("compare2 returned false")
	}

	// test folded float64 comparisons
	d1, d3, d5, d9 := float64(1), float64(3), float64(5), float64(9)
	if d3 == d5 {
		t.Errorf("d3 == d5 returned true")
	}
	if d3 != d3 {
		t.Errorf("d3 != d3 returned true")
	}
	if d3 > d5 {
		t.Errorf("d3 > d5 returned true")
	}
	if d3 >= d9 {
		t.Errorf("d3 >= d9 returned true")
	}
	if d5 < d1 {
		t.Errorf("d5 < d1 returned true")
	}
	if d9 <= d1 {
		t.Errorf("d9 <= d1 returned true")
	}
	if math.NaN() == math.NaN() {
		t.Errorf("math.NaN() == math.NaN() returned true")
	}
	if math.NaN() >= math.NaN() {
		t.Errorf("math.NaN() >= math.NaN() returned true")
	}
	if math.NaN() <= math.NaN() {
		t.Errorf("math.NaN() <= math.NaN() returned true")
	}
	if math.Copysign(math.NaN(), -1) < math.NaN() {
		t.Errorf("math.Copysign(math.NaN(), -1) < math.NaN() returned true")
	}
	if math.Inf(1) != math.Inf(1) {
		t.Errorf("math.Inf(1) != math.Inf(1) returned true")
	}
	if math.Inf(-1) != math.Inf(-1) {
		t.Errorf("math.Inf(-1) != math.Inf(-1) returned true")
	}
	if math.Copysign(0, -1) != 0 {
		t.Errorf("math.Copysign(0, -1) != 0 returned true")
	}
	if math.Copysign(0, -1) < 0 {
		t.Errorf("math.Copysign(0, -1) < 0 returned true")
	}
	if 0 > math.Copysign(0, -1) {
		t.Errorf("0 > math.Copysign(0, -1) returned true")
	}

	// test folded float32 comparisons
	s1, s3, s5, s9 := float32(1), float32(3), float32(5), float32(9)
	if s3 == s5 {
		t.Errorf("s3 == s5 returned true")
	}
	if s3 != s3 {
		t.Errorf("s3 != s3 returned true")
	}
	if s3 > s5 {
		t.Errorf("s3 > s5 returned true")
	}
	if s3 >= s9 {
		t.Errorf("s3 >= s9 returned true")
	}
	if s5 < s1 {
		t.Errorf("s5 < s1 returned true")
	}
	if s9 <= s1 {
		t.Errorf("s9 <= s1 returned true")
	}
	sPosNaN, sNegNaN := float32(math.NaN()), float32(math.Copysign(math.NaN(), -1))
	if sPosNaN == sPosNaN {
		t.Errorf("sPosNaN == sPosNaN returned true")
	}
	if sPosNaN >= sPosNaN {
		t.Errorf("sPosNaN >= sPosNaN returned true")
	}
	if sPosNaN <= sPosNaN {
		t.Errorf("sPosNaN <= sPosNaN returned true")
	}
	if sNegNaN < sPosNaN {
		t.Errorf("sNegNaN < sPosNaN returned true")
	}
	sPosInf, sNegInf := float32(math.Inf(1)), float32(math.Inf(-1))
	if sPosInf != sPosInf {
		t.Errorf("sPosInf != sPosInf returned true")
	}
	if sNegInf != sNegInf {
		t.Errorf("sNegInf != sNegInf returned true")
	}
	sNegZero := float32(math.Copysign(0, -1))
	if sNegZero != 0 {
		t.Errorf("sNegZero != 0 returned true")
	}
	if sNegZero < 0 {
		t.Errorf("sNegZero < 0 returned true")
	}
	if 0 > sNegZero {
		t.Errorf("0 > sNegZero returned true")
	}
}

// For GO386=387, make sure fucomi* opcodes are not used
// for float->int conversions.

//go:noinline
func cvt1(a float64) uint64 {
	return uint64(a)
}

//go:noinline
func cvt2(a float64) uint32 {
	return uint32(a)
}

//go:noinline
func cvt3(a float32) uint64 {
	return uint64(a)
}

//go:noinline
func cvt4(a float32) uint32 {
	return uint32(a)
}

//go:noinline
func cvt5(a float64) int64 {
	return int64(a)
}

//go:noinline
func cvt6(a float64) int32 {
	return int32(a)
}

//go:noinline
func cvt7(a float32) int64 {
	return int64(a)
}

//go:noinline
func cvt8(a float32) int32 {
	return int32(a)
}

// make sure to cover int, uint cases (issue #16738)
//go:noinline
func cvt9(a float64) int {
	return int(a)
}

//go:noinline
func cvt10(a float64) uint {
	return uint(a)
}

//go:noinline
func cvt11(a float32) int {
	return int(a)
}

//go:noinline
func cvt12(a float32) uint {
	return uint(a)
}

//go:noinline
func f2i64p(v float64) *int64 {
	return ip64(int64(v / 0.1))
}

//go:noinline
func ip64(v int64) *int64 {
	return &v
}

func TestFloatConvert(t *testing.T) {
	if got := cvt1(3.5); got != 3 {
		t.Errorf("cvt1 got %d, wanted 3", got)
	}
	if got := cvt2(3.5); got != 3 {
		t.Errorf("cvt2 got %d, wanted 3", got)
	}
	if got := cvt3(3.5); got != 3 {
		t.Errorf("cvt3 got %d, wanted 3", got)
	}
	if got := cvt4(3.5); got != 3 {
		t.Errorf("cvt4 got %d, wanted 3", got)
	}
	if got := cvt5(3.5); got != 3 {
		t.Errorf("cvt5 got %d, wanted 3", got)
	}
	if got := cvt6(3.5); got != 3 {
		t.Errorf("cvt6 got %d, wanted 3", got)
	}
	if got := cvt7(3.5); got != 3 {
		t.Errorf("cvt7 got %d, wanted 3", got)
	}
	if got := cvt8(3.5); got != 3 {
		t.Errorf("cvt8 got %d, wanted 3", got)
	}
	if got := cvt9(3.5); got != 3 {
		t.Errorf("cvt9 got %d, wanted 3", got)
	}
	if got := cvt10(3.5); got != 3 {
		t.Errorf("cvt10 got %d, wanted 3", got)
	}
	if got := cvt11(3.5); got != 3 {
		t.Errorf("cvt11 got %d, wanted 3", got)
	}
	if got := cvt12(3.5); got != 3 {
		t.Errorf("cvt12 got %d, wanted 3", got)
	}
	if got := *f2i64p(10); got != 100 {
		t.Errorf("f2i64p got %d, wanted 100", got)
	}
}

var sinkFloat float64

func BenchmarkMul2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m float64 = 1
		for j := 0; j < 500; j++ {
			m *= 2
		}
		sinkFloat = m
	}
}
func BenchmarkMulNeg2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m float64 = 1
		for j := 0; j < 500; j++ {
			m *= -2
		}
		sinkFloat = m
	}
}
