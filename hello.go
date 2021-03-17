package hello

import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

func Grass() string {
	return "I can eat grass and it doesn't hurt me." 
}
