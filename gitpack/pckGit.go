package gitpack

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getFunctionName() string {
	x, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s()", runtime.FuncForPC(x).Name())
}

func printStrings(format string, values ...any) {
	log.Printf(format, values...)
}

func printInfoStrings(funcName string, funcArg string) {
	log.Printf("Function name: <%s> arguments: <%s>. Status successfully!", funcName, funcArg)
}

func printErrorStrings(funcName string, funcArg string, err error) {
	log.Printf("Function name: <%s>, arguments: <%s>, error: <%s>. Status Error!!!", funcName, funcArg, err)
}

func GitRemoveDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	err := os.RemoveAll(path)
	if err != nil {
		printErrorStrings(getFunctionName(), path, err)
	} else {
		printStrings("Function name: <%s> arguments: <%s>. Status successfully!", getFunctionName(), path)
	}
}

func GitCreateDir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		printErrorStrings(getFunctionName(), path, err)
	} else {
		printInfoStrings(getFunctionName(), path)
	}
}

func GitCreateFile(path string, fileName string, content string) {
	mainGoFilePath := filepath.Join(path, fileName)
	file, err := os.Create(mainGoFilePath)
	messStr := fmt.Sprintf("%s/%s", path, fileName)
	if err != nil {
		printErrorStrings(getFunctionName(), messStr, err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		printErrorStrings(getFunctionName(), messStr, err)
	}
	file.Close()
	printInfoStrings(getFunctionName(), messStr)
}

func GitChangeFile(path string, fileName string, content string) {
	updateGoFilePath := filepath.Join(path, fileName)
	updateFile, err := os.OpenFile(updateGoFilePath, os.O_RDWR, 0644)
	messStr := fmt.Sprintf("%s/%s", path, fileName)
	if err != nil {
		printErrorStrings(getFunctionName(), messStr, err)
	}
	_, err = updateFile.Seek(0, 0)
	if err != nil {
		printErrorStrings(getFunctionName(), messStr, err)
	}
	_, err = updateFile.WriteString(content)
	if err != nil {
		printErrorStrings(getFunctionName(), messStr, err)
	}
	updateFile.Close()
	printInfoStrings(getFunctionName(), messStr)
}

func GitCommand(path string, arg ...string) {
	cmd := exec.Command(arg[0], arg[1:]...)
	cmd.Dir = path
	if _, err := cmd.Output(); err != nil {
		printErrorStrings(getFunctionName(), fmt.Sprintf("%s, %v", path, arg), err)
	} else {
		printInfoStrings(getFunctionName(), fmt.Sprintf("%s, %v", path, arg))
	}
}

func GitCommandOut(path string, arg ...string) {
	cmd := exec.Command(arg[0], arg[1:]...)
	cmd.Dir = path

	output, err := cmd.CombinedOutput()
	if err != nil {
		printErrorStrings(getFunctionName(), fmt.Sprintf("%s, %v", path, arg), err)
	} else {
		printInfoStrings(getFunctionName(), fmt.Sprintf("%s, %v", path, arg))
		printInfoStrings(getFunctionName(), fmt.Sprintf("\n\n%s\n\n", output))
	}
}
