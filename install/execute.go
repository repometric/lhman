package install

import (
	"bytes"
	"os/exec"
)

// Execute function runs binary with Context and arguments. Result: (string stdout, string stderr)
func Execute(context ExecuteContext) (string, string) {
	cmd := exec.Command(context.Binary, context.Args...)
	if len(context.WorkingDirectory) > 0 {
		cmd.Dir = context.WorkingDirectory
	}
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", context.Binary + " crashed with: " + err.Error()
	}

	return string(stdout.Bytes()), string(stderr.Bytes())
}
