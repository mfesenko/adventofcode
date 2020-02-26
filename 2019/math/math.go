package math

// Max returns a greater value
func Max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Min returns a smaller value
func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Abs returns an absolute value of a
func Abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

// Between returns true if c is between a and b
func Between(a, b, c int32) bool {
	min := Min(a, b)
	max := Max(a, b)
	return min <= c && c <= max
}

// GCD returns the greatest common denominator
func GCD(a int32, b int32) int32 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
