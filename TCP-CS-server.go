package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器通信协议，IP地址和端口号
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()
	//阻塞监听客户端请求
    fmt.Println("服务端等待客户端建立连接。。。")

	conn, err:=listener.Accept()
	if err!=nil {
		fmt.Println("listener.Accept() err:",err)
		return
	}
	defer conn.Close()
	fmt.Println("服务端与客户端建立连接！！！")

	//读客户端发送的数据
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("conn.Read err:",err)
		return
	}
	conn.Write(buf[:n])

	//处理数据
	fmt.Println("服务器读到数据：",string(buf[:n]))

}

