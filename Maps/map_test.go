package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("detest")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dict.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("attempt to add an existing word", func(t *testing.T) {
		word := "test"
		def := "just a test, ma'am"
		dict := Dictionary{word: def}
		err := dict.Add(word, "new test, ma'am")
		assertError(t, err, WordExistsError)
		assertDefinition(t, dict, word, def)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "just a test, ma'am"
		dict := Dictionary{word: def}
		newDef := "new test, ma'am"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "just a test, ma'am"

		err := dict.Update(word, def)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "just a test, ma'am"}
	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' wanted '%s' given '%s'", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s', want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, key, def string) {
	t.Helper()

	got, err := dict.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if def != got {
		t.Errorf("got '%s', want '%s'", got, def)
	}
}
