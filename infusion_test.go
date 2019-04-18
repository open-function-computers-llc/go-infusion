package infusion

import (
	"testing"
)

// TestCanInitializeThePackage validate a Config struct and verify package initialization
func TestCanInitializeThePackage(t *testing.T) {
	// Setting up a new config
	c := Config{}

	// Test to make sure config has valid properties
	err := Init(c)
	if err != nil {
		t.Errorf(err.Error())
	}
}
