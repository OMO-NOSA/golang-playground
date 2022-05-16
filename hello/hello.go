package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const germanyHelloPrefix = "Hallo, "
const spanishHelloPrefix = "Hallo, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch language{
	case "german":
		prefix = germanyHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}



func main() {
	fmt.Println(Hello("Nosa", "Germany"))
}