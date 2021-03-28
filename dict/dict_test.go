package main

import "testing"

func assertDef(t testing.TB, dict Dict, word, def string)  {
	t.Helper()
	want := def
	got, err := dict.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSearch(t *testing.T) {

	d := Dict{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		assertDef( t, d, "test", "this is just a test")
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unk")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertStrings(err.Error(), ErrNotFound.Error(), t)
	})
}

func assertStrings(got string, want string, t testing.TB) {
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func TestAdd(t *testing.T) {
	d := Dict{}
	d.Add("test", "this is just a test")
	assertDef( t, d, "test", "this is just a test")
}