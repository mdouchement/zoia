package caption

// IsValid checks if the given cation has valid characters and maximum length to 15.
func IsValid(caption []byte, n int) bool {
	if len(caption) > n {
		return false
	}

	for _, b := range caption {
		_, ok := AlphabetByte[b]
		if !ok {
			return false
		}
	}
	return true
}

// IsValidString checks if the given cation has valid characters and maximum length to 15.
func IsValidString(caption string, n int) bool {
	if len(caption) > n {
		return false
	}

	for _, r := range caption {
		_, ok := AlphabetRune[r]
		if !ok {
			return false
		}
	}
	return true
}

// Trim removes all incompatible characters.
func Trim(caption []byte) (dst []byte) {
	for _, b := range caption {
		if AlphabetByte[b] {
			dst = append(dst, b)
		}
	}
	return dst
}

// TrimToString removes all incompatible characters.
func TrimToString(caption []byte) string {
	return string(Trim(caption))
}

// TrimString removes all incompatible characters.
func TrimString(caption string) string {
	return string(Trim([]byte(caption)))
}

// Raw formats the given caption to raw format.
func Raw(caption []byte, n int) []byte {
	dst := make([]byte, n)
	for i := range dst {
		if i < len(caption) {
			dst[i] = caption[i]
		}
	}

	return dst
}
