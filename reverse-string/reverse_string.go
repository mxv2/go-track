// Package reverse provides implementation for string reverse operation.
package reverse

// Reverse reverses input string.
func Reverse(input string) string {
	output := []rune(input)

	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return string(output)
}
