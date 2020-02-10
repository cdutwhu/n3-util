package preprocess

import (
	"os"
	"os/exec"
	"path/filepath"

	cmn "github.com/cdutwhu/json-util/common"
)

func prepareJQ(jqDirs ...string) (jqWD, oriWD string, err error) {
	oriWD, err = os.Getwd()
	cmn.FailOnErr("%v", err)
	jqDirs = append(jqDirs, "./", "../", "../../")
	for _, jqWD = range jqDirs {
		if !sHasSuffix(jqWD, "/") {
			jqWD += "/"
		}
		if _, err = os.Stat(jqWD + jq); err == nil {
			cmn.FailOnErr("%v", os.Chdir(jqWD))
			jqWD, err = os.Getwd()
			cmn.FailOnErr("%v", err)
			return jqWD, oriWD, nil
		}
	}
	cmn.FailOnErr("Fatal: %v", fEf("%s is not found", jq))
	return "", "", nil
}

// FmtJSONStr : <json string> must not have single quote <'>
// May have 'Argument list too long' issue !
func FmtJSONStr(json string, jqDirs ...string) string {
	_, oriWD, err := prepareJQ(jqDirs...)
	cmn.FailOnErr("prepareJQ error @ %v", err)
	defer func() { os.Chdir(oriWD) }()

	json = "'" + sReplaceAll(json, "'", "\\'") + "'" // *** deal with <single quote> in "echo" ***
	// cmdstr := "echo " + json + `> temp.txt`       // May have 'Argument list too long' issue !
	cmdstr := "echo " + json + ` | ./` + jq + " ."
	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

	output, err := cmd.Output()
	cmn.FailOnErr("cmd.Output() error @ %v", err)
	return string(output)
}

// FmtJSONFile : <file> is the <Relative Path> to <LOCAL EXECUTABLE>, NOT to <JQ>
func FmtJSONFile(file string, jqDirs ...string) string {
	_, err := os.Stat(file)
	cmn.FailOnErr("file cannot be loaded @ %v", err)
	if !filepath.IsAbs(file) {
		absfile, err := filepath.Abs(file)
		cmn.FailOnErr("file path for absolute error @ %v", err)
		file = absfile
	}

	_, oriWD, err := prepareJQ(jqDirs...)
	cmn.FailOnErr("prepareJQ error @ %v", err)
	defer func() { os.Chdir(oriWD) }()

	cmdstr := "cat " + file + ` | ./` + jq + " ."
	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

	output, err := cmd.Output()
	cmn.FailOnErr("cmd.Output() error @ %v", err)
	return string(output)
}

// FmtJSONFile : <file> is the <relative path> to <jq>
// func FmtJSONFile(file string, jqDirs ...string) string {
// 	_, oriWD, err := prepareJQ(jqDirs...)
// 	cmn.FailOnErr("prepareJQ error @ %v", err)
// 	_, err = os.Stat(file)
// 	cmn.FailOnErr("file cannot be loaded @ %v", err)
// 	defer func() { os.Chdir(oriWD) }()

// 	if !filepath.IsAbs(file) {
// 		absfile, err := filepath.Abs(file)
// 		cmn.FailOnErr("file path for absolute error @ %v", err)
// 		file = absfile
// 	}

// 	cmdstr := "cat " + file + ` | ./` + jq + " ."
// 	cmd := exec.Command(execCmdName, execCmdP0, cmdstr)

// 	output, err := cmd.Output()
// 	cmn.FailOnErr("cmd.Output() error @ %v", err)
// 	return string(output)
// }
