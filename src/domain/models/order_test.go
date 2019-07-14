package models

import (
	"flag"
	"testing"
)

var integration = flag.Bool("integration", false, "run database integration tests")

func TestLatLng(t *testing.T) {
	err := validateLatLng("", "")

	if err == nil {
		t.Error("error should be thrown for wrong lat and lng")
	}

	err = validateLatLng("2", "2")

	if err != nil {
		t.Error(err)
	}

	err = validateLatLng("100", "190")

	if err == nil {
		t.Error("invalid coordinates")
	}

}
