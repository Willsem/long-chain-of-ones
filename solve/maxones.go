package solve

func MaxOnesAfterRemoveItem(sequence []byte) uint {
	// Длина наибольшей последовательности единиц с удалением нуля
	var maxLen uint = 0

	// Длина текущей последовательности единиц с учетом удаления нуля
	var lenOfOnesWithoutZero uint = 0

	// Длина текущей последовательности единиц
	var lenOfOnes uint = 0

	n := len(sequence)
	for i := 0; i < n; i++ {
		if sequence[i] == 1 {
			lenOfOnes++
			lenOfOnesWithoutZero++
		} else {
			if lenOfOnesWithoutZero > maxLen {
				maxLen = lenOfOnesWithoutZero
			}

			lenOfOnesWithoutZero = lenOfOnes
			lenOfOnes = 0
		}
	}

	if lenOfOnesWithoutZero > maxLen {
		maxLen = lenOfOnesWithoutZero
	}

	// В последовательности нет нулей
	if maxLen == uint(n) {
		return maxLen - 1
	}

	return maxLen
}
