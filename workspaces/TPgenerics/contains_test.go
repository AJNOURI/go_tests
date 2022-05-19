package generics

import (
	"testing"
)

func TestContains_int(t *testing.T) {
	testContains(t, []containsTest[int]{
		{[]int{1664, 1337}, 6, false},
		{[]int{1664, 1337}, 1337, true},
		{nil, 1, false},
	})
}

func TestContains_string(t *testing.T) {
	testContains(t, []containsTest[string]{
		{[]string{"foo", "bar"}, "foo", true},
		{[]string{"foo", "bar"}, "baz", false},
		{nil, "", false},
	})
}

type Point struct {
	x, y int
}

func TestContains_Point(t *testing.T) {
	testContains(t, []containsTest[Point]{
		{[]Point{{-2, 1}, {9, -12}}, Point{1, 2}, false},
		{[]Point{{-2, 1}, {9, -12}}, Point{9, -12}, true},
		{nil, Point{}, false},
	})
}

func testContains[T comparable](t *testing.T, tests []containsTest[T]) {
	for _, test := range tests {
		actual := Contains(test.haystack, test.needle)
		if actual != test.expected {
			t.Errorf("Contains(%#v, %#v) returned %#v, expected %#v", test.haystack, test.needle, actual, test.expected)
		}
	}
}

type containsTest[T comparable] struct {
	haystack []T
	needle   T
	expected bool
}
