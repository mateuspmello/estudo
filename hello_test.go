package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Testando saida do hello", func(t *testing.T) {
		got := Hello("Mateus")
		want := "Hello, Mateus"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
