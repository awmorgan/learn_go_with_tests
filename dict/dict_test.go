package main

import "testing"

func assertDef(t testing.TB, dict Dict, word, def string) {
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
		assertDef(t, d, "test", "this is just a test")
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
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dict{}
		w := "test"
		def := "this is just a test"
		err := d.Add(w, def)
		assertErr(t, err, nil)
		assertDef(t, d, w, def)
	})
	t.Run("existing word", func(t *testing.T) {
		w := "test"
		def := "this is just a test"
		d := Dict{w:def}
		err := d.Add(w, def)
		assertErr(t, err, ErrWordExists)
		assertDef(t, d, w, def)
	})
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
