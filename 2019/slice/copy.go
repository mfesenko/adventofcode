package slice

func CopyInt64(a []int64) []int64 {
	if a == nil {
		return nil
	}

	c := make([]int64, len(a))
	copy(c, a)
	return c
}
