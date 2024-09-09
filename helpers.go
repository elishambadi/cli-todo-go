package main

// prepend adds new values to beginning of array
//
// Parameters:
//   - slice: existing slice
//   - values: values to prepend
//
// Returns:
//   - the new slice as array of strings
//
// Usage example:
//
//	sliceA := []string {"B","C","D"}
//	sliceB := []string {"A"}
//	combinedSlice := prepend(sliceA, sliceB)
func Prepend(slice, values []string) []string {
	newSlice := make([]string, len(slice)+len(values))

	copy(newSlice, values)
	// copies from where newSlice ended to end, effectively appends
	copy(newSlice[len(values):], slice)

	return newSlice
}
