package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ryo-kagawa/go-utils/commandline"
	"golang.org/x/sys/windows"
)

type Command struct{}

var _ = (commandline.RootCommand)(Command{})

func (Command) Execute(arguments []string) (string, error) {
	args := Args{
		LockType: "share",
	}
	for _, argument := range arguments {
		if strings.HasPrefix(argument, "--lockType=") {
			args.LockType = argument
		} else {
			args.FilePaths = append(args.FilePaths, argument)
		}
	}
	if err := args.Validate(); err != nil {
		return "", err
	}

	for _, filePath := range arguments {
		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			return "", err
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			return "", err
		}

		handle := windows.Handle(file.Fd())
		overlapped := windows.Overlapped{}
		flags := uint32(windows.LOCKFILE_FAIL_IMMEDIATELY)
		if args.LockType == "exclusive" {
			flags |= windows.LOCKFILE_EXCLUSIVE_LOCK
		}

		if err := windows.LockFileEx(
			handle,
			uint32(flags),
			0,
			uint32(stat.Size()+1),
			0,
			&overlapped,
		); err != nil {
			return "", err
		}
	}

	fmt.Print("Enterキーでロック解除します")
	fmt.Scanln()

	return "finish", nil
}
