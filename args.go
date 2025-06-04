package main

import (
	"errors"
	"slices"
)

type Args struct {
	LockType string `key:"--lockType" default:"share"`
}

func (a Args) Validate() error {
	if !slices.Contains([]string{"share", "exclusive"}, a.LockType) {
		return errors.New("invalid lockType")
	}
	return nil
}
