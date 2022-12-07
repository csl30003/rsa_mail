package cmd

import (
	"RSA_mail/utils"
	"flag"
	"fmt"
	"os"
)

type MyFlagSet struct {
	*flag.FlagSet
	cmdComment string
}

func Cmd() {

	sendCmd := &MyFlagSet{
		flag.NewFlagSet("send", flag.ExitOnError),
		"发送邮件",
	}
	sendCmd.String("from", "", "你的邮箱地址")
	sendCmd.String("to", "", "对方的邮箱地址")
	sendCmd.String("theme", "", "主题")
	sendCmd.String("message", "", "内容")
	sendCmd.String("secret", "", "密钥")

	receiveCmd := &MyFlagSet{
		flag.NewFlagSet("receive", flag.ExitOnError),
		"接收邮件",
	}
	receiveCmd.String("user", "", "你的邮箱地址")
	receiveCmd.String("secret", "", "密钥")

	//  用map保存所有二级子命令 方便快速查找
	subCommands := map[string]*MyFlagSet{
		sendCmd.Name():    sendCmd,
		receiveCmd.Name(): receiveCmd,
	}

	usage := func() {
		//  整个命令行的帮助信息
		fmt.Printf("Usage: RSA_mail COMMAND\n\n")
		for _, v := range subCommands {
			fmt.Printf("%s %s\n", v.Name(), v.cmdComment)
			v.PrintDefaults() // 使用 flag 库自带的格式输出子命令的选项帮助信息
			fmt.Println()
		}
		os.Exit(2)
	}

	if len(os.Args) < 2 {
		//  即没有输入子命令
		usage()
	}

	//  第二个参数必须是我们支持的子命令
	cmd := subCommands[os.Args[1]]
	if cmd == nil {
		usage()
	}

	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if cmd.Name() == "send" {
		if cmd.NFlag() == 5 {
			//  发邮件
			encryptText, err := utils.RSAEncrypt([]byte(cmd.Lookup("message").Value.(flag.Getter).Get().(string)))
			if err != nil {
				fmt.Println(err)
				return
			}

			changeSlice := make([]byte, 1)
			changeSlice[0] = 255
			//  如果字节数组里有0 把它改
			for i := range encryptText {
				if encryptText[i] == 0 {
					encryptText[i] = 255
					changeSlice = append(changeSlice, byte(i))
				}
			}
			if len(changeSlice) > 1 {
				changeSlice[0] = byte(len(changeSlice) - 1)
			}

			finalText := append(changeSlice, encryptText...)
			//fmt.Println(finalText)
			//fmt.Println(len(finalText))

			config, content := utils.Prepare(
				cmd.Lookup("from").Value.(flag.Getter).Get().(string),
				cmd.Lookup("to").Value.(flag.Getter).Get().(string),
				cmd.Lookup("theme").Value.(flag.Getter).Get().(string),
				string(finalText),
				cmd.Lookup("secret").Value.(flag.Getter).Get().(string),
			)

			err = utils.SendEmail(&config, &content)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("发送成功")
		} else {
			usage()
		}
	} else if cmd.Name() == "receive" {
		if cmd.NFlag() == 2 {
			//  收邮件
			utils.Receive(
				cmd.Lookup("user").Value.(flag.Getter).Get().(string),
				cmd.Lookup("secret").Value.(flag.Getter).Get().(string),
			)
		} else { //yokvwwmicasndhbd
			usage()
		}
	} else {
		usage()
	}
}
