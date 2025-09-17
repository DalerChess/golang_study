package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_putnbr(nb int) {
	if nb == -2147483648 {
		os.Stdout.Write([]byte("-2147483648"))
		return
	}

	if nb < 0 {
		ft_putchar('-')
		nb = -nb
	}

	if nb > 9 {
		ft_putnbr(nb / 10)
	}
	ft_putchar(rune(nb%10 + '0'))
}

func main() {
	ft_putnbr(5)
	ft_putchar('\n')
	ft_putnbr(0)
	ft_putchar('\n')
	ft_putnbr(-2147483648)
	ft_putchar('\n')
}
