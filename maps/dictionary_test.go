package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	actual, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find addeed word:", err)
	}

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}

}
