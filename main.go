package main

import (
	"github.com/titus12/ma-commons-go/net"
	"github.com/titus12/ma-commons-go/testconsole"
	pb "ma-client-go/pb"
)

func main() {
	net.LoadProtocol(&pb.Response{}, &pb.Request{})

	testconsole.EtcdRoot = "/root/backends/game"

	console := testconsole.NewConsole()
	console.Command("start", StartCommand)
	console.Command("send", SendCommand)
	console.Command("loop", LoopCommand)
	console.Command("openlog", OpenLogPrint)
	console.Command("closelog", CloseLogPrint)
	console.Command("query", testconsole.QueryRequest)

	console.Run()
}
