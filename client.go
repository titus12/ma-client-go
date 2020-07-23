package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

import (
	pb "ma-client-go/pb"
)

var (
	sequenceId uint32
	userId     int64
)

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
}

func ReadLine() (string, error) {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimRight(text, "\n\r")
	return text, nil
}

// Start TCPClient
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
	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			reply := make([]byte, 1024)
			_, err = conn.Read(reply)
			if err != nil {
				panic(err)
			}
			fmt.Printf("received from server: [%s]\n", string(reply))
		}
	}()

	for {
		fmt.Print(">> ")
		text, err := ReadLine()
		if err != nil {
			panic("Error reading console input")
		}
		command := strings.Split(text, " ")
		switch command[0] {
		case "LoginC2S":
			if len(command) != 2 {
				fmt.Printf("命令解析错误，LEN != 2\n")
				break
			}
			userId, _ = strconv.ParseInt(command[1], 10, 64)
			message := &pb.LoginC2S{
				UserId:   userId,
				Version:  "",
				Device:   "",
				Provider: "",
			}
			data := CreateMessage(101, message)
			conn.Write(data)
		case "PingC2S":
			message := &pb.PingC2S{
				Time: uint32(time.Now().Second()),
			}
			data := CreateMessage(100, message)
			conn.Write(data)
		case "IntoGameC2S":
			message := &pb.IntoGameC2S{
				UserId: userId,
			}
			data := CreateMessage(102, message)
			conn.Write(data)
		default:
			fmt.Println("command not exist..")
		}
	}
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
