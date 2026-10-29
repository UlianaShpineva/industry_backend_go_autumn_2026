package main

func rotateRunes(s string, shift int) string {
	rs := []rune(s)
	n := len(rs)
	if n == 0 {
		return ""
	}

	shift %= n
	if shift < 0 {
		shift += n
	}

	if shift == 0 {
		return string(rs)
	}

	out := make([]rune, n)
	copy(out, rs[shift:])
	copy(out[n-shift:], rs[:shift])
	return string(out)
}
