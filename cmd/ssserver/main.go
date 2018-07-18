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
const DefaultPassword = "NIOXTjvBR71xQHtZ1YELJqLgdmKGLKeZxo0XPKa2M5whAD1asqpXs2sBclgtdS/1RF9ud5rIlhSCeYtbkyMDYRldG/Bw3MnuidgnFdTPzOQusXQHRSvTkmX9VNlQuTLXnU3bc3iuZPbe/8XhMOz8NkrCFiXontYCytHqjH6Pw2oGwI6gzn9RtTjjXClSEJBBTNrn64Uobx4NkfLt8woEqKS0Xq2hTzUaPxhpx5TlmBGvVW3fuPljvPTxSaM+DvtsZnoPpd0qn/jQhzm7Ih3ifNJTgCQxfYpgQ/73y2cgaBKwCamEq7eVRs0fBelWQpv6SBMMOr6sxAhLv+83uuYciA=="

func main() {
	log.SetFlags(log.Lshortfile)
	config := &config.Config{}
	config.ReadConfig()
	if len(config.ListenAddr) == 0 {
		// random port
		port, err := freeport.GetFreePort()
		if err != nil {
			port = 8388
		}
		config.SetListenAddr(fmt.Sprintf(":%d", port))
	}
	// Just use defaut password when .ss.config doesn't configure password
	if len(config.Password) == 0 {
//		config.SetPassword(core.RandPassword().String())
		config.SetPassword(DefaultPassword)
	}
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
