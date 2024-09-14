package ebird

import (
	"testing"
)

func TestProcessOptions(t *testing.T) {
	t.Run("Back", func(t *testing.T) {
		options := Back(15)
		requestOptions := processOptions(options)

		if requestOptions.URLParams.Get("back") != "15" {
			t.Error("Back option not set correctly")
		}
	})

	t.Run("Cat", func(t *testing.T) {
		options := Cat("bird")
		requestOptions := processOptions(options)

		if requestOptions.URLParams.Get("cat") != "bird" {
			t.Error("Cat option not set correctly")
		}
	})

	t.Run("Multiple Options", func(t *testing.T) {
		options := []RequestOption{Back(10), Cat("sparrow"), Dist(50)}
		requestOptions := processOptions(options...)

		if requestOptions.URLParams.Get("back") != "10" ||
			requestOptions.URLParams.Get("cat") != "sparrow" ||
			requestOptions.URLParams.Get("dist") != "50" {
			t.Error("Multiple options not set correctly")
		}
	})
}
