package file

import "testing"

func TestGetFileNameWithoutExtension(t *testing.T) {
	testCases := []struct {
		fileName       string
		expectedResult string
	}{
		{"./file.txt", "./file"},
		{"example.docx", "example"},
		{"data.csv", "data"},
		{"picture.jpg", "picture"},
		{"README", "README"},
	}

	for _, testCase := range testCases {
		result := GetFileNameWithoutExtension(testCase.fileName)
		if result != testCase.expectedResult {
			t.Errorf("For '%s', expected '%s', but got '%s'", testCase.fileName, testCase.expectedResult, result)
		}
	}
}
