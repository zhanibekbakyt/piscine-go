package piscine

func BasicAtoi2(s string) int {
	_s := []byte(s)
	count := 0
	for range _s {
		count += 1
	}
	result := 0
	pow := 1
	for i := count - 1; i >= 0; i-- {
		if int(_s[i]) > 58 || int(_s[i]) < 48 {
			return 0
		}
		result += ((int(_s[i]) - 48) * pow)
		pow *= 10
	}
	return result
}
