package main

func HelloYou(name string) string {
	return "Hello, " + name
}

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "

func HelloYou2(language string, name string) string {
	if name == "" {
		name = "World"
	}
	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	println(HelloYou("Yuri"))
	println(HelloYou2("", "Yuri"))
}
