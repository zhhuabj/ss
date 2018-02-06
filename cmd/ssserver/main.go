package main

import (
	"fmt"
	"log"
	"net"
	"github.com/phayes/freeport"
	"github.com/zhhuabj/ss/core"
	"github.com/zhhuabj/ss/config"
	"github.com/zhhuabj/ss/server"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// random port
	port, err := freeport.GetFreePort()
	if err != nil {
		port = 7448
	}
	config := &config.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		Password: core.RandPassword().String(),
	}
	config.ReadConfig()
	config.SaveConfig()

	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	lsServer := server.New(password, listenAddr)
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("Use configuration：", fmt.Sprintf(`
local listen：
%s
password：
%s
	`, listenAddr, password))
		log.Printf("ssserver:%s successfully listen on %s\n", version, listenAddr.String())
	}))
}
