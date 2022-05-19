package generics

func Find(list []int, predicate func(int) bool) (*int, bool) {
	for _, v := range list {
		if predicate(v) {
			return &v, true
		}
	}
	return nil, false
}
