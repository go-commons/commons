// Copyright (c) 2017 go-commons. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bytes

import (
	"math/bits"
	"math/rand"
)

// RandomFrom fills dst with randomly chosen bytes from the set.
// It returns the number of bytes written into dst, which is len(dst).
// The function will panic if len(set) == 0.
func RandomFrom(set, dst []byte, src rand.Source64) (n int) {
	switch {
	case len(dst) == 0:
		return 0
	case len(set) == 0:
		panic("at least one byte in set must be specified")
	}

	idxBits := uint64(bits.Len(uint(len(set)))) // Number of bits needed to represent all bytes from the set.
	idxMask := uint64(1<<idxBits) - 1           // Mask allows to pass only bits that can represent max number from idxBits.
	rndMax := 64 / idxBits                      // Max count of integers that can be extracted from one random uint64.

	for u, c := src.Uint64(), rndMax; n < len(dst); u, c = u>>idxBits, c-1 {
		if c == 0 {
			u, c = src.Uint64(), rndMax
		}
		if i := int(u & idxMask); i < len(set) {
			dst[n] = set[i]
			n++
		}
	}
	return n
}
