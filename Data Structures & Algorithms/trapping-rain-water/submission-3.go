func trap(h []int) int {
	v := 0
	l, r := 0, len(h) - 1
	l_m, r_m := h[l], h[r]
	for l < r{
		if l_m < r_m{
			l ++
			l_m = max(l_m, h[l])
			v += l_m - h[l]
		} else {
			r --
			r_m = max(r_m, h[r])
			v += r_m - h[r]
		}
	}
	return v
}
