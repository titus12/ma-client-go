package main

import (
	"github.com/titus12/ma-commons-go/testconsole"
)

func main() {
	console := testconsole.NewConsole()
	console.Command("start", StartCommand)
	console.Command("send", SendCommand)
	console.Command("openlog", OpenLogPrint)

	console.Run()
}
