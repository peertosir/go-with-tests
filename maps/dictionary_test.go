package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("test")
		want := ErrNoItemInDict
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add value not exists in dict", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("test", "this is just a test")
		assertNoError(t, err)

		want := "this is just a test"
		got, err := dict.Search("test")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
	t.Run("cannot add value exists in dict", func(t *testing.T) {
		dict := Dictionary{"test": "not rewritten"}
		err := dict.Add("test", "this is just a test")

		want := "not rewritten"
		got, _ := dict.Search("test")

		assertError(t, err, ErrItemAlreadyInDict)
		assertStrings(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update value not exists in dict", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update("test", "updated")

		assertError(t, err, ErrNoItemInDict)
	})
	t.Run("update value exists in dict", func(t *testing.T) {
		dict := Dictionary{"test": "initial"}
		err := dict.Update("test", "updated")
		assertNoError(t, err)

		want := "updated"
		got, _ := dict.Search("test")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		dict := Dictionary{"test": "to delete"}

		err := dict.Delete("test")
		assertNoError(t, err)

		_, err = dict.Search("test")
		assertError(t, err, ErrNoItemInDict)
	})
	t.Run("delete non existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("test")
		assertError(t, err, ErrNoItemInDict)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', but want '%s'", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expect an error, but didnt get one")
	}

	if got != want {
		t.Errorf("want %q error, but got %q", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error, but didnt want one")
	}
}
