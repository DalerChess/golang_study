package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_print_comb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				ft_putchar(a)
				ft_putchar(b)
				ft_putchar(c)

				if !(a == '7' && b == '8' && c == '9') {
					os.Stdout.Write([]byte(", "))
				}
			}
		}
	}
}

func main() {
	ft_print_comb()
	ft_putchar('\n')
}
