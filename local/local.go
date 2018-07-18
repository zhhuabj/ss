package local

import (
	"github.com/zhhuabj/ss/core"
	"log"
	"net"
)

type LsLocal struct {
	*core.SecureSocket
}

// 新建一个本地端
// 本地端的职责是:
// 1. 监听来自本机浏览器的代理请求
// 2. 转发前加密数据
// 3. 转发socket数据到墙外代理服务端
// 4. 把服务端返回的数据转发给用户的浏览器
func New(password *core.Password, listenAddr, remoteAddr *net.TCPAddr) *LsLocal {
	return &LsLocal{
		SecureSocket: &core.SecureSocket{
			Cipher:     core.NewCipher(password),
			ListenAddr: listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

// 本地端启动监听，接收来自本机浏览器的连接
func (local *LsLocal) Listen(didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		userConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// userConn被关闭时直接清除所有数据 不管没有发送的数据
		userConn.SetLinger(0)
		go local.handleConn(userConn)
	}
	return nil
}

func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()
	proxyServer, err := local.DialRemote()
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyServer.Close()
	//Conn被关闭时直接清除所有数据不管没有发送的数据
	proxyServer.SetLinger(0)

	// 进行转发, 从proxyServer读取数据发送到localUser
	go func() {
		err := local.DecodeCopy(userConn, proxyServer)
		if err != nil {
			// 在copy的过程中可能会存在网络超时等error被return，只要有一个发生了错误就退出本次工作
			userConn.Close()
			proxyServer.Close()
		}
	}()
	// 从 localUser 发送数据发送到 proxyServer，这里因为处在翻墙阶段出现网络错误的概率更大
	local.EncodeCopy(proxyServer, userConn)
}
