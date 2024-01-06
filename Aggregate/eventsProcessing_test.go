package aggregate

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAggregateEvents(t *testing.T) {
	// Create a temporary input file for testing
	tmpInputFile, err := CreateTempFile("test_input.json", TestInputJSON)
	if err != nil {
		t.Fatal("Error creating temporary input file:", err)
	}
	defer os.Remove(tmpInputFile)

	// Create a temporary output file for testing
	tmpOutputFile, err := CreateTempFile("test_output.json", "")
	if err != nil {
		t.Fatal("Error creating temporary output file:", err)
	}
	defer os.Remove(tmpOutputFile)

	// Run the AggregateEvents function with the temporary files
	AggregateEvents(tmpInputFile, tmpOutputFile)

	// Read and compare the actual output with the expected output
	actualOutput, err := os.ReadFile(tmpOutputFile)
	if err != nil {
		t.Fatal("Error reading actual output file:", err)
	}

	expectedOutput, err := json.MarshalIndent(TestExpectedOutput, "", "  ")
	if err != nil {
		t.Fatal("Error encoding expected output to JSON:", err)
	}

	if string(actualOutput) != string(expectedOutput) {
		t.Errorf("Actual output does not match expected output:\nActual:\n%s\nExpected:\n%s", actualOutput, expectedOutput)
	}
}

// Helper function to create a temporary file with the given content
func CreateTempFile(fileName, content string) (string, error) {
	tmpFile, err := os.CreateTemp("", fileName)
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

// Test data
var TestInputJSON = `[
	{"userId": 1, "eventType": "post", "timestamp": 1672444800},
	{"userId": 1, "eventType": "likeReceived", "timestamp": 1672444801},
	{"userId": 2, "eventType": "comment", "timestamp": 1672531201}
]`

var TestExpectedOutput = []map[string]interface{}{
	{"userId": 1, "date": "2022-12-31", "post": 1, "likeReceived": 1},
	{"userId": 2, "date": "2023-01-01", "comment": 1},
}
