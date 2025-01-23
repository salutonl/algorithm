package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello("Chris", "Chinese")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func (t *testing.T){
		got := Hello("Tom", "Spanish")
		want := "Hola, Tom!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func (t *testing.T) {
		got := Hello("Fre", "French")
		want := "Bonjour, Fre!"
		
		assertCorrectMessage(t, got, want)
		
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}