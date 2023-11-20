package goutils

// Keys returns a slice composed of map keys
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, len(in), len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

// Values returns a slice composed of map values
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}
