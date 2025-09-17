package main

import "os"

func ft_putchar(c rune) {
	os.Stdout.Write([]byte(string(c)))
}

func ft_print_digits(comb []int, n int) {
	for i := 0; i < n; i++ {
		ft_putchar(rune(comb[i] + '0'))
	}
}

func ft_print_combn(n int) {
	comb := make([]int, n)
	for i := 0; i < n; i++ {
		comb[i] = i
	}
	for {
		ft_print_digits(comb, n)
		if comb[0] == 10-n {
			break
		}
		os.Stdout.Write([]byte(", "))

		i := n - 1
		for i >= 0 && comb[i] == 9-(n-1-i) {
			i--
		}

		comb[i]++

		for j := i + 1; j < n; j++ {
			comb[j] = comb[j-1] + 1
		}
	}
}

func main() {
	ft_print_combn(2)
	ft_putchar('\n')
	ft_print_combn(3)
	ft_putchar('\n')
}
