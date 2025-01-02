package main

import "net"

type Usr struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建一个usr的API
func NewUsr(con net.Conn) *Usr {
	userAddr := con.RemoteAddr().String()

	usr := &Usr{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: con,
	}

	//启动监听当前usr channel消息的goroutine
	go usr.ListenMessage()
	return usr
}

// 监听当前usr的channel的方法，一旦有消息，就直接发送给客户端
func (ths *Usr) ListenMessage() {
	for {
		msg := <-ths.C

		ths.conn.Write([]byte(msg + "\n"))
	}
}
