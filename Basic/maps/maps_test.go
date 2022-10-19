package maps

import "testing"

func assertSearchRightDictValue(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expect %q, while got %q", want, got)
	}
}

func assertDefinition(t *testing.T, dictionary Dict, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}
	got, _ := dict.Search("test")
	want := "this is just a test"
	assertSearchRightDictValue(t, got, want)
}

func TestTableDrivenSeach(t *testing.T) {
	dict := Dict{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertSearchRightDictValue(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		if err.Error() == "" {
			t.Errorf("Error not occurred when search unkown key in map")
		}
	})
}

func TestAdd(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("expect %q, while got %q", want, got)
	}

}

func TestUpdate(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dict := Dict{key: value}
	newDefinition := "new definition"

	dict.Update(key, newDefinition)
	assertDefinition(t, dict, key, newDefinition)
}

func TestWrongUpdate(t *testing.T) {
	key := "test"
	value := "this is just a test"
	dict := Dict{}
	err := dict.Update(key, value)

	assertError(t, err, ErrKeyDoesNotExistMapValue)
}

func TestDelete(t *testing.T) {
	key := "test"
	dict := Dict{key: "value for test"}
	dict.Delete(key)

	_, err := dict.Search(key)
	if err != ErrNotFoundMapValue {
		t.Errorf("Expected %q to be deleted", key)
	}
}
