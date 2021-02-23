package pkg_test

import (
	"fmt"
	"testing"

	"github.com/mfroes/gotest/main/pkg"
)

func TestMetadata_toJSON(t *testing.T) {
	expected := `{"version":"99.99","description":"teST dEscripTIon","lastcommitsha":"RANDOMSHA"}`
	metadata := pkg.Metadata{
		"99.99", "teST dEscripTIon", "RANDOMSHA",
	}
	stringMetadata := metadata.ToJSON()

	if stringMetadata != expected {
		t.Error(
			fmt.Printf("Incorrect JSON received:\n expected[ %v ],\n got[ %v ]", expected, stringMetadata),
		)
		t.FailNow()
	}
}
