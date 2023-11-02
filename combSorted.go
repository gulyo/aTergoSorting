package main

func combSorted(a, b []WordWeight) []WordWeight {
	var indexA, indexB int = 0, 0

	var result []WordWeight

	for indexA < len(a) && indexB < len(b) {
		if a[indexA].lessThan(b[indexB]) {
			result = append(result, a[indexA])
			indexA++
		} else {
			result = append(result, b[indexB])
			indexB++
		}
	}
	if len(a) > indexA {
		result = append(result, a[indexA:]...)
	}
	if len(b) > indexB {
		result = append(result, b[indexB:]...)
	}

	return result
}
