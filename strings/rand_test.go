// Copyright (c) 2017 go-commons. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package strings

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandomFrom(t *testing.T) {
	tt := []struct {
		set  string
		size int
		want string
	}{
		// The order of tests matters!
		{"", 0, ""},

		{"a", 1, "a"},
		{"a", 2, "aa"},
		{"🎈", 1, "🎈"},
		{"🎈", 2, "🎈🎈"},

		{"ab", 1, "b"},
		{"ab", 2, "bb"},
		{"ab", 3, "aaa"},
		{"a🦋", 1, "🦋"},
		{"a🦋", 2, "aa"},
		{"a🦋", 3, "a🦋🦋"},

		{"a1b", 1, "b"},
		{"a1b", 2, "1b"},
		{"a1b", 3, "11b"},
		{"a1b", 4, "b111"},
		{"a🍭b", 1, "a"},
		{"a🍭b", 2, "aa"},
		{"a🍭b", 3, "🍭ba"},
		{"a🍭b", 4, "🍭b🍭🍭"},

		{"1234", 1, "4"},
		{"1234", 2, "33"},
		{"1234", 3, "414"},
		{"1234", 4, "3422"},
		{"1234", 5, "42113"},
		{"1⛄3🥕", 1, "1"},
		{"1⛄3🥕", 2, "31"},
		{"1⛄3🥕", 3, "🥕1⛄"},
		{"1⛄3🥕", 4, "🥕1⛄3"},
		{"1⛄3🥕", 5, "⛄1🥕1⛄"},

		{"0123456789", 10, "0950608567"},
		{"0123456789abcdef", 10, "fb73789108"},
		{"0123456789ABCDEF", 20, "96644CA506360F544EEE"},
		{"abcdefghijklmnopqrstuvwxyz", 20, "zwnfvotnftneqvjitofd"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 30, "ADTBIYHEPSVLFUHKPHFAMTAEDGQMZI"},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 30, "vXVNySTpMuHDcdplgYeQOgyBFvHHnX"},
		{"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 40, "qVTTdCNAhv8ptjaQMYJ4iTSKEro9YKErUA37BKlh"},
	}

	rnd := rand.New(rand.NewSource(8128))
	for i, tc := range tt {
		have := make([]rune, tc.size)
		n := RandomFrom(tc.set, have, rnd)
		if n != len(have) {
			t.Errorf("#%d: RandomFrom() = %d; want %d", i, n, len(have))
		}
		if string(have) != tc.want {
			t.Errorf("#%d: RandomFrom():\nhave: %#v\nwant: %#v", i, have, tc.want)
		}
	}
}

func BenchmarkRandomFrom(b *testing.B) {
	const set = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rnd := rand.NewSource(int64(time.Now().Nanosecond())).(rand.Source64)
	for _, size := range [...]int{1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9} {
		dst := make([]rune, size)
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandomFrom(set, dst, rnd)
			}
		})
	}
}
