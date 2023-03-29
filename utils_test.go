package utils_test

import (
	"fmt"
	"testing"

	"github.com/joroovb/go-utils"
)

func TestHead(t *testing.T) {
	t.Run("opens a file and returns the top 2 lines", func(t *testing.T) {
		want := []string{"one", "two"}

		got, err := utils.Head("testdata/somefile", 2)
		if err != nil {
			fmt.Println(err)
		}

		if !Equal(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("opens a file and attempts to read more lines than are in the file", func(t *testing.T) {
		want := []string{"one", "two", "three", "four", "five"}

		got, err := utils.Head("testdata/somefile", 100)
		if err != nil {
			fmt.Println(err)
		}

		if !Equal(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("attempts to open unknown file and errors out", func(t *testing.T) {
		got, err := utils.Head("testdata/ghost", 2)

		if err == nil {
			t.Errorf("should return an error on missing file")
		}

		if got != nil {
			t.Errorf("should return nil lines on missing file")
		}
	})

}

func TestTail(t *testing.T) {
	t.Run("opens a file and returns the bottom 2 lines", func(t *testing.T) {
		want := []string{"four", "five"}

		got, err := utils.Tail("testdata/somefile", 2)
		if err != nil {
			fmt.Println(err)
		}

		if !Equal(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("opens a file and attempts to read more lines than are in the file", func(t *testing.T) {
		want := []string{"one", "two", "three", "four", "five"}

		got, err := utils.Tail("testdata/somefile", 100)
		if err != nil {
			fmt.Println(err)
		}

		if !Equal(got, want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("attempts to open unknown file and errors out", func(t *testing.T) {
		got, err := utils.Tail("testdata/ghost", 2)

		if err != nil {
			t.Errorf("should return an error on missing file")
		}

		if got != nil {
			t.Errorf("should return nil lines on missing file")
		}
	})
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
