package urunce2etesting

import (
	"os"
	"testing"
)

func TestLoadTestCases(t *testing.T) {
	// Create a temporary config file
	configContent := `{
		"dummyTool": [
			{
				"Name": "Test1",
				"TestFuncName": "pingTest",
				"Image": "busybox"
			},
			{
				"Name": "Test2",
				"TestFuncName": "", 
				"Image": "busybox"
			}
		]
	}`
	tmpfile, err := os.CreateTemp("", "config_test_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(configContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Case 1: Success
	tests, err := loadTestCases(tmpfile.Name(), "dummyTool")
	if err != nil {
		t.Fatalf("Expected success, got error: %v", err)
	}
	if len(tests) != 2 {
		t.Errorf("Expected 2 tests, got %d", len(tests))
	}
	if tests[0].TestFunc == nil {
		t.Error("Expected TestFunc to be mapped for pingTest")
	}

	// Case 2: Tool not found
	_, err = loadTestCases(tmpfile.Name(), "nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent tool, got nil")
	}

	// Case 3: Invalid JSON (using a bad file path for simplicity of testing error path)
	_, err = loadTestCases("non_existent_file.json", "dummyTool")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}
