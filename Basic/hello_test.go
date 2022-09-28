package main

import "testing"

func TestName(t *testing.T) {
	got := Hello("David", "")
	want := "Hello, David"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestEmptyHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
}

func gotAssertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMultiLanguageHello(t *testing.T) {

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Halo, Elodie"

		gotAssertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		gotAssertCorrectMessage(t, got, want)
	})
}
