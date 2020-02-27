package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	for {
		n, _ := fmt.Scan(&s)
		if n == 0 {
			break
		} else {
			count := 0
			for i := len(s) - 1; i >= 0; i-- {
				if s[i] == byte(32) {
					fmt.Println(count)
					break
				}
				count++
			}
			fmt.Println(s)
			s = "13213\n"
			b, err := strconv.Atoi(s[:len(s)-1])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(b)
		}
	}
}
