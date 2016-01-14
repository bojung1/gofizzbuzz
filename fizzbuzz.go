package main

import "fmt"
// Herp is mod 3 0, Derp is mod 5 0

func main() {
	for i := 0; i <= 1000000; i++ {
		switch {
				case i%3 == 0 && i%5 == 0:
					fmt.Println("FizzBuzz")
				case i%3 == 0 && i%5 != 0:
					fmt.Println("Fizz")
				case i%3 != 0 && i%5 == 0:
					fmt.Println("Buzz")
				default:
					fmt.Println(i)
		}
	}
}