package main

import "fmt"

func ft_putchar(c rune) {
	fmt.Printf("%c", c)
}

func main() {
	ft_putchar('a')
	ft_putchar('\n')
}
