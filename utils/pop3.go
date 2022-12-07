package utils

import (
	"fmt"
	"github.com/knadh/go-pop3"
	"io"
	"log"
)

func Receive(user, secret string) {
	p := pop3.New(pop3.Opt{
		Host:       "pop.qq.com",
		Port:       995,
		TLSEnabled: true,
	})

	c, err := p.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	if err := c.Auth(user, secret); err != nil {
		log.Fatal(err)
	}

	count, _, _ := c.Stat()

	for id := 1; id <= count; id++ {
		m, _ := c.Retr(id)

		fmt.Println("========================================================================================")
		fmt.Println("|", "第", id, "封邮件")
		fmt.Println("|", "发件人", m.Header.Get("From"))
		fmt.Println("|", "收件人", m.Header.Get("To"))
		fmt.Println("|", "标题", m.Header.Get("Subject"))

		content, err := io.ReadAll(m.Body)
		var newContent []byte
		if content[0] != 255 {
			temp := content[:content[0]+1]
			newContent = content[content[0]+1:]
			for i := 1; i <= int(temp[0]); i++ {
				//fmt.Println(newContent[temp[i]])
				newContent[temp[i]] = 0
			}
		} else {
			newContent = content[1:]
		}

		noChange := true
		newContent = newContent[:len(newContent)-2]
		for noChange {
			noChange = false
			for i := range newContent {
				if newContent[i] == 13 && i < len(newContent)-1 && newContent[i+1] == 10 {
					noChange = true
					newContent = append(newContent[:i], newContent[i+1:]...)
					break
				}
			}
		}

		//fmt.Println(newContent)
		//fmt.Println(len(newContent))
		decryptText, err := RSADecrypt(newContent)
		if err != nil {
			fmt.Println("|", "内容", err)
		} else {
			fmt.Println("|", "内容", string(decryptText))
		}
	}
	fmt.Println("========================================================================================")
}

func Debug() {
	p := pop3.New(pop3.Opt{
		Host:       "pop.qq.com",
		Port:       995,
		TLSEnabled: true,
	})

	c, err := p.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	if err := c.Auth("1733786384@qq.com", "yokvwwmicasndhbd"); err != nil {
		log.Fatal(err)
	}

	c.Stat()
}
