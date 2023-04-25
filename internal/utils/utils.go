package utils

func IntToRunes(a int) []rune {
	var b []rune
	for a > 0 {
		b = append(b, 48+rune(a%10))
		a = a / 10
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}

func BytesToRunes(b []byte) []rune {
	res := make([]rune, len(b))
	for i, v := range b {
		res[i] = rune(v)
	}
	return res
}

func RunesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
