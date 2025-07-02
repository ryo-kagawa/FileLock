package main

import (
	"errors"
	"slices"
)

type Args struct {
	LockType  string `key:"--lockType" default:"share"`
	FilePaths []string
}

func (a Args) Validate() error {
	if !slices.Contains([]string{"share", "exclusive"}, a.LockType) {
		return errors.New("invalid lockType")
	}
	if len(a.FilePaths) == 0 {
		return errors.New("filePath is empty")
	}
	return nil
}
