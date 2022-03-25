package search

// This function uses recirsive call to itself
func binarySearch(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}

	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

// Iterative method as opposed to the recursive method
func binaryIterative(array []int, target int, lowIndex int, highInded int) (int, error) {
	startIndex := lowIndex
	endIndex := highInded
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}
