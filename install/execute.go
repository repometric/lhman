package install

import (
	"bytes"
	"os/exec"
)

// Execute function runs binary with Context and arguments. Result: (string stdout, string stderr)
func Execute(context Context, binary string, args ...string) (string, string) {
	cmd := exec.Command(binary, args...)
	if len(context.Folder) > 0 {
		cmd.Dir = context.Folder
	}
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", binary + " crashed with: " + err.Error()
	}

	return string(stdout.Bytes()), string(stderr.Bytes())
}
