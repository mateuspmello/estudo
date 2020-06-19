package main

import "fmt"

const portugues = "Portugues"
const italiano = "Italiano"

const helloEnglish = "Hello, "
const olaPortugues = "Olá, "
const ciaoItaliano = "Ciao, "
const helloWorld = "World"

//Hello dá as boa vindas
func Hello(nome string, lingua string) string {
	if nome == "" {
		nome = helloWorld
	}
	return prefixo(lingua) + nome
}

func prefixo(lingua string) (prefixo string) {
	switch lingua {
	case portugues:
		prefixo = olaPortugues
	case italiano:
		prefixo = ciaoItaliano
	default:
		prefixo = helloEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Mateus", ""))
}
