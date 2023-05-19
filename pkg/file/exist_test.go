package file

import "testing"

func TestFileExists(t *testing.T) {
	// Existing file
	existingFilePath := "./testdata/existing_file.txt"
	exists := Exists(existingFilePath)
	if !exists {
		t.Errorf("Expected file to exist, but it doesn't.")
	}

	// Non-existent file
	nonexistentFilePath := "./testdata/nonexistent_file.txt"
	exists = Exists(nonexistentFilePath)
	if exists {
		t.Errorf("Expected file to not exist, but it does.")
	}
}
