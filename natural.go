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

	for {
		var li int
		var lok bool
		if idx < leftLen {
			li, lok = parseInt(left[idx])
		}
		var ri int
		var rok bool
		if idx < rightLen {
			ri, rok = parseInt(right[idx])
		}

		idx++

		if !lok {
			if rok {
				rn = 10*rn + ri
			}

			if rn > ln || rok && ln == rn {
				return -1
			}

			// Read rest of right until it's larger than left.
			for idx < rightLen {
				n, ok := parseInt(right[idx])
				if !ok {
					break
				}
				rn = 10*rn + n
				// Greater than or equal because right is longer than left (i.e. more leading zeros).
				if rn >= ln {
					return -1
				}
				idx++
			}

			return 1
		}
		if !rok {
			if lok {
				ln = 10*ln + li
			}

			if ln > rn || lok && ln == rn {
				return 1
			}

			// Read rest of left until it's larger than right.
			for idx < leftLen {
				n, ok := parseInt(left[idx])
				if !ok {
					break
				}
				ln = 10*ln + n
				// Greater than or equal because left is longer than right (i.e. more leading zeros).
				if ln >= rn {
					return 1
				}
				idx++
			}

			return -1
		}

		ln = 10*ln + li
		rn = 10*rn + ri
	}
}

func parseInt(b byte) (int, bool) {
	i := int(b) - '0'
	return i, 0 <= i && i <= 9
}
