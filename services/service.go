package services

import (
	"find_pairs/models"
)

func FindPairs(i models.Input) (models.Solutions, error) {
	// map to store pairs of indices with desired sum
	var result models.Solutions

	// map to store indices of all elements
	m := make(map[int]int)

	// traversing through all the elements
	for j := 0; j < len(i.Numbers); j++ {
		rem := i.Target - i.Numbers[j]

		// check if a pair can be formed with current index j
		idx, found := m[rem]
		if found || (found && m[i.Numbers[j]] == 0) {
			result = append(result, []int{idx, j})
		}

		// storing the current index in m
		m[i.Numbers[j]] = j
	}

	return result, nil
}
