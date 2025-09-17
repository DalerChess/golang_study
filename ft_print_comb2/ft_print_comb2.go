package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_putnbr2(n int) {
	ft_putchar(rune(n/10 + '0'))
	ft_putchar(rune(n%10 + '0'))
}

func ft_print_comb2() {
	for a := 0; a <= 98; a++ {
		for b := 1; b <= 99; b++ {
			ft_putnbr2(a)
			ft_putchar(' ')
			ft_putnbr2(b)

			if !(a == 98 && b == 99) {
				os.Stdout.Write([]byte(", "))
			}
		}
	}
}

func main() {
	ft_print_comb2()
	ft_putchar('\n')
}
