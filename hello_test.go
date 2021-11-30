package main

import "testing"

func TestHello(t *testing.T) {
	//	got := Hello("sakura endo")
	//	want:= "Hello, sakura endo"
	//
	//	if got != want {
	//		t.Errorf("got %v want %v" , got, want)
	//	}
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v , want %v", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sakura Endo", " jepang")
		want := "Hello, Sakura Endo"
		assertCorrectMessage(t, got, want)
	})

	// t.Run("saying hello to people", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello, Sakura Endo"
	// 	assertCorrectMessage(t, got, want)
	// })
}
