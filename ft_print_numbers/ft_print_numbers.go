package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_print_numbers() {
	c := '0'
	for c <= '9' { // аналог while
		ft_putchar(c)
		c++
	}
}

func main() {
	ft_print_numbers()
	ft_putchar('\n') // перенос строки
}
