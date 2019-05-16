package main

import (
	"fmt"
	"net"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8001")
	if err!=nil {
		fmt.Println("net.Dial err:",err)
		return
	}
	defer conn.Close()
	//客户端主动写数据
	conn.Write([]byte("Are you ok?"))

	//接收服务器回发的数据
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil {
		fmt.Println("conn.Read err:",err)
		return
	}

	fmt.Println("服务器回发：",string(buf[:n]))


}
