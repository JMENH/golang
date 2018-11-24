
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	stdout = new(bytes.Buffer) // captured output
	main()
	got := stdout.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Errorf("Fail result")
	}
}

func Example_one() {

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		words    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 90}, 3},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}, 8},
	}
	for _, test := range tests {
		var i IntSet
		for _, w := range test.words {
			i.Add(w)
		}
		if i.Len() != test.expected {
			t.Errorf("Result = %v, Expected %v", i.Len(), test.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		words    []int
		removes  []int
		expected string
	}{
		{[]int{}, []int{1}, "{}"},
		{[]int{1, 2, 32}, []int{1, 2, 90, 32}, "{}"},
		{[]int{1, 2, 90}, []int{2}, "{1 90}"},
		{[]int{1, 2, 3, 4, 200, 6, 7, 8}, []int{200, 4}, "{1 2 3 6 7 8}"},
	}

	for _, test := range tests {
		var i IntSet
		for _, w := range test.words {
			i.Add(w)
		}
		for _, r := range test.removes {
			i.Remove(r)
		}

		if i.String() != test.expected {
			t.Errorf("Result = %v, Expected %v", i.String(), test.expected)
		}

	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		is IntSet
	}{
		{IntSet{}},
		{IntSet{words: []uint64{1, 2, 3}}},
	}

	for _, test := range tests {
		test.is.Clear()
		if len(test.is.words) != 0 {
			t.Fatalf("Result = %q", test.is)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		expected IntSet
	}{
		{IntSet{}},
		{IntSet{words: []uint64{1, 2, 3, 4, 5, 6, 7, 8}}},
		{IntSet{words: []uint64{1, 9, 2, 9, 3, 9, 4, 9, 5, 9}}},
		{IntSet{words: []uint64{1, 2, 3, 5}}},
	}

	for _, test := range tests {
		result := test.expected.Copy()
		for i := 0; i < len(test.expected.words); i++ {
			if result.words[i] != test.expected.words[i] {
				t.Fatalf("Result = %v, Expected %v", result.words[i], test.expected.words[i])
			}
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		words    []int
		values   []int
		expected string
	}{
		{[]int{}, []int{1, 30, 400}, "{1 30 400}"},
		{[]int{1, 2, 32}, []int{}, "{1 2 32}"},
		{[]int{1, 2, 90}, []int{2}, "{1 2 90}"},
		{[]int{1, 2, 3, 4, 6, 7, 8}, []int{200}, "{1 2 3 4 6 7 8 200}"},
	}

	for _, test := range tests {
		var i IntSet
		for _, w := range test.words {
			i.Add(w)
		}

		i.AddAll(test.values...)

		if i.String() != test.expected {
			t.Errorf("Result = %v, Expected %v", i.String(), test.expected)
		}

	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
		expected  string
	}{
		{[]int{3}, []int{1, 30, 400}, "{}"},
		{[]int{1, 2, 32}, []int{1, 2, 32}, "{1 2 32}"},
		{[]int{1, 2, 90}, []int{2}, "{2}"},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}, "{2 200}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, e := range test.xelements {
			x.Add(e)
		}
		for _, e := range test.yelements {
			y.Add(e)
		}
		x.IntersectWith(&y)

		if x.String() != test.expected {
			t.Errorf("Result = %v, Expected %v", x.String(), test.expected)
		}

	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
		expected  string
	}{
		{[]int{3}, []int{1, 30, 400}, "{3}"},
		{[]int{1, 2, 32}, []int{1, 2, 32}, "{}"},
		{[]int{1, 2, 90}, []int{2}, "{1 90}"},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2}, "{1 3 4 6 7}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, e := range test.xelements {
			x.Add(e)
		}
		for _, e := range test.yelements {
			y.Add(e)
		}
		x.DifferenceWith(&y)

		if x.String() != test.expected {
			t.Errorf("Result = %v, Expected %v", x.String(), test.expected)
		}

	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		xelements []int
		yelements []int
		expected  string
	}{
		{[]int{3}, []int{1, 30, 400}, "{1 3 30 400}"},
		{[]int{1, 2, 32}, []int{1, 2, 32}, "{}"},
		{[]int{1, 2, 90}, []int{2}, "{1 90}"},
		{[]int{1, 2, 3, 4, 6, 7, 200}, []int{200, 2, 5}, "{1 3 4 5 6 7}"},
	}

	for _, test := range tests {
		var x, y IntSet
		for _, e := range test.xelements {
			x.Add(e)
		}
		for _, e := range test.yelements {
			y.Add(e)
		}
		x.SymmetricDifference(&y)

		if x.String() != test.expected {
			t.Errorf("Result = %v, Expected %v", x.String(), test.expected)
		}

	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		words []int
		c     int
	}{
		{[]int{}, 0},
		{[]int{3}, 1},
		{[]int{1, 2, 32}, 3},
		{[]int{1, 2, 90}, 3},
		{[]int{1, 2, 3, 4, 6, 7, 200}, 7},
	}

	for _, test := range tests {
		var x IntSet
		for _, e := range test.words {
			x.Add(e)
		}

		for i, e := range x.Elems() {
			if e != test.words[i] {
				t.Errorf("Result = %v, Expected %v", e, test.words[i])
			}
		}
	}
}
