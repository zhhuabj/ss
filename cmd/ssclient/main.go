package main

import (
	"fmt"
	"log"
	"net"
	"github.com/zhhuabj/ss/core"
	"github.com/zhhuabj/ss/config"
	"github.com/zhhuabj/ss/local"
)

const (
	DefaultListenAddr = ":7070"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)
	config := &config.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	if len(config.RemoteAddr) == 0 {
		log.Fatalln("There is no remote in .ss.json")
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	lsLocal := local.New(password, listenAddr, remoteAddr)
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("Use configuration：", fmt.Sprintf(`
local listen：
%s
remote remote：
%s
password password：
%s
	`, listenAddr, remoteAddr, password))
		log.Printf("ssclient:%s sucessfully listen on %s to remote address %s\n",
			version, listenAddr.String(), remoteAddr.String())
	}))
}