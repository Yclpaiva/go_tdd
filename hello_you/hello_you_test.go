package main

import "testing"

func TestHelloYou(t *testing.T) {
	got := HelloYou("Yuri")
	want := "Hello, Yuri"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloYour2(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		//Test1
		got := HelloYou2("", "Yuri")
		want := "Hello, Yuri"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		//Test2
		got := HelloYou2("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hola, ' spanish prefix", func(t *testing.T) {
		//Test3
		got := HelloYou2("Spanish", "Yuri")
		want := "Hola, Yuri"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Bonjour, ' french prefix", func(t *testing.T) {
		//Test4
		got := HelloYou2("French", "Yuri")
		want := "Bonjour, Yuri"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Olá, ' portuguese prefix", func(t *testing.T) {
		//Test5
		got := HelloYou2("Portuguese", "Yuri")
		want := "Olá, Yuri"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
