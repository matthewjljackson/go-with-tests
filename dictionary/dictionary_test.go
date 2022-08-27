package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrMissingKey.Error()

		if err == nil {
			t.Fatal("expected an error but got none")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	key := "test"
	value := "this is just a test"

	t.Run("New key value pair", func(t *testing.T) {
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("Already existing key value pair", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		newValue := "testy"
		err := dictionary.Add(key, newValue)

		assertError(t, err, ErrExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err1 := dictionary.Delete(word)

		assertError(t, err1, nil)

		_, err := dictionary.Search(word)
		if err != ErrMissingKey {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

	t.Run("No word to delete", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("test")

		if err != ErrMissingKey {
			t.Error("Expected missing key error")
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != value {
		t.Errorf("got %q want %q", got, value)
	}
}
