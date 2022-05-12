package greeter

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying \"Hello, World!\" when name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello to Chris in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello to Chris in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris!"

		assertCorrectMessage(t, got, want)
	})
}
