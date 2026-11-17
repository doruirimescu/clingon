package config

import "errors"

func mkfifo(path string, mode uint32) error { return errors.New("no named pipes on Windows") }
