package funcs

import (
	"os"
	"output/funcs"
	"strings"
	"testing"
)

// Test for character A
func TestPrintAsciiArt_A(t *testing.T) {
	testPrintAsciiArtToFile(t, []string{"A"})
}

// Test for character B
func TestPrintAsciiArt_B(t *testing.T) {
	testPrintAsciiArtToFile(t, []string{".!/?  /n/npaok"})
}

// Test for character C
func TestPrintAsciiArt_C(t *testing.T) {
	testPrintAsciiArtToFile(t, []string{"3fd:][2]"})
}

// Test for multiple characters
func TestPrintAsciiArt_Hello(t *testing.T) {
	testPrintAsciiArtToFile(t, []string{"Hello"})
}

// Helper function to test PrintAsciiArt with input and expected output
func testPrintAsciiArtToFile(t *testing.T, input []string) {
	styleBanner := "standard"
	bannerFileName := "../banners/" + styleBanner + ".txt"
	// Read the banner file to obtain the ASCII art characters
	fileContent, err := os.ReadFile(bannerFileName)
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}
	fileContentString := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
	bannerLines := strings.Split(fileContentString, "\n")
	// Generate expected output
	expectedOutput := generateExpectedOutput(bannerLines, input)
	// Create a temporary file to capture the output
	tempFile, err := os.CreateTemp("", "ascii_output_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure cleanup after test
	// Use PrintToFile to write output to temp file
	funcs.PrintToFile(tempFile, input, bannerLines)
	// Close the temp file so we can read from it
	tempFile.Close()
	// Read actual output from the temp file
	actualOutput, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read from temporary file: %v", err)
	}
	// Compare expected and actual outputs
	if string(actualOutput) != expectedOutput {
		t.Errorf("For input %v, expected:\n%s\nGot:\n%s", input, expectedOutput, actualOutput)
	}
}

// Helper function to generate expected output based on input
func generateExpectedOutput(bannerLines []string, input []string) string {
	var output strings.Builder
	for h := 1; h < 9; h++ {
		for _, word := range input {
			for i := 0; i < len(word); i++ {
				charIndex := int(word[i]) - 32
				lineIndex := charIndex*9 + h
				output.WriteString(bannerLines[lineIndex])
			}
		}
		output.WriteString("\n")
	}
	return output.String()
}
