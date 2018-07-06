package sorts

import (
	"sort"
)

func Bubble(sortable sort.Interface) {
	if sortable.Len() < 2 {
		return
	}

	sortedFromIndex := sortable.Len() - 1

	for sortedFromIndex > 0 {

		for i := 0; i < sortedFromIndex; i++ {
			current := i
			next := i + 1

			if sortable.Less(next, current) {
				sortable.Swap(next, current)
			}
		}

		sortedFromIndex--

	}

}
