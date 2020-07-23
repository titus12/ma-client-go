package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	userMap map[int64]*TCPClient
)

func init()  {
	userMap = make(map[int64]*TCPClient)
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

	if eUid <= sUid {
		fmt.Println("命令格式错误, 第二个参数不能小于等于第一个参数....")
		return
	}

	for uid := sUid; uid<= eUid; uid ++ {
		user := &TCPClient{
			Host:   "127.0.0.1",
			Port:   8888,
			UserId: int64(uid),
		}
		user.Start()
		userMap[int64(uid)] = user
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

	if eUid <= sUid {
		fmt.Println("命令格式错误, 第二个参数不能小于等于第一个参数....")
		return
	}

	for uid := sUid; uid<= eUid; uid ++ {
		user, ok := userMap[int64(uid)]
		if ok {
			user.IntoGame()
		}
	}
}

func OpenLogPrint(msgstr interface{}) {
	print = true
}