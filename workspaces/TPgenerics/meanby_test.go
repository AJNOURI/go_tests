package generics

import (
	"testing"
)

func TestMeanBy_students(t *testing.T) {
	type Student struct {
		Name  string
		Grade float32
	}

	var students = []Student{
		{"John", 12.5},
		{"Brenda", 17},
		{"Max", 15.5},
	}

	var studentGrade = func(s Student) float32 {
		return s.Grade
	}

	testMeanBy(t, students, studentGrade, 15)
}

func TestMeanBy_cities(t *testing.T) {
	type Celsius float32

	type City struct {
		Name        string
		Temperature Celsius
	}

	var cities = []City{
		{"Nantes", 29},
		{"Paris", 34.5},
		{"Nice", 58},
	}

	var cityTemperature = func(c City) Celsius {
		return c.Temperature
	}

	testMeanBy(t, cities, cityTemperature, 40.5)
}

func testMeanBy[T any, R testNumber](t *testing.T, list []T, mapper func(T) R, expected R) {
	actual := MeanBy(list, mapper)
	if actual != expected {
		t.Errorf("MeanBy(%#v, %T) returned ok=%#v, expected %#v", list, mapper, actual, expected)
	}

}

type testNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}
