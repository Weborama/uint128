# github.com/weborama/uint128

[![Go Report Card](https://goreportcard.com/badge/github.com/weborama/uint128)](https://goreportcard.com/report/github.com/weborama/uint128)[![GolangCI Report Card](https://github.com/golangci/golangci-web/blob/master/src/assets/images/badge_a_plus_flat.svg)](https://golangci.com/r/github.com/weborama/uint128)

This package provides an implementation of a 128 bit unsigned integer with some
implementations of:

- arithmetic operators (Add, Sub, Incr, Decr)
- binary operations (AND, OR, XOR, NOT, AND NOT, shifts)
- math/bits operations (LeadingZeros, Len, OnesCount, Reverse, ReverseBytes, TrailingZeros)

Missing operators (Mult, Div, Mod, RotateLeft etc.) to be added later.

Uses a modified copy (changes to the generated package declaration) of make_tables.go from [math/bits](https://github.com/golang/go/tree/master/src/math/bits) as well as code sourced from [github.com/mengzhuo/uint128](https://github.com/mengzhuo/uint128) and [github.com/davidminor/uint128](https://github.com/davidminor/uint128).

Original use case for this package is implementing IPv6 calculations (see
[github.com/weborama/cidr](http://github.com/weborama/cidr)).
