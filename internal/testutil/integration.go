package testutil

import (
	"os"
	"testing"
)

// IntegrationTest can be called at the start of a particular test to mark it
// as an integration test. Integration tests are are skipped unless the
// environment variable INTEGRATION is set to true.
func IntegrationTest(t *testing.T) {
	t.Helper()
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integration test...")
	}
}
