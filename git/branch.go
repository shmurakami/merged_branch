package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func branch(name string, args ...string) ([]string, error) {
	var out, e bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &out
	cmd.Stderr = &e

	err := cmd.Run()
	if err != nil {
		fmt.Println(e.String())
		return nil, err
	}

	o := out.String()
	branches := strings.Split(o, "\n")

	result := make([]string, len(branches))
	i := 0
	for _, b := range branches {
		if b == "" {
			continue
		}

		// ignore HEAD
		m, _ := regexp.MatchString("origin/HEAD", b)
		if m {
			continue
		}

		// current branch
		m, _ = regexp.MatchString("^\\*", b)
		if m {
			b = strings.Replace(b, "*", "", 1)
		}

		b = strings.TrimLeft(b, " ")
		result[i] = b
		i++
	}

	return result, nil
}

func BrancheList() ([]string, error) {
	return branch("git", "branch", "-a")
}

func MergedBranchList() ([]string, error) {
	return branch("git", "branch", "-a", "--merged")
}
