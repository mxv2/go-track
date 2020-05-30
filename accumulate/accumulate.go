// Package accumulate provides functional map
package accumulate

// Accumulate applies `binaryFn` function to each elements
// from `in` array and returns array with new elements.
func Accumulate(in []string, binaryFn func(string) string) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = binaryFn(v)
	}

	return out
}
