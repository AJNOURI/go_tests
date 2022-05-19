package generics

import (
	"testing"
)

func TestFind_int(t *testing.T) {
	var isEven = func(i int) bool {
		return i%2 == 0
	}

	testFind(t, []findTest[int]{
		{[]int{1, 2, 3}, isEven, 2, true},
		{[]int{1, 3, 5, 7, 9}, isEven, 0, false},
		{nil, isEven, 0, false},
	})
}

func TestFind_string(t *testing.T) {
	var hasLenGt3 = func(s string) bool {
		return len(s) > 3
	}

	testFind(t, []findTest[string]{
		{[]string{"foo", "bar", "bazzz"}, hasLenGt3, "bazzz", true},
		{[]string{"foo", "bar", "baz"}, hasLenGt3, "", false},
		{[]string{}, hasLenGt3, "", false},
	})
}

func testFind[T comparable](t *testing.T, tests []findTest[T]) {
	for _, test := range tests {
		result, ok := Find(test.list, test.predicate)
		if ok != test.expectedOk {
			t.Errorf("Find(%#v, %T) returned ok=%v, expected %v", test.list, test.predicate, ok, test.expectedOk)
			continue
		}
		if !ok {
			continue
		}
		if *result != test.expectedResult {
			t.Errorf("Find(%#v, %T) returned result=%v, expected %v", test.list, test.predicate, *result, test.expectedResult)
		}
	}
}

type findTest[T any] struct {
	list           []T
	predicate      func(T) bool
	expectedResult T
	expectedOk     bool
}
