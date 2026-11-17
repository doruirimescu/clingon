package tmux

import (
	"errors"
	"io/fs"
)

// clingon's tmux window is macOS-only for now: these only keep the package
// building on Windows.

func foregroundGroup(pid int) (int, error) { return 0, errors.New("not supported on Windows") }

func flushInput(fd int) error { return errors.New("not supported on Windows") }

func setBlocking(fd int) {}

func ownedByUs(fi fs.FileInfo) bool { return false }

func listening(path string) bool { return false }
