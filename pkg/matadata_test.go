package pkg_test

import (
	"fmt"
	"testing"

	"github.com/mfroes/gotest/main/pkg"
)

func TestMetadata_toJSON(t *testing.T) {
	expected := `{"version":"99.99","description":"teST dEscripTIon","lastcommitsha":"abc57858585"}`
	pkg.BuildDescription = "teST dEscripTIon"
	pkg.BuildVersion = "99.99"
	pkg.BuildSHA = "abc57858585"
	pkg.BuildTime = "today"

	metadata := pkg.GetMetadata()
	stringMetadata := metadata.ToJSON()

	if stringMetadata != expected {
		t.Errorf("Incorrect JSON received:\n got [%v]\n want [%v]",
			stringMetadata, expected)
	}
}
