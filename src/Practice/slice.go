package main

import "fmt"

var months = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func main() {
	Q1 := months[4:7]
	Q2 := months[6:9]
	for _, s := range Q1 {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}
