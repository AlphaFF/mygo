package main

import (
	"fmt"
	"net"
	"time"
)

// Client 创建用户结构体类型
type Client struct {
	Name string
	Addr string
	C    chan Msg
}

// Msg 传递的消息类型
type Msg struct {
	Addr    string
	Content string
}

// 创建全局map, 存储在线用户
var onlineMap map[string]Client

// 创建全局channel, 传递用户信息
var message = make(chan Msg)

// WriteMsgToClient 向客户端写数据
func WriteMsgToClient(clnt Client, conn net.Conn) {
	// 监听用户自带channel上是否有消息
	for msg := range clnt.C {
		if msg.Addr != clnt.Addr {
			conn.Write([]byte(msg.Content + "\n"))
		}
	}
}

// MakeMsg 封装消息
func MakeMsg(clnt Client, msg string) (buf Msg) {
	buf = Msg{clnt.Name, "[" + clnt.Addr + "]" + clnt.Name + ": " + msg}
	return
}

// HandlerConnect 处理每一个请求链接
func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	netAddr := conn.RemoteAddr().String()
	clnt := Client{netAddr, netAddr, make(chan Msg)}

	// 将新连接用户, 添加到在线用户map中
	onlineMap[netAddr] = clnt
	// 创建专门给当前用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	// 发送用户上线消息到全局channel中
	message <- MakeMsg(clnt, "login")

	// 判断用户是否下线
	isQuit := make(chan bool)

	// 监听用户是否超时
	hasData := make(chan bool)

	// 创建go程, 专门处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println(clnt.Name, "客户端退出")
				return
			}
			if err != nil {
				fmt.Println("conn Read err:", err)
				return
			}
			msg := string(buf[:n-1])
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("users list:\n"))
				for _, user := range onlineMap {
					conn.Write([]byte(user.Addr + ":" + user.Name + "\n"))
				}
			} else if msg[:7] == "rename|" && len(msg) >= 8 {
				newName := msg[7:]
				clnt.Name = newName
				onlineMap[netAddr] = clnt
				conn.Write([]byte("rename successful.\n"))
			} else {
				message <- MakeMsg(clnt, msg)
			}
			hasData <- true
		}
	}()

	// 保证不退出
	for {
		select {
		case <-isQuit:
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		case <-hasData:
		case <-time.After(time.Second * 10):
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		}
	}
}

// Manager 处理服务器数据传递
func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-message
		for _, clnt := range onlineMap {
			fmt.Println(msg)
			clnt.C <- msg
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err.", err)
		return
	}
	defer listener.Close()

	// 创建管理者go程和全局channel
	go Manager()

	// 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err.", err)
			return
		}
		go HandlerConnect(conn)
	}
}
