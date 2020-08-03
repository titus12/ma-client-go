package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	net2 "github.com/titus12/ma-commons-go/net"
	"github.com/titus12/ma-commons-go/utils"
	"io"
	"net"
	"sync/atomic"
	"time"
)

import (
	pb "ma-client-go/pb"
)

var print = false

// TCPClient struct
type TCPClient struct {
	SequenceId uint32

	Host string
	Port int
	UserId int64


	conn *net.TCPConn
}


func (c *TCPClient) Login() error {
	reqId := utils.GenUuidWithUint64()
	message := &pb.LoginC2S{
		UserId:   c.UserId,
		ReqId: reqId,
	}
	data := c.CreateMessage(101, message)
	_, err := c.conn.Write(data)
	return err
}

func (c *TCPClient) Ping() error {
	reqId := utils.GenUuidWithUint64()
	message := &pb.PingC2S{
		ReqId:                reqId,
		UserId:               c.UserId,
	}
	data := c.CreateMessage(100, message)
	_, err :=  c.conn.Write(data)
	return err
}

func (c *TCPClient) IntoGame() error {
	reqId := utils.GenUuidWithUint64()
	message := &pb.IntoGameC2S{
		ReqId:reqId,
		UserId: c.UserId,
	}
	data := c.CreateMessage(102, message)
	_, err := c.conn.Write(data)
	return err
}

// 解码
func (c *TCPClient) decode() error {
	sizebuf := make([]byte, 2)
	_, err := io.ReadFull(c.conn, sizebuf)
	if err != nil {
		return fmt.Errorf("TCPClient decode parse sizebuf err %w", err)
	}

	// 拿到整个协议的长度
	size := binary.BigEndian.Uint16(sizebuf)
	bodybuf := make([]byte, size)


	_, err = io.ReadFull(c.conn, bodybuf)
	if err != nil {
		return fmt.Errorf("TCPClient decode parse bodybuf err %w", err)
	}

	reader := net2.Reader(bodybuf)
	protocolCode, err := reader.ReadS16()
	if err != nil {
		return fmt.Errorf("TCPClient decode parse protocolCode err %w", err)
	}

	flagBits, err := reader.ReadU16()
	if err != nil {
		return fmt.Errorf("TCPClient decode parse flagBits err %w", err)
	}

	errCode, err := reader.ReadS16()
	if err != nil {
		return fmt.Errorf("TCPClient decode parse errCode err %w", err)
	}

	if errCode != 0 {
		return fmt.Errorf("TCPClient ErrCode %d", errCode)
	}

	payload, err := reader.ReadBinary()
	if err != nil {
		return fmt.Errorf("TCPClient decode parse payload err %w", err)
	}

	msg, err := net2.Deserialize(protocolCode, flagBits, payload)
	if err != nil {
		return fmt.Errorf("TCPClient decode parse net2.Deserialize err %w", err)
	}

	if protoMsg, ok := msg.(*pb.LoginS2C); ok {
		setUser(c)
		fmt.Printf("ProtocolCode Login Message %v\n", protoMsg)
	} else {
		if print {
			fmt.Printf("ProtocolCode %d Message %v\n", protocolCode, msg)
		}
	}

	return nil
}

func (c *TCPClient) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//defer conn.Close()
	if err != nil {
		panic(err)
	}

	c.conn = conn
	go func() {
		for {
			if err := c.decode(); err != nil {
				delUser(c.UserId)
				fmt.Printf("删除玩家[%d]ERROR:%v\n",c.UserId, err)
				return
			}
		}
	}()


	err = c.Login()
	if err != nil {
		fmt.Println("Login Err: ", err)
		return
	}



	go func() {
		// 10秒一次心跳
		ticker := time.NewTicker(time.Second * 30)

		for {
			<- ticker.C
			err := c.Ping()
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		ticker.Stop()
	}()
}


func (c *TCPClient) CreateMessage(protocolId uint16, msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal message error!!!")
		return nil
	}
	message := &net2.Packet{}
	message.WriteU16(0)
	message.WriteU32(atomic.AddUint32(&c.SequenceId, 1))
	message.WriteU16(protocolId)
	message.WriteU16(0)
	message.WriteRawBytes(data)
	binary.BigEndian.PutUint16(message.Data(), uint16(message.Length()-2))
	return message.Data()
}
