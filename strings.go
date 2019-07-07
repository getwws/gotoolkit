package gotoolkit

func isUpperASCII(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isLowerASCII(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func toUpperASCII(c byte) byte {
	if isLowerASCII(c) {
		return c - ('a' - 'A')
	}
	return c
}

func toLowerASCII(c byte) byte {
	if isUpperASCII(c) {
		return c + 'a' - 'A'
	}
	return c
}
