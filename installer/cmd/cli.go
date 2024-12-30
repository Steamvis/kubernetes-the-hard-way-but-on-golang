package main

import (
	"fmt"
	"installer/internal/args"
	"installer/internal/config"
	"installer/internal/logger"
	"installer/internal/ssh"
	"log"
)

func main() {
	logger.Init()

	argsConfig := args.GetArgs()
	err := config.Parse(argsConfig.ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	sshClient, err := ssh.NewConnect(
		config.Cfg.JumpHost.Username,
		config.Cfg.JumpHost.Password,
		config.Cfg.JumpHost.IP.String(),
	)
	if err != nil {
		log.Fatalln("Failed to dial: ", err)
	}
	defer sshClient.CloseConnect()

	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session: ", err)
	}
	defer session.CloseSession()

	result, err := session.RunCommand("echo $USER && echo `nproc` && cat /etc/os-release")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
