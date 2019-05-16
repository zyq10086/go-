package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {

	defer conn.Close()
	addr:=conn.RemoteAddr()

	fmt.Println(addr,"客户端连接成功！")

	//循环读取客户端发送的数据
	buf:=make([]byte,4096)
	for  {
		n,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println("conn.Read:",err)
			return
		}
		fmt.Println("服务器读到：",string(buf[:n]))

		//小写转大写，会发给客户端
		x:=strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(x))
	}
}

func main() {
	//创建监听套接字
	listener,err:=net.Listen("tcp","127.0.0.1:8001")
	if err!=nil {
		fmt.Println("net.Listen:",err)
		return
	}
	defer listener.Close()

	//监听客户端请求连接
	fmt.Println("服务器等待客户端连接")
	for  {
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("listener.Accept：",err)
			return
		}
		//封装一个函数，具体完成服务端和客户端的通信
		go HandlerConnect(conn)
	}

}
