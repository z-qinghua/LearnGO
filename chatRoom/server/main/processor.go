package main

import (
	"LearnGo/chatRoom/common/message"
	"LearnGo/chatRoom/server/process2"
	"LearnGo/chatRoom/server/utils"
	"fmt"
	"io"
	"net"
)

// 先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	// 循环处理客户端发送的信息
	for {

		// 创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return err
			} else {
				fmt.Println("readPkg err =", err)
				return err
			}
		}
		// fmt.Println("mes =", mes)
		err = this.serverProcessMess(&mes)
		if err != nil {
			return err
		}
	}
}
