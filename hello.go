package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, bahasa string) string {
	//	return englishHelloPrefix + name

	if name == "" && bahasa == "" {
		name = "Sakura Endo"
		bahasa = "jepang"
	}

	return englishHelloPrefix + name + bahasa
}
func main() {
	fmt.Println(Hello("world", "jepang"))
}
