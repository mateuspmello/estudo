package interation

//Imprime imprime as 5 primeiras letras
func Imprime(a string, b int) string {
	var texto string
	for i := 0; i < b; i++ {
		texto += a
	}
	return texto
}
