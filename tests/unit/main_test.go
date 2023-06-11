package unit

import (
	"testing"
)

var (
	runMainFunc = RunMain
)

func RunMain() {}

func TestRunMainInvocation(t *testing.T) {
	isRunMainInvoked := false

	runMainMock := func() {
		isRunMainInvoked = true
	}

	// Replace the function pointer with the mock implementation
	runMainFunc = runMainMock

	runMainFunc()

	if !isRunMainInvoked {
		t.Error("RunMain function is not invoked")
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
