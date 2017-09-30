package strcmp

// Natural compares two strings by comparing embedded numbers as numbers.
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
		idx++
		if l == r {
			continue
		}

		// Bytes l and r differ.

		ln, lok := parseInt(l)
		rn, rok := parseInt(r)
		if !lok || !rok {
			// At least one of the bytes is not a number.
			// Compare as bytes.
			if l < r {
				return -1
			}
			return 1
		}

		// Both bytes are numbers.
		// Resolve and compare these numbers.
		li := idx
		for li < leftLen {
			n, ok := parseInt(left[li])
			if !ok {
				break
			}
			li++
			ln = 10*ln + n
		}
		ri := idx
		for ri < rightLen {
			n, ok := parseInt(right[ri])
			if !ok {
				break
			}
			ri++
			rn = 10*rn + n
		}

		// Compare numbers.
		if ln < rn {
			return -1
		}
		if ln > rn {
			return 1
		}

		// Numbers are equal, so one of them muast have leading zeros.
		// Compare indices - higher is greater.
		if li < ri {
			return -1
		}
		return 1
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

func parseInt(b byte) (int, bool) {
	i := int(b) - '0'
	return i, 0 <= i && i <= 9
}
