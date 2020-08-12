package main

import (
	"context"
	"fmt"
	"github.com/titus12/ma-commons-go/control"
	control2 "github.com/titus12/ma-commons-go/control/proto"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	userMap map[int64]*TCPClient
	mu sync.RWMutex
)

func init()  {
	userMap = make(map[int64]*TCPClient)
}


func getUser(id int64) *TCPClient {
	mu.RLock()
	defer mu.RUnlock()
	return userMap[id]
}

func setUser(u *TCPClient) {
	mu.Lock()
	defer mu.Unlock()

	userMap[u.UserId] = u
}

func delUser(id int64) {
	mu.Lock()
	defer mu.Unlock()

	delete(userMap, id)
}


// 列出指定服务的所有节点
func ListNodeCommand(msgstr interface{}) {
	str, ok := msgstr.(string)
	if !ok {
		fmt.Println("不能正常转换成 string 消息")
		return
	}
	strarr := strings.Split(str, " ")

	serviceName := strarr[0]

	nodes := control.GetAllNodeData(serviceName)

	for _, n := range nodes {
		fmt.Println(n)
	}
}


func StopNodeCommand(msgstr interface{}) {
	str, ok := msgstr.(string)
	if !ok {
		fmt.Println("不能正常转换成 string 消息")
		return
	}
	strarr := strings.Split(str, " ")

	serviceName := strarr[0]
	id := strarr[1]

	nodeKey := fmt.Sprintf("%s%s", serviceName, id)

	nodes := control.GetAllNodeData(serviceName)

	var hitNode *control.NodeInfo
	for idx , n := range nodes {
		if n.Key == nodeKey {
			hitNode = nodes[idx]
			break
		}
	}

	if hitNode == nil {
		fmt.Println("没有找到合适的节点，节点不存在....")
		return
	}

	// 执行远程调用
	control.ExecGrpcCall(hitNode, func(ctx context.Context, cli control2.ControlServiceClient) error {
		_, err := cli.StopNode(ctx, &control2.Request{
			NodeKey:              hitNode.Key,
			Addr:                 hitNode.Addr,
		})
		return err
	})
}



func StartCommand(msgstr interface{}) {
	str, ok := msgstr.(string)
	if !ok {
		fmt.Println("不能正常转换成 string 消息")
		return
	}

	strarr := strings.Split(str, " ")

	sUid, err := strconv.Atoi(strarr[0])
	if err != nil {
		fmt.Println("命令格式错误, 第一个参数需要是数字...", strarr[0])
		return
	}

	eUid, err := strconv.Atoi(strarr[1])
	if err != nil {
		fmt.Println("命令格式错误, 第二个参数需要是数字...", strarr[1])
		return
	}

	if eUid < sUid {
		fmt.Println("命令格式错误, 第二个参数不能小于等于第一个参数....")
		return
	}

	for uid := sUid; uid<= eUid; uid ++ {
		u := getUser(int64(uid))
		if u != nil {
			fmt.Println("已经存在用户， 先进行关闭......")
			u.conn.Close()
			delUser(int64(uid))
		}

		user := &TCPClient{
			Host:   "127.0.0.1",
			Port:   8888,
			UserId: int64(uid),
		}
		user.Start()
	}
}


func LoopCommand(msgstr interface{}) {
	str, ok := msgstr.(string)
	if !ok {
		fmt.Println("不能正常转换成 string 消息")
		return
	}

	strarr := strings.Split(str, " ")
	if len(strarr) != 3 {
		fmt.Println("命令格式不正确，需要3个参数...")
		return
	}

	sUid, err := strconv.Atoi(strarr[0])
	if err != nil {
		fmt.Println("命令格式错误, 第一个参数需要是数字...", strarr[0])
		return
	}

	eUid, err := strconv.Atoi(strarr[1])
	if err != nil {
		fmt.Println("命令格式错误, 第二个参数需要是数字...", strarr[1])
		return
	}

	if eUid < sUid {
		fmt.Println("命令格式错误, 第二个参数不能小于等于第一个参数....")
		return
	}

	num, err := strconv.Atoi(strarr[2])
	if err != nil {
		fmt.Println("命令格式不正确，第3个参数必须是数字")
		return
	}

	for i:=0; i<num; i++ {
		_send(int64(sUid), int64(eUid))
		time.Sleep(time.Second)
	}
}


func _send(sUid, eUid int64) {
	for uid := sUid; uid<= eUid; uid ++ {
		user := getUser(uid)
		if user != nil {
			err := user.IntoGame()
			if err != nil {
				fmt.Println("IntoGame Err ", err)
			}
		}
	}
}

func SendCommand(msgstr interface{}) {
	str, ok := msgstr.(string)
	if !ok {
		fmt.Println("不能正常转换成 string 消息")
		return
	}

	strarr := strings.Split(str, " ")

	sUid, err := strconv.Atoi(strarr[0])
	if err != nil {
		fmt.Println("命令格式错误, 第一个参数需要是数字...", strarr[0])
		return
	}

	eUid, err := strconv.Atoi(strarr[1])
	if err != nil {
		fmt.Println("命令格式错误, 第二个参数需要是数字...", strarr[1])
		return
	}

	if eUid < sUid {
		fmt.Println("命令格式错误, 第二个参数不能小于等于第一个参数....")
		return
	}

	_send(int64(sUid), int64(eUid))
}

func OpenLogPrint(msgstr interface{}) {
	print = true
}

func CloseLogPrint(msgstr interface{}) {
	print = false
}