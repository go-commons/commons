// Copyright (c) 2017 go-commons. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"math/bits"
	"math/rand"
)

// RandomFrom fills dst with randomly chosen runes from the set.
// It returns the number of runes written into dst, which is len(dst).
// The function will panic if len(set) == 0.
func RandomFrom(set string, dst []rune, src rand.Source64) (n int) {
	switch {
	case len(dst) == 0:
		return 0
	case len(set) == 0:
		panic("at least one character must be specified")
	}

	runes := []rune(set)
	idxBits := uint64(bits.Len(uint(len(runes)))) // Number of bits needed to represent all runes from the set.
	idxMask := uint64(1<<idxBits) - 1             // Mask allows to pass only bits that can represent max number from idxBits.
	rndMax := 64 / idxBits                        // Max count of integers that can be extracted from one random uint64.

	for u, c := src.Uint64(), rndMax; n < len(dst); u, c = u>>idxBits, c-1 {
		if c == 0 {
			u, c = src.Uint64(), rndMax
		}
		if i := int(u & idxMask); i < len(runes) {
			dst[n] = runes[i]
			n++
		}
	}
	return n
}
