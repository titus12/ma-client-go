package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"sync/atomic"
	"time"
)

import (
	pb "ma-client-go/pb"
)

var (
	sequenceId uint32
	//userId     int64
)

var print = false

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
	UserId int64

	conn *net.TCPConn
}


func (c *TCPClient) Login() {
	message := &pb.LoginC2S{
		UserId:   c.UserId,
		Version:  "",
		Device:   "",
		Provider: "",
	}
	data := CreateMessage(101, message)
	c.conn.Write(data)
}

func (c *TCPClient) Ping() {
	message := &pb.PingC2S{
		Time: uint32(time.Now().Second()),
	}
	data := CreateMessage(100, message)
	c.conn.Write(data)
}

func (c *TCPClient) IntoGame() {
	message := &pb.IntoGameC2S{
		UserId: c.UserId,
	}
	data := CreateMessage(102, message)
	c.conn.Write(data)
}

func (c *TCPClient) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	c.conn = conn
	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			reply := make([]byte, 1024)
			_, err = conn.Read(reply)
			if err != nil {
				panic(err)
			}

			if print {
				fmt.Printf("received from server: [%s]\n", string(reply))
			}
		}
	}()


	c.Login()

	go func() {
		// 10秒一次心跳
		ticker := time.NewTicker(time.Second * 10)

		for {
			<- ticker.C
			c.Ping()
		}
	}()
}

func CreateMessage(protocolId uint16, msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal message error!!!")
		return nil
	}
	message := &Packet{}
	message.WriteU16(0)
	message.WriteU32(atomic.AddUint32(&sequenceId, 1))
	message.WriteU16(protocolId)
	message.WriteU16(0)
	message.WriteRawBytes(data)
	binary.BigEndian.PutUint16(message.Data(), uint16(message.Length()-2))
	return message.data
}
