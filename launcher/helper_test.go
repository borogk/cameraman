package launcher

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func assertArrayEquals(t *testing.T, expected []string, actual []string) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("length mismatch: got %d, want %d", len(actual), len(expected))
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %s, want %s", i, actual[i], expected[i])
		}
	}
}

func assertDirFiles(t *testing.T, expectedNames []string, actualDir string) {
	t.Helper()
	entries, err := os.ReadDir(actualDir)
	assertNoError(t, err)

	actualNames := make([]string, len(entries))
	for i, entry := range entries {
		actualNames[i] = filepath.Base(entry.Name())
	}

	sort.Strings(expectedNames)
	sort.Strings(actualNames)
	assertArrayEquals(t, expectedNames, actualNames)
}

func assertFileContent(t *testing.T, expectedLines []string, actualFile string) {
	t.Helper()
	file, err := os.Open(actualFile)
	assertNoError(t, err)

	defer func() {
		_ = file.Close()
	}()

	actualLines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		actualLines = append(actualLines, scanner.Text())
	}

	assertArrayEquals(t, expectedLines, actualLines)
}

func assertNoError(t *testing.T, actualError error) {
	t.Helper()
	if actualError != nil {
		t.Fatal(actualError)
	}
}

func assertErrorMessage(t *testing.T, expectedMessage string, actualError error) {
	t.Helper()
	if actualError == nil {
		t.Fatal("expected error, got none")
	}
	if actualError.Error() != expectedMessage {
		t.Fatalf("unexpected error: %s", actualError)
	}
}
