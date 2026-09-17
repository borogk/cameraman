package launcher

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"
)

func TestScanAndSaveExportedProfiles_EmptyLog(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog()

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{}, outputDir)
}

func TestScanAndSaveExportedProfiles_NonEmptyLogWithNoExports(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"Goodbye",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{}, outputDir)
}

func TestScanAndSaveExportedProfiles_EmptyExport(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"--- BEGIN CAMERAMAN ---",
		"--- END CAMERAMAN ---",
		"Goodbye",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{"export-0001.cman"}, outputDir)
	assertFileContent(t, []string{}, path.Join(outputDir, "export-0001.cman"))
}

func TestScanAndSaveExportedProfiles_SingleExport(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"--- END CAMERAMAN ---",
		"Goodbye",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{"export-0001.cman"}, outputDir)
	assertFileContent(t, []string{
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	}, path.Join(outputDir, "export-0001.cman"))
}

func TestScanAndSaveExportedProfiles_MultipleExports(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"--- END CAMERAMAN ---",
		"MAP 02 - Underhalls",
		"Picked up a medikit",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 1.0",
		"x1 = 1.1",
		"x2 = 1.2",
		"--- END CAMERAMAN ---",
		"Goodbye",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{
		"export-0001.cman",
		"export-0002.cman",
	}, outputDir)
	assertFileContent(t, []string{
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	}, path.Join(outputDir, "export-0001.cman"))
	assertFileContent(t, []string{
		"x0 = 1.0",
		"x1 = 1.1",
		"x2 = 1.2",
	}, path.Join(outputDir, "export-0002.cman"))
}

func TestScanAndSaveExportedProfiles_OverlappingExports(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"x1 = 0.1",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 1.0",
		"x1 = 1.1",
		"x2 = 1.2",
		"--- END CAMERAMAN ---",
		"x2 = 0.2",
		"--- END CAMERAMAN ---",
		"MAP 02 - Underhalls",
		"Picked up a medikit",
		"Goodbye",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertErrorMessage(t, "unexpected profile begin before the previous export has ended", err)

	// Unfinished profile is expected to still be saved
	assertDirFiles(t, []string{"export-0001.cman"}, outputDir)
	assertFileContent(t, []string{
		"x0 = 0.0",
		"x1 = 0.1",
	}, path.Join(outputDir, "export-0001.cman"))
}

func TestScanAndSaveExportedProfiles_LastProfileInterrupted(t *testing.T) {
	outputDir := testOutputDir()
	log := testLog(
		"ZDoom version 6.66",
		"MAP 01 - Entryway",
		"You suck!",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"--- END CAMERAMAN ---",
		"MAP 02 - Underhalls",
		"Picked up a medikit",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 1.0",
		"x1 = 1.1",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertErrorMessage(t, "last profile export was interrupted", err)

	// Both finished and interrupted profiles are expected to still be saved
	assertDirFiles(t, []string{
		"export-0001.cman",
		"export-0002.cman",
	}, outputDir)
	assertFileContent(t, []string{
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	}, path.Join(outputDir, "export-0001.cman"))
	assertFileContent(t, []string{
		"x0 = 1.0",
		"x1 = 1.1",
	}, path.Join(outputDir, "export-0002.cman"))
}

func TestScanAndSaveExportedProfiles_GenerateNewNameWhenSomeAlreadyExist(t *testing.T) {
	outputDir := testOutputDir()
	_, _ = os.Create(path.Join(outputDir, "export-0001.cman"))
	_, _ = os.Create(path.Join(outputDir, "export-0002.cman"))
	log := testLog(
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"--- END CAMERAMAN ---",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{
		"export-0001.cman",
		"export-0002.cman",
		"export-0003.cman",
	}, outputDir)
}

func TestScanAndSaveExportedProfiles_GenerateNewNameSkippingNumberGaps(t *testing.T) {
	outputDir := testOutputDir()
	_, _ = os.Create(path.Join(outputDir, "export-0001.cman"))
	_, _ = os.Create(path.Join(outputDir, "export-0002.cman"))
	_, _ = os.Create(path.Join(outputDir, "export-0004.cman"))
	_, _ = os.Create(path.Join(outputDir, "export-0006.cman"))
	log := testLog(
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"--- END CAMERAMAN ---",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, []string{
		"export-0001.cman",
		"export-0002.cman",
		"export-0004.cman",
		"export-0006.cman",
		"export-0007.cman",
	}, outputDir)
}

func TestScanAndSaveExportedProfiles_GenerateNewNameOverflow(t *testing.T) {
	outputDir := testOutputDir()

	expectedFileNames := make([]string, 0)
	for i := 1; i <= 9999; i++ {
		name := fmt.Sprintf("export-%04d.cman", i)
		_, _ = os.Create(path.Join(outputDir, name))
		expectedFileNames = append(expectedFileNames, name)
	}
	expectedFileNames = append(expectedFileNames, "export-10000.cman", "export-10001.cman")

	log := testLog(
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"--- END CAMERAMAN ---",
		"--- BEGIN CAMERAMAN ---",
		"x0 = 0.0",
		"--- END CAMERAMAN ---",
	)

	err := ScanAndSaveExportedProfiles(log, outputDir)
	assertNoError(t, err)

	assertDirFiles(t, expectedFileNames, outputDir)
}

func testOutputDir() string {
	temp, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		panic(err)
	}
	return temp
}

func testLog(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}
