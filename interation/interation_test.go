package interation

import "testing"

func TestImprimirCincoLetras(t *testing.T) {
	t.Run("Imprimir Cinco Letras", func(t *testing.T) {
		got := Imprime("a", 3)
		want := "aaa"

		if want != got {
			t.Errorf("Esperava %s mas recebeu %s", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Imprime("a", 6)
	}
}
