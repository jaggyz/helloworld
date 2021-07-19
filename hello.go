package hello

import "rsc.io/quote"

func Hello1() string {
	return "Hello, world."
}

func Hello2() string {
	return quote.Hello()
}
