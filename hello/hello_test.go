package hello

import "testing"

func TestHello(t *testing.T) {

	checkGotWant := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		checkGotWant(t, got, want)
		//assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jace", "")
		want := "Hello, Jace"
		checkGotWant(t, got, want)
		//assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		checkGotWant(t, got, want)
		//assertCorrectMessage(t, got, want)
	})
}

// Used the below function as a local variable in the above function
/*
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/
