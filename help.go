package main

import (
	"strings"

	"github.com/ryo-kagawa/go-utils/commandline"
)

type Help struct{}

var _ = (commandline.SubCommand)(Help{})

var usage = `
Usage: FileLock [--lockType] filePath [filePath ...]
--lockType string
    share: 共有ロック（デフォルト）
    exclusive: 排他ロック
filePath: ロックしたいファイルを複数指定します
Usage: FileLock help
    ここを表示します
Usage: FileLock version
    バージョンを表示します
`

func (Help) Execute(arguments []string) (string, error) {
	return strings.Trim(usage, "\n"), nil
}

func (h Help) Name() string {
	return "help"
}
