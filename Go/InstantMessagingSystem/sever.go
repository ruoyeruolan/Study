package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Sever struct {
	IP   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*Usr
	mapLock   sync.RWMutex
	Message   chan string
}

// 创建一个sever的接口
func NewSever(ip string, port int) *Sever {
	sever := &Sever{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*Usr),
		Message:   make(chan string),
	}
	return sever
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (ths *Sever) ListenMessager() {
	for {
		msg := <-ths.Message

		//将msg发送给全部的在线user
		ths.mapLock.Lock() //加锁
		for _, cli := range ths.OnlineMap {
			cli.C <- msg
		}
		ths.mapLock.Unlock() //解锁
	}
}

func (ths *Sever) BroadCast(usr *Usr, msg string) {
	sendMsg := "[" + usr.Addr + "]" + usr.Name + ":" + msg

	ths.Message <- sendMsg
}

func (ths *Sever) Handler(conn net.Conn) {
	//当前连接的业务
	// fmt.Println("Connection Successfully!")

	usr := NewUsr(conn)
	//用户上线， 将用户加入到onlinemap中
	ths.mapLock.Lock()
	ths.OnlineMap[usr.Name] = usr
	ths.mapLock.Unlock()
	//广播当前用户上线消息
	ths.BroadCast(usr, "Online")

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				ths.BroadCast(usr, "Offline")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			//提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])
			//将得到的消息进行广播
			ths.BroadCast(usr, msg)
		}

	}()

	//阻塞当前handler
	select {}
}

// 启动服务器的接口
func (ths *Sever) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ths.IP, ths.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go ths.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		//do handler
		go ths.Handler(conn)

	}

}
