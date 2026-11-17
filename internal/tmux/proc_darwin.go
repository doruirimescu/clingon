package tmux

import (
	"io/fs"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

// foregroundGroup is the process group running in the terminal that pid
// (a shell) belongs to — the shell itself when it's at its prompt.
func foregroundGroup(pid int) (int, error) {
	k, err := unix.SysctlKinfoProc("kern.proc.pid", pid)
	if err != nil {
		return 0, err
	}
	return int(k.Eproc.Tpgid), nil
}

// flushInput discards keys typed on the terminal but not yet read.
func flushInput(fd int) error { return unix.IoctlSetPointerInt(fd, unix.TIOCFLUSH, unix.TCIFLUSH) }

// setBlocking puts the terminal's descriptor back to blocking reads.
func setBlocking(fd int) { _ = syscall.SetNonblock(fd, false) }

// ownedByUs reports whether fi is ours alone: ours, and closed to others.
func ownedByUs(fi fs.FileInfo) bool {
	st, ok := fi.Sys().(*syscall.Stat_t)
	return ok && int(st.Uid) == os.Getuid() && fi.Mode().Perm()&0o077 == 0
}

// listening reports whether a server accepts connections on unix socket path.
func listening(path string) bool {
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return true // can't tell: leave it be
	}
	defer func() { _ = syscall.Close(fd) }()
	return syscall.Connect(fd, &syscall.SockaddrUnix{Name: path}) == nil
}
