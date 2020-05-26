package listops

type binFunc func(int, int) int

type predFunc func(int) bool

type unaryFunc func(int) int

// IntList is alias for int slice
type IntList []int

// Length returns count of list elements
func (il IntList) Length() int {
	return len(il)
}

// Append appends elements from passed list to called list and returns new list
func (il IntList) Append(list IntList) IntList {
	return append(il, list...)
}

// Concat returns list with the elements from all the lists
func (il IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		il = il.Append(list)
	}

	return il
}

// Foldl returns value calculated by applying `fn` to the accumulator `acc`
// and each elements from list beginning from left side
func (il IntList) Foldl(fn binFunc, acc int) int {
	for _, val := range il {
		acc = fn(acc, val)
	}

	return acc
}

// Foldr returns value calculated by applying `fn` to each elements from list
// and the accumulator `acc` beginning from right side
func (il IntList) Foldr(fn binFunc, acc int) int {
	for i := il.Length() - 1; i >= 0; i-- {
		val := il[i]
		acc = fn(val, acc)
	}

	return acc
}

// Filter returns list with the elements which passed filter function `fn`
func (il IntList) Filter(fn predFunc) IntList {
	res := IntList{}
	for _, val := range il {
		if fn(val) {
			res = append(res, val)
		}
	}

	return res
}

// Map returns list with the elements on which was called function `fn`
func (il IntList) Map(fn unaryFunc) IntList {
	res := make([]int, 0, il.Length())
	for _, val := range il {
		res = append(res, fn(val))
	}

	return res
}

// Reverse returns original list in reverse order
func (il IntList) Reverse() IntList {
	res := make([]int, 0, il.Length())
	for i := il.Length() - 1; i >= 0; i-- {
		res = append(res, il[i])
	}

	return res
}
