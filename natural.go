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
			return innerCompare(l, r, left, right, idx+1, minLen)
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

func innerCompare(l, r byte, left, right string, idx, minLen int) int {
	// Bytes l and r are assumed to be different.
	ln, lok := parseInt(l)
	rn, rok := parseInt(r)

	// Any number character is "larger" than any non-number one.
	if !lok && rok {
		return -1
	}
	if lok && !rok {
		return 1
	}

	leftLen := len(left)
	rightLen := len(right)

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

func parseInt(b byte) (int, bool) {
	i := int(b) - '0'
	return i, 0 <= i && i <= 9
}
