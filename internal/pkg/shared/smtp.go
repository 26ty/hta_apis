package shared

import (
	"net/smtp"
)

// 以下 variable 可參考 Gmail 的 smtp 設定說明
// var (
// 	host     = "smtp.gmail.com"
// 	username = "example@gmail.com"
// 	password = "Pass123"
// 	port = "587"
// )

func SendEmail(host string, port string, name string, username string, password string, to string,subject string, body string) (err error){
	hostport := host + ":" + port
	auth := smtp.PlainAuth("", username, password, host)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte("From:"+name+"<"+username+">"+"\n"+"Subject: "+subject+"\n" + mime + body)

	err = smtp.SendMail(
		hostport,
		auth,
		username,
		[]string{to},
		msg,
	)
	
	return err
}