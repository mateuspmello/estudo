package interation

import "testing"

func TestImprimirCincoLetras(t *testing.T) {
	t.Run("Imprimir Cinco Letras", func(t *testing.T) {
		want := "aaaaa"
		got := Imprime("a")

		if want != got {
			t.Errorf("Esperava %s mas recebeu %s", want, got)
		}
	})
}
