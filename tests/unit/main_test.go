package unit

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/xehap/jago/runner"
)

func TestRunMain(t *testing.T) {
	expectedClientID := "testClientID"
	expectedClientSecret := "testClientSecret"

	// Redirect os.Stdin for testing
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	r, w, _ := os.Pipe()
	defer r.Close()
	defer w.Close()
	os.Stdin = r

	// Write the expected client ID and client secret to the pipe
	_, _ = w.WriteString(expectedClientID + "\n" + expectedClientSecret + "\n")
	_ = w.Close()

	// Capture the output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Run the main logic
	runner.RunMain()

	// Check the output
	actualOutput := buf.String()
	if !strings.Contains(actualOutput, expectedClientID) {
		t.Errorf("Expected output to contain client ID '%s', but got: %s", expectedClientID, actualOutput)
	}
	if !strings.Contains(actualOutput, expectedClientSecret) {
		t.Errorf("Expected output to contain client secret '%s', but got: %s", expectedClientSecret, actualOutput)
	}
}
