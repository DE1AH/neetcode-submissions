func maxArea(h []int) int {
	l, r := 0, len(h) - 1
	max_V := -10000
	for l < r{
		if h[l] <= h[r]{
			max_V = max(max_V, h[l] * (r-l))
			l++
		} else {
			max_V = max(max_V, h[r] * (r-l))
			r --
		}
	}
	return max_V
}
