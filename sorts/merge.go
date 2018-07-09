package sorts

func Merge(s []int) {

	if len(s) < 2 {
		return
	}

	half := len(s) / 2
	left, right := s[0:half], s[half:]

	Merge(left)
	Merge(right)

	sortRelativeSortedSlices(left, right)
}

func sortRelativeSortedSlices(left, right []int) {

	lLength, rLength := len(left), len(right)
	totalLength := lLength + rLength
	merged := make([]int, totalLength)

	for tIndex, lIndex, rIndex := 0, 0, 0; tIndex < totalLength; tIndex++ {

		var lValue, rValue *int

		if lIndex < lLength {
			lValue = &left[lIndex]
		}

		if rIndex < rLength {
			rValue = &right[rIndex]
		}

		if lValue != nil && rValue == nil {
			merged[tIndex] = *lValue
			lIndex++
			continue
		}

		if rValue != nil && lValue == nil {
			merged[tIndex] = *rValue
			rIndex++
			continue
		}

		if *lValue <= *rValue {
			merged[tIndex] = *lValue
			lIndex++
		} else {
			merged[tIndex] = *rValue
			rIndex++
		}

	}

	for i, lIndex, rIndex := 0, 0, 0; i < len(merged); i++ {
		value := merged[i]

		if lIndex < lLength {
			left[lIndex] = value
			lIndex++
		} else {
			right[rIndex] = value
			rIndex++
		}

	}

}
