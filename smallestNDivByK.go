package FamousProblemsWithGo

func SmallestNDivByK(k int) int {
	if k % 2 == 0 || k % 5 == 0 {
		return -1
	}
	var res = 1
	var r = res % k

	for res = 1; res <= k; {
		if r == 0 {
			return res
		}
		r = (r * 10 + 1) % k
		res++
	}
	return -1
}
