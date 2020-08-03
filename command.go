package main

import (
	"fmt"
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