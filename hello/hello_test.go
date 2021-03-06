package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t testing.TB, got, want string)  {
		t.Helper()
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to fine humans", func (t *testing.T)  {
		got := Hello("Nosa", " ")
		want := "Hello, Nosa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello world' when an empty string is supplied", func (t *testing.T)  {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func (t *testing.T)  {
		got := Hello("Nosa", "german")
		want:= "Hallo, Nosa"
		assertCorrectMessage(t, got, want)
	})

}