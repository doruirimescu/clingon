package config

import "errors"

// Watch is how the menu hears tmux and the shell inside clingon's window.
// There's no such window on Windows (see `clingon print`), so nothing here
// can be watched.
type Watch struct{}

// WatchPath fails: file watching is macOS-only for now.
func WatchPath(path string) (*Watch, error) {
	return nil, errors.New("watching files isn't supported on Windows")
}

// Wait says false: there's nothing to wait for.
func (w *Watch) Wait() bool { return false }

// Close does nothing.
func (w *Watch) Close() {}
