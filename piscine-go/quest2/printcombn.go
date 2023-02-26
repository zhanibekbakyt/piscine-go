package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb(n int) {
	switch n {
	case 1:
		PrintComb1()
	case 2:
		PrintComb2()
	case 3:
		PrintComb3()
	case 4:
		PrintComb4()
	case 5:
		PrintComb5()
	case 6:
		PrintComb6()
	case 7:
		PrintComb7()
	case 8:
		PrintComb8()
	case 9:
		PrintComb9()
	}
}

func PrintComb1() {
	for i := '0'; i <= '9'; i++ {
		if i < '9' {
			z01.PrintRune(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune(i)
			z01.PrintRune('\n')
		}
	}
}

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			if i <= '8' && j <= '9' {
				if i == '8' && j == '9' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune('\n')
					break
				} else {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}

func PrintComb3() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if i <= '7' && j <= '8' && k <= '9' {
					if i == '7' && j == '8' && k == '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune('\n')
						break
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func PrintComb4() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					if i <= '6' && j <= '7' && k <= '8' && e <= '9' {
						if i == '6' && j == '7' && k == '8' && e == '9' {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(e)
							z01.PrintRune('\n')
							break
						} else {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(e)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}

func PrintComb5() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					for f := e + 1; f <= '9'; f++ {
						if i <= '5' && j <= '6' && k <= '7' && e <= '8' && f <= '9' {
							if i == '5' && j == '6' && k == '7' && e == '8' && f == '9' {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(e)
								z01.PrintRune(f)
								z01.PrintRune('\n')
								break
							} else {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(e)
								z01.PrintRune(f)
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb6() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					for f := e + 1; f <= '9'; f++ {
						for g := f + 1; g <= '9'; g++ {
							if i <= '4' && j <= '5' && k <= '6' && e <= '7' && f <= '8' && g <= '9' {
								if i == '4' && j == '5' && k == '6' && e == '7' && f == '8' && g == '9' {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(g)
									z01.PrintRune('\n')
									break
								} else {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(g)
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb7() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					for f := e + 1; f <= '9'; f++ {
						for g := f + 1; g <= '9'; g++ {
							for l := g + 1; l <= '9'; l++ {
								if i <= '3' && j <= '4' && k <= '5' && e <= '6' && f <= '7' && g <= '8' && l <= '9' {
									if i == '3' && j == '4' && k == '5' && e == '6' && f == '7' && g == '8' && l == '9' {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune(l)
										z01.PrintRune('\n')
										break
									} else {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune(l)
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb8() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					for f := e + 1; f <= '9'; f++ {
						for g := f + 1; g <= '9'; g++ {
							for l := g + 1; l <= '9'; l++ {
								for m := l + 1; m <= '9'; m++ {
									if i <= '2' && j <= '3' && k <= '4' && e <= '5' && f <= '6' && g <= '7' && l <= '8' && m <= '9' {
										if i == '2' && j == '3' && k == '4' && e == '5' && f == '6' && g == '7' && l == '8' && m == '9' {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune('\n')
											break
										} else {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb9() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for e := k + 1; e <= '9'; e++ {
					for f := e + 1; f <= '9'; f++ {
						for g := f + 1; g <= '9'; g++ {
							for l := g + 1; l <= '9'; l++ {
								for m := l + 1; m <= '9'; m++ {
									for n := m + 1; n <= '9'; n++ {
										if i <= '1' && j <= '2' && k <= '3' && e <= '4' && f <= '5' && g <= '6' && l <= '7' && m <= '8' && n <= '9' {
											if i == '1' && j == '2' && k == '3' && e == '4' && f == '5' && g == '6' && l == '7' && m == '8' && n == '9' {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(e)
												z01.PrintRune(f)
												z01.PrintRune(g)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(n)
												z01.PrintRune('\n')
												break
											} else {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(e)
												z01.PrintRune(f)
												z01.PrintRune(g)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(n)
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
