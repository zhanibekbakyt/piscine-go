package printnbr

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	RealPrintNbr(n)
}

func RealPrintNbr(n int) {
	x := '0'
	if n == 0 {
		z01.PrintRune(x)
		return
	}
	for i := 1; i <= n%10; i++ {
		x++
	}
	for i := -1; i >= n%10; i-- {
		x++
	}
	if n/10 != 0 {
		RealPrintNbr(n / 10)
	}
	z01.PrintRune(x)
	return
}
