package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	//	return englishHelloPrefix + name

	if name == "" {
		name = "Sakura Endo"
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world"))
}
