package main

import "fmt"

func main() {
	switch {
	case '0' > '1':
		fmt.println("0 > 1!")
	case '0' < '1':
		fmt.println("0 < 1!")
	case '0' == '1':
		fmt.println("0 == 1!")
	}
}
