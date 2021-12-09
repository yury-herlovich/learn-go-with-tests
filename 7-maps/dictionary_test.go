package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word")
		}

		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "new definition")

		assertError(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.Update("test", "new definition")

		want := "new definition"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("expected to not get an error")
		}

		assertStrings(t, got, want)
	})

	t.Run("update not existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "new definition")

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.Delete("test")

		_, err := dictionary.Search("test")

		assertError(t, err, ErrNotFound)
	})

	t.Run("delete not existing key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error")
	}

	assertStrings(t, got.Error(), want.Error())

}
