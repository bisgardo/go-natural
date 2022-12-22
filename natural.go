package strcmp

// Natural compares two strings using "natural" semantics where numbers are compared in their entirety as opposed to byte for byte.
func Natural(left, right string) int {
	leftLen := len(left)
	rightLen := len(right)

	minLen := leftLen
	if minLen > rightLen {
		minLen = rightLen
	}

	idx := 0
	for idx < minLen {
		l := left[idx]
		r := right[idx]
		// Always skip equal bytes.
		if l == r {
			idx++
			continue
		}
		if isDigit(l) || isDigit(r) {
			// Skip any leading zeros.
			leftStart := idx
			for l == '0' {
				leftStart++
				if leftStart == leftLen {
					break
				}
				l = left[leftStart]
			}
			rightStart := idx
			for r == '0' {
				rightStart++
				if rightStart == rightLen {
					break
				}
				r = right[rightStart]
			}

			// Find end of numbers "in parallel"; stopping as soon as one turns out to be longer than the other
			// (could do fast path when leftStart==rightStart).
			leftEnd := leftStart
			rightEnd := rightStart
			for {
				leftIsDigit := leftEnd < leftLen && isDigit(left[leftEnd])
				rightIsDigit := rightEnd < rightLen && isDigit(right[rightEnd])
				if !leftIsDigit && !rightIsDigit {
					break
				}
				if !leftIsDigit && rightIsDigit {
					// Right number is longer.
					return -1
				}
				if leftIsDigit && !rightIsDigit {
					// Left number is longer.
					return 1
				}
				leftEnd++
				rightEnd++
			}

			// Numbers have the same length; compare as standard strings.
			leftNumStr := left[leftStart:leftEnd]
			rightNumStr := right[rightStart:rightEnd]
			if leftNumStr < rightNumStr {
				return -1
			}
			if leftNumStr > rightNumStr {
				return 1
			}

			// Numbers are equal. Let the one with the most leading zeros win.
			if leftStart < rightStart {
				return -1
			}
			if leftStart > rightStart {
				return 1
			}

			// Everything is equal. Skip what we read and start over.
			idx = leftEnd
			continue
		}
		// Do standard comparison of bytes.
		if l < r {
			return -1
		}
		// We already know that 'l' and 'r' are not equal.
		if l > r {
			return 1
		}
		idx++
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

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
