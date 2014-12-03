package main

import (
	"bytes"
	"fmt"
	"git"
	"os/exec"
	"strings"
)

func main() {
	branches, err := git.MergedBranchList()
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range branches {
		if b == "" {
			continue
		}
		fmt.Println(b)
		cmd := exec.Command("git", "show", "--summary", "--pretty=format:%aN (%aE)", b)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// --summaryだとcreate/delete/rename modeも出るのでbranch名のみ
		r := strings.Split(out.String(), "\n")
		branch := r[0]
		fmt.Println(branch, "\n")
	}

}
