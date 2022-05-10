package main

import "fmt"

const englishHelloPrefix = "Hello, "
const germanyHelloPrefix = "Hallo, "

func hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	if language == "German" {
		return germanyHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Nosa", "Germany"))
}