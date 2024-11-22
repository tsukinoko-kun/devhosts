//go:build !windows
// +build !windows

package sudo

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func WriteFile(name string, data []byte, perm fs.FileMode) error {
	// write to temp file
	tmpfile := filepath.Join(os.TempDir(), "devhosts.tmp")
	if err := os.WriteFile(tmpfile, data, 0666); err != nil {
		return errors.Join(errors.New("failed to create temp file"), err)
	}
	defer os.Remove(tmpfile)

	cmd := exec.Command("sudo", "sh", "-c", "cat "+tmpfile+" > "+name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Join(errors.New("failed to write file"), err)
	}

	return nil
}
