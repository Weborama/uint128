// Copyright 2018 Weborama. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package uint128 // import "github.com/weborama/uint128"

import "math/bits"

// Cmp compares two Uint128 and returns one of the following values:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
//
// nolint: varnamelen
func Cmp(x, y Uint128) int {
	if x.H < y.H {
		return -1
	} else if x.H > y.H {
		return 1
	}

	if x.L < y.L {
		return -1
	} else if x.L > y.L {
		return 1
	}

	return 0
}

// IsZero returns true if x is zero.
func IsZero(x Uint128) bool {
	return x.H == 0 && x.L == 0
}

// ShiftLeft shifts x to the left by the provided number of bits.
// nolint: varnamelen
func ShiftLeft(x Uint128, bits uint) Uint128 {
	switch {
	case bits >= numBits:
		x.H = 0
		x.L = 0
	case bits >= numHalfBits:
		x.H = x.L << (bits - numHalfBits)
		x.L = 0
	default:
		x.H <<= bits
		x.H |= x.L >> (numHalfBits - bits)
		x.L <<= bits
	}

	return x
}

// ShiftRight shifts x to the right by the provided number of bits.
// nolint: varnamelen
func ShiftRight(x Uint128, bits uint) Uint128 {
	switch {
	case bits >= numBits:
		x.H = 0
		x.L = 0
	case bits >= numHalfBits:
		x.L = x.H >> (bits - numHalfBits)
		x.H = 0
	default:
		x.L >>= bits
		x.L |= x.H << (numHalfBits - bits)
		x.H >>= bits
	}

	return x
}

// And returns the logical AND of x and y.
func And(x, y Uint128) Uint128 {
	x.H &= y.H
	x.L &= y.L

	return x
}

// AndNot returns the logical AND NOT of x and y.
func AndNot(x, y Uint128) Uint128 {
	x.H &^= y.H
	x.L &^= y.L

	return x
}

// Not returns the logical NOT of x and o.
func Not(x Uint128) Uint128 {
	x.H = ^x.H
	x.L = ^x.L

	return x
}

// Xor returns the logical XOR of x and y.
func Xor(x, y Uint128) Uint128 {
	x.H ^= y.H
	x.L ^= y.L

	return x
}

// Or returns the logical OR of x and y.
func Or(x, y Uint128) Uint128 {
	x.H |= y.H
	x.L |= y.L

	return x
}

// Add adds x and y.
func Add(x, y Uint128) Uint128 {
	sum, _ := Add128(x, y, Zero())

	return sum
}

// Add128 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
func Add128(x, y, carry Uint128) (sum, carryOut Uint128) {
	sum.L, carryOut.L = bits.Add64(x.L, y.L, carry.L)
	sum.H, carryOut.L = bits.Add64(x.H, y.H, carryOut.L)

	return
}

// Incr increments x by one.
func Incr(x Uint128) Uint128 {
	return Add(x, Uint128{H: 0, L: 1})
}

// Sub subtracts x and y.
// nolint: ifshort, varnamelen
func Sub(x, y Uint128) Uint128 {
	pL := x.L
	x.L -= y.L
	x.H -= y.H

	if x.L > pL {
		x.H--
	}

	return x
}

// Decr decrements x by one.
// nolint: ifshort, varnamelen
func Decr(x Uint128) Uint128 {
	pL := x.L
	x.L--

	if x.L > pL {
		x.H--
	}

	return x
}

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len(x Uint128) int {
	if x.H == 0 {
		return bits.Len64(x.L)
	}

	return numHalfBits + bits.Len64(x.H)
}

// LeadingZeros returns the number of leading zero bits in x; the result is 128 for x == 0.
func LeadingZeros(x Uint128) int {
	return numBits - Len(x)
}

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x Uint128) int {
	return bits.OnesCount64(x.H) + bits.OnesCount64(x.L)
}

// TrailingZeros returns the number of trailing zero bits in x; the result is 128 for x == 0.
func TrailingZeros(x Uint128) int {
	if x.L == 0 {
		return bits.TrailingZeros64(x.H) + numHalfBits
	}

	return bits.TrailingZeros64(x.L)
}

// Reverse returns the value of x with its bits in reversed order.
func Reverse(x Uint128) (y Uint128) {
	y.L, y.H = bits.Reverse64(x.H), bits.Reverse64(x.L)

	return
}

// ReverseBytes returns the value of x with its bytes in reversed order.
func ReverseBytes(x Uint128) (y Uint128) {
	y.L, y.H = bits.ReverseBytes64(x.H), bits.ReverseBytes64(x.L)

	return
}
