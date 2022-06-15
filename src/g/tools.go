package g

import (
	"bytes"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/errgo.v2/errors"
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
	log.Println(cmd.String())

	stdErr, stdOut := new(bytes.Buffer), new(bytes.Buffer)
	cmd.Stderr = stdErr
	cmd.Stdout = stdOut

	if e := cmd.Run(); e != nil {
		return errors.Note(e, nil, stdErr.String())
	}
	return nil
}
