package main

import "fmt"

func ft_putchar(c rune) {
	fmt.Printf("%c", c)
}

func ft_print_alphabet() {
	c := 'a'
	for c <= 'z' {
		ft_putchar(c)
		c++
	}
}

func main() {
	ft_print_alphabet()
	fmt.Println()
}
