package g

import (
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os/exec"
	"runtime"
	"strings"
)

func Clone(u string, target string) error {
	addr, e := Parse(u)
	if e != nil {
		return e
	}

	return clone(*addr, target)
}

func clone(addr Address, target string) error {
	cmdArgs := []string{"git", "clone"}
	cmdArgs = append(cmdArgs, addr.SSHUrl())
	cmdArgs = append(cmdArgs, target)

	cmd := exec.Command("bash", "-c", strings.Join(cmdArgs, " "))
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", strings.Join(cmdArgs, " "))
	}
	cmd.Start()
	if e := cmd.Wait(); e != nil {
		stdErr, err := cmd.CombinedOutput()
		var msg string
		if err == nil {
			msg = fmt.Sprintf("%v", stdErr)
		} else {
			msg = err.Error()
		}
		return errors.Note(e, nil, msg)
	}
	return nil
}
