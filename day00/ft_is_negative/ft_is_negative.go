package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_is_negative(n int) {
	if n > 0 {
		ft_putchar('P')
	} else {
		ft_putchar('N')
	}
}

func main() {
	ft_is_negative(-5)
	ft_putchar('\n')
	ft_is_negative(0)
	ft_putchar('\n')
	ft_is_negative(42)
	ft_putchar('\n')
}
