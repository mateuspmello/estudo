package main

import "fmt"

const helloEnglish = "Hello, "
const helloWorld = "World"

//Hello dá as boa vindas
func Hello(nome string) string {
	if nome == "" {
		nome = helloWorld
	}
	return helloEnglish + nome

}

func main() {
	fmt.Println(Hello("Mateus"))
}
