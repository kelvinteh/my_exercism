package hamming

import "errors"

// Distance takes 2 strings sequences and returns the number of differences
func Distance(a, b string) (int, error) {

	// early exit if the strings don't match in length since no way to do comparison
	if len(a) != len(b) {
		return 0, errors.New("Both strings are of unequal length")
	}

	ham := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ham++
		}
	}
	return ham, nil
}
