package utils

import (
	"net/smtp"
)

// ConfigInfo 邮箱服务器配置信息
type ConfigInfo struct {
	smtpAddr string
	smtpPort string
	secret   string
}

// EmailContent 邮件内容信息
type EmailContent struct {
	fromAddr    string
	contentType string
	theme       string
	message     string
	toAddr      string
}

func SendEmail(c *ConfigInfo, e *EmailContent) error {
	//  拼接smtp服务器地址
	smtpAddr := c.smtpAddr + ":" + c.smtpPort
	//  认证信息
	auth := smtp.PlainAuth("", e.fromAddr, c.secret, c.smtpAddr)
	//  配置邮件内容类型
	if e.contentType == "html" {
		e.contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		e.contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	//  发送
	msg := []byte("To: " + e.toAddr + "\r\n" +
		"From: " + e.fromAddr + "\r\n" +
		"Subject: " + e.theme + "\r\n" +
		e.contentType + "\r\n\r\n" +
		e.message)
	err := smtp.SendMail(smtpAddr, auth, e.fromAddr, []string{e.toAddr}, msg)
	if err != nil {
		return err
	}
	Debug()
	return nil
}

func Prepare(from, to, theme, message, secret string) (config ConfigInfo, content EmailContent) {
	// 收集配置信息
	config = ConfigInfo{
		// smtp服务器地址
		smtpAddr: "smtp.qq.com",
		// smtp服务器密钥
		secret: secret, //  hazylqlepicgbcai
		// smtp服务器端口
		smtpPort: "587",
	}
	// 收集邮件内容
	content = EmailContent{
		// 发件人
		fromAddr: from,
		// 收件人
		toAddr: to,
		// 邮件格式
		contentType: "text",
		// 邮件主题
		theme: theme,
		// 邮件内容
		message: message,
	}
	return
}
