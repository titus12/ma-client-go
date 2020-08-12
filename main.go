package main

import (
	"github.com/sirupsen/logrus"
	"github.com/titus12/ma-commons-go/net"
	"github.com/titus12/ma-commons-go/setting"
	"github.com/titus12/ma-commons-go/testconsole"

	"gopkg.in/urfave/cli.v2"
	pb "ma-client-go/pb"

	"os"
)

func main() {

	app := &cli.App{
		Name:    "控制台",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Value:   "Debug",
				Usage:   "log level of logrus",
				EnvVars: []string{"LOG_LEVEL"},
			},

			&cli.StringSliceFlag{
				Name:    "etcd-hosts",
				Value:   cli.NewStringSlice("http://127.0.0.1:2379"),
				Usage:   "etcd hosts",
				EnvVars: []string{"ETCD_HOST"},
			},

			&cli.StringFlag{
				Name:  "etcd-root",
				Value: "/backends",
				Usage: "etcd root path",
			},

			&cli.StringFlag{
				Name:  "vars-root",
				Value: "/vars",
				Usage: "etcd vars path",
			},

			&cli.StringSliceFlag{
				Name:    "kafka-brokers",
				Value:   cli.NewStringSlice("127.0.0.1:9092"),
				Usage:   "kafka brokers address",
				EnvVars: []string{"KAFKA_BROKERS"},
			},

			&cli.StringFlag{
				Name:  "log-topic",
				Value: "game-log",
				Usage: "game log topic",
			},
		},
		Action: action,
	}
	app.Run(os.Args)
}


func action(ctx *cli.Context) error {
	logrus.Println("log-level:", ctx.String("log-level"))
	logrus.Println("etcd-hosts:", ctx.StringSlice("etcd-hosts"))
	logrus.Println("etcd-root:", "/root" + ctx.String("etcd-root"))
	logrus.Println("vars-root:", "/root" + ctx.String("vars-root"))
	logrus.Println("kafka-brokers:", ctx.StringSlice("kafka-brokers"))
	logrus.Println("log-topic:", ctx.String("log-topic"))

	setting.EtcdRoot = "/root" + ctx.String("etcd-root")
	setting.EtcdHosts = ctx.StringSlice("etcd-hosts")


	net.LoadProtocol(&pb.Response{}, &pb.Request{})

	// todo: 目前用于测试，之后要重置 testconsole包，统一到client项目中
	testconsole.EtcdRoot = "/root/backends/game"
	console := testconsole.NewConsole()

	console.Command("start", StartCommand)
	console.Command("send", SendCommand)
	console.Command("loop", LoopCommand)
	console.Command("openlog", OpenLogPrint)
	console.Command("closelog", CloseLogPrint)
	console.Command("query", testconsole.QueryRequest)
	console.Command("listnode", ListNodeCommand)

	console.Run()
	return nil
}