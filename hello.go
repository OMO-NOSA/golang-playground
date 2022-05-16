package integers

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

func Add(x, y int) int {

	return x + y
}

func main() {
	fmt.Println(Hello("Nosa", "Germany"))
	fmt.Println(Add(55,6))
}