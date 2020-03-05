package math

// Max returns a greater value
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Min returns a smaller value
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// Abs returns an absolute value of a
func Abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// Between returns true if c is between a and b
func Between(a, b, c int64) bool {
	min := Min(a, b)
	max := Max(a, b)
	return min <= c && c <= max
}

// GCD returns the greatest common denominator
func GCD(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM returns the least common multiple
func LCM(a int64, b int64) int64 {
	return a * b / GCD(a, b)
}
