package preprocess

import (
	"os"
	"os/exec"
	"path/filepath"

	eg "github.com/cdutwhu/json-util/n3errs"
)

func prepare(exe string, wkDirs ...string) (cWD, oriWD string, err error) {
	oriWD, err = os.Getwd()
	failOnErr("%v", err)
	wkDirs = append(wkDirs, "./", "../", "../../")
	for _, cWD = range wkDirs {
		cWD = sTrimRight(cWD, "/") + "/"
		if _, err = os.Stat(cWD + exe); err == nil {
			failOnErr("%v", os.Chdir(cWD))
			cWD, err = os.Getwd()
			failOnErr("%v", err)
			return cWD, oriWD, nil
		}
	}
	failOnErr("%v: %s", eg.FILE_NOT_FOUND, exe)
	return "", "", nil
}

// FmtJSONStr : May have 'Argument list too long' issue !
func FmtJSONStr(json string, jqDirs ...string) string {
	_, oriWD, err := prepare(jq, jqDirs...)
	failOnErr("prepare jq error @ %v", err)
	defer func() { os.Chdir(oriWD) }()

	json = sReplaceAll(json, `\`, `\\`)
	json = sReplaceAll(json, `'`, `\'`)
	json = `$'` + json + `'` // *** deal with <quotes> in "echo" ***
	// cmdstr := "echo " + json + `> temp.txt`          // May have 'Argument list too long' issue AND Apostrophe issue !
	cmdstr := "echo " + json + ` | ./` + jq + " ."
	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

	output, err := cmd.Output()
	failOnErr("cmd.Output() error @ %v", err)
	return string(output)
}

// FmtJSONFile : <file> is the <Relative Path> to <LOCAL EXECUTABLE>, NOT to <JQ>
func FmtJSONFile(file string, jqDirs ...string) string {
	_, err := os.Stat(file)
	failOnErr("file cannot be loaded @ %v", err)
	if !filepath.IsAbs(file) {
		absfile, err := filepath.Abs(file)
		failOnErr("file path for absolute error @ %v", err)
		file = absfile
	}

	_, oriWD, err := prepare(jq, jqDirs...)
	failOnErr("prepare jq error @ %v", err)
	defer func() { os.Chdir(oriWD) }()

	cmdstr := "cat " + file + ` | ./` + jq + " ."
	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

	output, err := cmd.Output()
	failOnErr("cmd.Output() error @ %v", err)
	return string(output)
}

// FmtJSONFile : <file> is the <relative path> to <jq>
// func FmtJSONFile(file string, jqDirs ...string) string {
// 	_, oriWD, err := prepare(jq, jqDirs...)
// 	failOnErr("prepare jq error @ %v", err)
// 	_, err = os.Stat(file)
// 	failOnErr("file cannot be loaded @ %v", err)
// 	defer func() { os.Chdir(oriWD) }()

// 	if !filepath.IsAbs(file) {
// 		absfile, err := filepath.Abs(file)
// 		failOnErr("file path for absolute error @ %v", err)
// 		file = absfile
// 	}

// 	cmdstr := "cat " + file + ` | ./` + jq + " ."
// 	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

// 	output, err := cmd.Output()
// 	failOnErr("cmd.Output() error @ %v", err)
// 	return string(output)
// }
