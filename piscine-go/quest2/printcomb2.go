package printcomb2

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	a := '0'
	b := '0'
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {

			if a == '9' && b == '8' {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(' ')
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune('\n')
				break
			}

			if i == '0' && j == '0' {
				j++
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(' ')
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(' ')
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

			if i == '9' && j == '9' {
				if b < '9' {
					b++
				} else if b == '9' {
					b = '0'
					a++
				}
				i = a
				j = b
			}
		}
	}
}
