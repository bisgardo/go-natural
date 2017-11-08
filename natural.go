package strcmp

// Natural compares two strings naturally.
func Natural(left, right string) int {
	leftLen := len(left)
	rightLen := len(right)

	minLen := leftLen
	if minLen > rightLen {
		minLen = rightLen
	}

	for idx := 0; idx < minLen; idx++ {
		l := left[idx]
		r := right[idx]
		if l != r {
			return innerCompare(l, r, left, right, idx+1)
		}
	}

	// One string is a prefix of the other - longer is greater.
	if leftLen < rightLen {
		return -1
	}
	if leftLen > rightLen {
		return 1
	}

	// Strings are equal.
	return 0
}

func innerCompare(l, r byte, left, right string, idx int) int {
	// Bytes l and r are assumed to be different.
	li, lok := parseInt(l)
	ri, rok := parseInt(r)

	if !lok && !rok {
		// Both are non-numbers. Compare as bytes.
		if l < r {
			return -1
		}
		return 1
	}

	// Any number character is "larger" than any non-number one.
	if !lok {
		return -1
	}
	if !rok {
		return 1
	}

	// Both are numbers.
	return innerCompareRemaining(left, right, int64(li), int64(ri), idx)
}

func innerCompareRemaining(left, right string, leftNum, rightNum int64, idx int) int {
	leftLen := len(left)
	rightLen := len(right)
	for {
		var li, ri int8
		var lok, rok bool
		if idx < leftLen {
			li, lok = parseInt(left[idx])
		}
		if idx < rightLen {
			ri, rok = parseInt(right[idx])
		}

		if !lok && !rok {
			if leftNum < rightNum {
				return -1
			}
			return 1
		}

		idx++

		if !lok {
			rightNum = 10*rightNum + int64(ri)

			if rightNum >= leftNum {
				return -1
			}

			// Read rest of right until it's larger than left.
			for idx < rightLen {
				n, ok := parseInt(right[idx])
				if !ok {
					break
				}
				rightNum = 10*rightNum + int64(n)
				// Include '=' because right is longer than left (i.e. more leading zeros).
				if rightNum >= leftNum {
					return -1
				}
				idx++
			}

			// After reading all of right, rightNum remains strictly smaller than leftNum.
			return 1
		}
		if !rok {
			leftNum = 10*leftNum + int64(li)

			if leftNum >= rightNum {
				return 1
			}

			// Read rest of left until it's larger than right.
			for idx < leftLen {
				n, ok := parseInt(left[idx])
				if !ok {
					break
				}
				leftNum = 10*leftNum + int64(n)
				// Include '=' because left is longer than right (i.e. more leading zeros).
				if leftNum >= rightNum {
					return 1
				}
				idx++
			}

			// After reading all of left, leftNum remains strictly smaller than rightNum.
			return -1
		}

		leftNum = 10*leftNum + int64(li)
		rightNum = 10*rightNum + int64(ri)
	}
}

func parseInt(b byte) (int8, bool) {
	i := int8(b) - '0'
	return i, 0 <= i && i <= 9
}
