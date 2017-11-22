// Copyright (c) 2017 go-commons. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bytes

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomFrom(t *testing.T) {
	tt := []struct {
		set  []byte
		size int
		want []byte
	}{
		// The order of tests matters!
		{nil, 0, nil},

		{[]byte("a"), 1, []byte("a")},
		{[]byte("a"), 2, []byte("aa")},
		{[]byte{0x8}, 1, []byte{0x8}},
		{[]byte{0x8}, 2, []byte{0x8, 0x8}},

		{[]byte("ab"), 1, []byte("b")},
		{[]byte("ab"), 2, []byte("bb")},
		{[]byte("ab"), 3, []byte("aaa")},
		{[]byte{0x61, 0xA}, 1, []byte{0xA}},
		{[]byte{0x61, 0xA}, 2, []byte{0x61, 0x61}},
		{[]byte{0x61, 0xA}, 3, []byte{0x61, 0xA, 0xA}},

		{[]byte("a1b"), 1, []byte("b")},
		{[]byte("a1b"), 2, []byte("1b")},
		{[]byte("a1b"), 3, []byte("11b")},
		{[]byte("a1b"), 4, []byte("b111")},
		{[]byte{0x61, 0x20, 0x62}, 1, []byte{0x61}},
		{[]byte{0x61, 0x20, 0x62}, 2, []byte{0x61, 0x61}},
		{[]byte{0x61, 0x20, 0x62}, 3, []byte{0x20, 0x62, 0x61}},
		{[]byte{0x61, 0x20, 0x62}, 4, []byte{0x20, 0x62, 0x20, 0x20}},

		{[]byte("1234"), 1, []byte("4")},
		{[]byte("1234"), 2, []byte("33")},
		{[]byte("1234"), 3, []byte("414")},
		{[]byte("1234"), 4, []byte("3422")},
		{[]byte("1234"), 5, []byte("42113")},
		{[]byte{0x31, 0x20, 0x33, 0x1B}, 1, []byte{0x31}},
		{[]byte{0x31, 0x20, 0x33, 0x1B}, 2, []byte{0x33, 0x31}},
		{[]byte{0x31, 0x20, 0x33, 0x1B}, 3, []byte{0x1b, 0x31, 0x20}},
		{[]byte{0x31, 0x20, 0x33, 0x1B}, 4, []byte{0x1b, 0x31, 0x20, 0x33}},
		{[]byte{0x31, 0x20, 0x33, 0x1B}, 5, []byte{0x20, 0x31, 0x1b, 0x31, 0x20}},

		{[]byte("0123456789"), 10, []byte("0950608567")},
		{[]byte("0123456789abcdef"), 10, []byte("fb73789108")},
		{[]byte("0123456789ABCDEF"), 20, []byte("96644CA506360F544EEE")},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), 20, []byte("zwnfvotnftneqvjitofd")},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), 30, []byte("ADTBIYHEPSVLFUHKPHFAMTAEDGQMZI")},
		{[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 30, []byte("vXVNySTpMuHDcdplgYeQOgyBFvHHnX")},
		{[]byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), 40, []byte("qVTTdCNAhv8ptjaQMYJ4iTSKEro9YKErUA37BKlh")},
	}

	rnd := rand.New(rand.NewSource(8128))
	for i, tc := range tt {
		have := make([]byte, tc.size)
		n := RandomFrom(tc.set, have, rnd)
		if n != len(have) {
			t.Errorf("#%d: RandomFrom() = %d; want %d", i, n, len(have))
		}
		if !bytes.Equal(have, tc.want) {
			t.Errorf("#%d: RandomFrom():\nhave: %#v\nwant: %#v", i, have, tc.want)
		}
	}
}

func BenchmarkRandomFrom(b *testing.B) {
	set := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rnd := rand.NewSource(int64(time.Now().Nanosecond())).(rand.Source64)
	for _, size := range [...]int{1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9} {
		dst := make([]byte, size)
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandomFrom(set, dst, rnd)
			}
		})
	}
}
