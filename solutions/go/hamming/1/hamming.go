package hamming

import "errors"

// Distance calculate the differences between two DNA strings
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("both strings should have the same length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
