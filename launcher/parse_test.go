package launcher

import (
	"bufio"
	"os"
	"testing"
)

func TestParseLaunchArgs_NoArgs(t *testing.T) {
	args, err := ParseLaunchArgs([]string{"cm-editor"}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{}, args)
}

func TestParseLaunchArgs_OnlyUserArgs(t *testing.T) {
	args, err := ParseLaunchArgs([]string{"cm-editor", "-iwad", "doom2", "-file", "sunlust.wad"}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{"-iwad", "doom2", "-file", "sunlust.wad"}, args)
}

func TestParseLaunchArgs_OnlyLoadArgs(t *testing.T) {
	cmanFile := testCmanFile(
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "load", cmanFile}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{
		"+cman_x0 0.0",
		"+cman_x1 0.1",
		"+cman_x2 0.2",
		"+pukename", "Test_Script",
	}, args)
}

func TestParseLaunchArgs_BothUserAndLoadArgs(t *testing.T) {
	cmanFile := testCmanFile(
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "load", cmanFile, "-iwad", "doom2", "-file", "sunlust.wad"}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{
		"-iwad", "doom2", "-file", "sunlust.wad",
		"+cman_x0 0.0",
		"+cman_x1 0.1",
		"+cman_x2 0.2",
		"+pukename", "Test_Script",
	}, args)
}

func TestParseLaunchArgs_MisplacedLoadCommand(t *testing.T) {
	cmanFile := testCmanFile(
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "-iwad", "doom2", "load", cmanFile, "-file", "sunlust.wad"}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{"-iwad", "doom2", "load", cmanFile, "-file", "sunlust.wad"}, args)
}

func TestParseLaunchArgs_MissingLoadPath(t *testing.T) {
	_, err := ParseLaunchArgs([]string{"cm-editor", "load"}, "Test_Script")
	assertErrorMessage(t, "missing filename after 'load' command", err)
}

func TestParseLaunchArgs_EveryAllowedCvar(t *testing.T) {
	cmanFile := testCmanFile(
		"path_mode = 0",
		"speed_mode = 1",
		"angle_mode = 2",
		"delay = 500",
		"speed = 1000",
		"overshoot = 0",
		"warp_player = 1",
		"hide_player = 1",
		"ga_buffer_len = 128",
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"y0 = 1.0",
		"y1 = 1.1",
		"y2 = 1.2",
		"z0 = 2.0",
		"z1 = 2.1",
		"z2 = 2.2",
		"a0 = 3.0",
		"a1 = 3.1",
		"p0 = 4.0",
		"p1 = 4.1",
		"ra0 = 5.0",
		"ra1 = 5.1",
		"cx0 = 6.0",
		"cx1 = 6.1",
		"cy0 = 7.0",
		"cy1 = 7.1",
		"r0 = 8.0",
		"r1 = 8.1",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "load", cmanFile}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{
		"+cman_path_mode 0",
		"+cman_speed_mode 1",
		"+cman_angle_mode 2",
		"+cman_delay 500",
		"+cman_speed 1000",
		"+cman_overshoot 0",
		"+cman_warp_player 1",
		"+cman_hide_player 1",
		"+cman_ga_buffer_len 128",
		"+cman_x0 0.0",
		"+cman_x1 0.1",
		"+cman_x2 0.2",
		"+cman_y0 1.0",
		"+cman_y1 1.1",
		"+cman_y2 1.2",
		"+cman_z0 2.0",
		"+cman_z1 2.1",
		"+cman_z2 2.2",
		"+cman_a0 3.0",
		"+cman_a1 3.1",
		"+cman_p0 4.0",
		"+cman_p1 4.1",
		"+cman_ra0 5.0",
		"+cman_ra1 5.1",
		"+cman_cx0 6.0",
		"+cman_cx1 6.1",
		"+cman_cy0 7.0",
		"+cman_cy1 7.1",
		"+cman_r0 8.0",
		"+cman_r1 8.1",
		"+pukename", "Test_Script",
	}, args)
}

func TestParseLaunchArgs_SkipUnallowedCvars(t *testing.T) {
	cmanFile := testCmanFile(
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"sv_cheats = 1",
		"xxx = 100",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "load", cmanFile}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{
		"+cman_x0 0.0",
		"+cman_x1 0.1",
		"+cman_x2 0.2",
		"+pukename", "Test_Script",
	}, args)
}

func TestParseLaunchArgs_SkipNonCvars(t *testing.T) {
	cmanFile := testCmanFile(
		"# comment",
		"x0 = 0.0",
		"x1 = 0.1",
		"x2 = 0.2",
		"// another comment",
		"some extra line that doesn't conform",
	)

	args, err := ParseLaunchArgs([]string{"cm-editor", "load", cmanFile}, "Test_Script")
	assertNoError(t, err)

	assertArrayEquals(t, []string{
		"+cman_x0 0.0",
		"+cman_x1 0.1",
		"+cman_x2 0.2",
		"+pukename", "Test_Script",
	}, args)
}

func testCmanFile(lines ...string) string {
	file, _ := os.CreateTemp(os.TempDir(), "*.cman")
	defer func() {
		_ = file.Close()
	}()

	writer := bufio.NewWriter(file)
	defer func() {
		_ = writer.Flush()
	}()

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	return file.Name()
}
