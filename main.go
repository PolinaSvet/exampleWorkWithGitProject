package main

import (
	"bufio"
	"fmt"
	"mygitpack/gitpack"
	"os"
	"strings"
	"time"
)

const (
	pathWorkDir        string = "workDir"
	pathWorkDirForGit  string = "workDir/prgForGit"
	pathCloneDirForGit string = "prgForGitClone"
	pathWorkFile       string = "main.go"
	contentInitFile    string = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"=== v0 === Hello, world!!!\")\n}"
	commitInitFile     string = "first commit"
	tagInitFile        string = "v0"
	branchName         string = "main"
	originName         string = "origin"
	pathGit            string = "https://github.com/PolinaSvet/exampleWorkWithGit.git"
	countChange        int    = 5
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Application to demonstrate working with Git ===")
		fmt.Println("Input command:")
		fmt.Println("1. git init,... -> format input <1>                     Create new file, change and push to Git")
		fmt.Println("2. git log      -> format input <2>                     Shows the commit logs")
		fmt.Println("3. git show     -> format input <3 v1>, vX - name tag.  Shows one or more objects tags")
		fmt.Println("4. git clone    -> format input <4 v2>, vX - name tag.  Clone a repository into a new directory")
		fmt.Println("5. git reset    -> format input <5 v3>, vX - name tag.  Hard reset current HEAD to the specified state")
		fmt.Println("0. Exit")

		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("=== 1 ======================")
			menuCreateAndCommitAppGit()
		case "2":
			fmt.Println("=== 2 ======================")
			menuLogGit()
		case "0":
			fmt.Println("Exit")
			return
		default:
			arg1, arg2 := parseInputString(input)
			switch arg1 {
			case "3":
				fmt.Println("=== 3 ======================")
				menuShowGit(arg2)
			case "4":
				fmt.Println("=== 4 ======================")
				menuCloneCheckoutGit(arg2)
			case "5":
				fmt.Println("=== 5 ======================")
				menuResetHardGit(arg2)
			default:
				fmt.Println("Wrong choice. Try again.")
			}
		}

		fmt.Println()
	}

}

func parseInputString(input string) (string, string) {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// 1
func menuCreateAndCommitAppGit() {
	// create dir and file for work
	gitpack.GitRemoveDir(pathWorkDir)
	gitpack.GitCreateDir(pathWorkDir)
	gitpack.GitCreateDir(pathWorkDirForGit)
	gitpack.GitCreateFile(pathWorkDirForGit, pathWorkFile, contentInitFile)

	// …or create a new repository on the command line
	gitpack.GitCommand(pathWorkDirForGit, "git", "init")
	gitpack.GitCommand(pathWorkDirForGit, "git", "add", ".")
	gitpack.GitCommand(pathWorkDirForGit, "git", "commit", "-m", commitInitFile)
	gitpack.GitCommand(pathWorkDirForGit, "git", "tag", tagInitFile)
	gitpack.GitCommand(pathWorkDirForGit, "git", "branch", "-M", branchName)
	gitpack.GitCommand(pathWorkDirForGit, "git", "remote", "add", originName, pathGit)
	gitpack.GitCommand(pathWorkDirForGit, "git", "push", "-u", originName, branchName, tagInitFile)

	//…or push an existing repository from the command line
	for i := 1; i <= countChange; i++ {
		tagChangeFile := fmt.Sprintf("v%d", i)
		commitChangeFile := fmt.Sprintf("Code was updated %d times on %s.", i, time.Now().Format("2006-01-02 15:04:05"))
		contentChangeFile := fmt.Sprintf("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"=== %s === %s Hello, world!!!\")\n}", tagChangeFile, commitChangeFile)

		gitpack.GitChangeFile(pathWorkDirForGit, pathWorkFile, contentChangeFile)
		gitpack.GitCommand(pathWorkDirForGit, "git", "add", ".")
		gitpack.GitCommand(pathWorkDirForGit, "git", "commit", "-m", commitChangeFile)
		gitpack.GitCommand(pathWorkDirForGit, "git", "tag", tagChangeFile)
		gitpack.GitCommand(pathWorkDirForGit, "git", "push", "-u", originName, branchName, tagChangeFile)
	}
}

// 2
func menuLogGit() {
	gitpack.GitCommandOut(pathWorkDirForGit, "git", "log")
}

// 3
func menuShowGit(arg string) {
	gitpack.GitCommandOut(pathWorkDirForGit, "git", "show", arg)
}

// 4
func menuCloneCheckoutGit(arg string) {
	dirPath := fmt.Sprintf("%s/%s", pathWorkDir, pathCloneDirForGit)
	gitpack.GitRemoveDir(dirPath)
	gitpack.GitCommandOut(pathWorkDir, "git", "clone", pathGit, pathCloneDirForGit)
	if arg != "" {
		gitpack.GitCommand(pathWorkDir, "git", "-C", pathCloneDirForGit, "checkout", arg)
	}
	gitpack.GitCommandOut(dirPath, "go", "run", pathWorkFile)
}

// 5
func menuResetHardGit(arg string) {
	gitpack.GitCommandOut(pathWorkDirForGit, "git", "reset", "--hard", arg)
	gitpack.GitCommandOut(pathWorkDirForGit, "git", "push", "--force")
}
