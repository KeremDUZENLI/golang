package Challenge18_MinRotations

func MinRotations(array []int) int {
	// Implement me :)
	size := len(array)
	if size <= 1 {
		return 0
	}

	lowIndex := 0
	middleIndex := 0
	highIndex := size - 1
	for array[lowIndex] > array[highIndex] {
		middleIndex = (lowIndex + highIndex) / 2
		if array[middleIndex] > array[highIndex] {
			lowIndex = middleIndex + 1
		} else {
			highIndex = middleIndex
		}
	}
	return lowIndex
}
