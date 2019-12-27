package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "1030468074@qq.com", "一去、二三里")  // 发件人
	m.SetHeader("To",  // 收件人
		m.FormatAddress("283593214@qq.com", "乔峰"),
	)
	m.SetHeader("Subject", "Gomail")  // 主题
	m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>")  // 正文
	// 这里可能会因为dns解析问题，导致给的ip不准确，进而超时
	// 解决办法：可以去ip138.com，获取跟ISP运营商相同线路的ip（这里是移动线路），然后写到hosts文件中
	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewDialer("smtp.qq.com", 465, "1030468074@qq.com", "caosuqepcncybbjb")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

