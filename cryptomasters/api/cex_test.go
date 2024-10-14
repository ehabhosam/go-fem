package api_test

import (
	"cryptomasters-project/api"
	"testing"
)

// simple unit test
func TestAPIGetRate(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected an error")
	}
}
