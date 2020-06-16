package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	mensagemDeErro := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testando saida do hello", func(t *testing.T) {
		got := Hello("Mateus")
		want := "Hello, Mateus"

		mensagemDeErro(t, got, want)
	})

	t.Run("Testando com nome vazio", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		mensagemDeErro(t, got, want)
	})
}
