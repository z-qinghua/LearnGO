package process2

import (
	"LearnGo/chatRoom/common/message"
	"LearnGo/chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段
	Conn net.Conn
}

// 写一个serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码...
	// 1. 先从mes中取出mes.Data，直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}

	// 1. 先声明一个reMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2. 再声明一个loginResMes，并完成赋值
	var loginResMes message.LoginResMes

	// 如果用户id=100,密码=123456，认为合法，否则不合法
	// 实际生产环境，应到数据库中验证
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500 // 500状态码，表示该用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用"
	}

	// 3. 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// 4. 将data赋给resmes
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal faile, err =", err)
		return
	}

	// 6. 发送data封装到writePkg函数中
	// 因为使用了分层模式（mvc），我们先创建了一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
