// Package proc holds the few process calls that differ by OS.
package proc

import (
	"os/exec"
	"syscall"
)

// Alive reports whether a process with pid is running.
func Alive(pid int) bool {
	err := syscall.Kill(pid, 0)
	return err == nil || err == syscall.EPERM
}

// Detach starts cmd in a session of its own, off the terminal, so it
// outlives the process that started it.
func Detach(cmd *exec.Cmd) { cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true} }

// Exec replaces this process with path.
func Exec(path string, argv, env []string) error { return syscall.Exec(path, argv, env) }
