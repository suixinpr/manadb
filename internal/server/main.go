package server

import (
	"bufio"
	"github/suixinpr/manadb/internal/session"
	"log"
	"net"
)

func ListenAndServe(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	defer listener.Close()
	log.Printf("bind: %s, start listening...", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept err: %v", err)
		}

		/* 处理连接 */
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		log.Println("session exit")
		conn.Close()
	}()

	cli := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	sess := session.NewSesion(cli)

	/* 开始运行会话 */
	if err := sess.Main(); err != nil {
		log.Println(err)
	}
}
