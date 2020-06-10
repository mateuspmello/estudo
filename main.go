package main

import "fmt"

//Hello dá as boa vindas
func Hello(nome string) string {
	return "Hello, " + nome
}

func main() {
	fmt.Println(Hello("Mateus"))
}
