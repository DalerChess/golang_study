package main

import "fmt"

func ft_putchar(c rune) {
	fmt.Printf("%c", c)
}

func ft_print_reverse_alphabet() {
	c := 'z'
	for c >= 'a' {
		ft_putchar(c)
		c--
	}
}

func main() {
	ft_print_reverse_alphabet()
	fmt.Println()
}
