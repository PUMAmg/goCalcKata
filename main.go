package main

import "fmt"

func main() {
	fmt.Println("Hello i am calculator. Right now I'll show 4 operations with arabic numbers:")
	a, b := 8, 4
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}
