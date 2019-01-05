package decimal_test

import (
	"path/filepath"
	"testing"

	"github.com/ericlagergren/decimal/dectest"
)

func TestDecTests(t *testing.T) {
	files, err := filepath.Glob("_dectest/*.decTest")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		file := file // shadow range variable
		t.Run(filepath.Base(file), func(t *testing.T) {
			dectest.Test(t, file)
		})
	}
}
