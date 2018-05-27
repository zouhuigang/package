package main

import (
	"github.com/zouhuigang/package/zsend"
	"log"
)

func main() {
	/*smtp_host = smtp.qq.com
	  smtp_port = 25
	  smtp_username = notifications@anooc.com
	  smtp_password = gfsnnsqjcxhwbcgj
	  from_email = notifications@anooc.com*/

	username := "notifications@anooc.com"
	host := "smtp.qq.com"
	password := "gfsnnsqjcxhwbcgj"
	port := 25
	subject := "主题"
	content := "内容"
	contentType := "text/html" //text/plain
	attach := ""               //附件路径
	to := []string{"903788390@qq.com", "zouhuigang888@gmail.com"}
	cc := []string{"903788390@qq.com", "952750120@qq.com"}

	message := zsend.NewEmailMessage(username, subject, contentType, content, attach, to, cc)
	email := zsend.NewEmailClient(host, username, password, port, message)
	ok, err := email.SendMessage()
	if err != nil {
		log.Fatalf("发送邮件失败了: %s", err)
	}

	log.Println(ok)

}
