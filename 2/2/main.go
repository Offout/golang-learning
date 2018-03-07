package main

import (
	"fmt"
	"strings"
)

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	cache := make(map[int]interface{})
	return func(n int, other ...int) interface{} {
		if cache[n] == nil {
			cache[n] = function(n, other...)
		}
		return cache[n]
	}
}


// TODO обернуть функции fibonacci и roman в memoize
func init() {
	fibonacci = func(n int, other ...int) interface{} {
		if n == 1 || n == 2 {
			return 1
		}
		return fibonacci(n - 1).(int) + fibonacci(n - 2).(int)
	}

	fibonacci = memoize(fibonacci)

	romanForDecimal = func(n int, other ...int) interface{} {
		str := ""
		switch {
		case n > 1000:
			{
				str = strings.Repeat("M", int(n / 1000))
				n = n % 1000
			}
		case n == 1000:
			{
				str = "M"
				n = n - 1000
			}
		case n >= 900:
			{
				str = "CM"
				n = n - 900
			}
		case n > 500:
			{
				str = strings.Repeat("D", int(n / 500))
				n = n % 500
			}
		case n == 500:
			{
				str = "D"
				n = n - 500
			}
		case n >= 400:
			{
				str = "CD"
				n = n - 400
			}
		case n > 100:
			{
				str = strings.Repeat("C", int(n / 100))
				n = n % 100
			}
		case n == 100:
			{
				str = "C"
				n = n - 100
			}
		case n >= 90:
			{
				str = "XC"
				n = n - 90
			}
		case n >= 50:
			{
				str = "L"
				n = n % 50
			}
		case n >= 40:
			{
				str = "XL"
				n = n - 40
			}
		case n >= 10:
			{
				str = strings.Repeat("X", int(n / 10))
				n = n % 10
			}
		case n == 9:
			{
				str = "IX"
				n = 0
			}
		case n >= 5:
			{
				str = "V" + strings.Repeat("I", n-5)
				n = 0
			}
		case n == 4:
			{
				str = strings.Repeat("IV", 1)
				n = 0
			}
		case n <= 3:
			{
				str = strings.Repeat("I", n)
				n = 0
			}
		default:
			str = ".."
			n = 0
		}
		if n > 0 {
			return str + romanForDecimal(n, other...).(string)
		}
		return str
	}

	romanForDecimal = memoize(romanForDecimal)
}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}

