// Implement these additional methods:
//  `func (*IntSet) Len() int      // return the number of elements`
//  `func (*IntSet) Remove(x int)  // remove x from the set`
//  `func (*IntSet) Clear()        // remove all elements from the set`
//  `func (*IntSet) Copy() *IntSet // return a copy of the set`
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the number of elements by running a population count (by clearing)
// on each word.
func (s *IntSet) Len() int {
	count := 0
	for _, w := range s.words {
		for w != 0 {
			w = w & (w - 1) // clear rightmost non-zero bit
			count++
		}
	}
	return count
}

// Remove removes `x` from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil // is this a cheat? :thinking_face:
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	ns := new(IntSet)
	for _, word := range s.words {
		ns.words = append(ns.words, word)
	}
	return ns
}

// String returns the set as a string of the form "{1 2 3}".
// TODO re-write using strings.Builder?
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Printf("%s has %d items\n", x.String(), x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Printf("%s has %d items\n", y.String(), y.Len())

	x.UnionWith(&y)
	fmt.Printf("%s union has %d items\n", x.String(), x.Len())

	fmt.Println(x.Has(9), x.Has(123))

	x.Clear()
	fmt.Printf("%s is cleared\n", x.String())

	z := y.Copy()
	y.Remove(9)
	fmt.Printf("%s has %d items\n", y.String(), y.Len())

	fmt.Printf("%s copy has %d items\n", z.String(), z.Len())

}
