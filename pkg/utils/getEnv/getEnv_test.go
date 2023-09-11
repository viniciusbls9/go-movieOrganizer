package getEnv

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Test case 1: Environment variable is set
	expectedValue := "ENV_DB"
	os.Setenv("DB_URL", expectedValue)

	actualValue := GetEnv()
	if actualValue != expectedValue {
		t.Errorf("Expected: %s, Got: %s", expectedValue, actualValue)
	}
}
