package main

import (
	"bufio"
	"flag"
	"fmt"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/pkg/driver"
	"net"
	"os"
	"strings"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "host, default is 127.0.0.1")
	flag.IntVar(&port, "p", 9315, "port number, default is 9315")
}

func main() {
	flag.Parse()
	/* 建立连接 */
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	defer conn.Close()

	/* 用户和数据库 */
	user := bufio.NewReader(os.Stdin)
	server := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	db := driver.NewConnection(server)
	handle(user, db)
}

func handle(user *bufio.Reader, db *driver.Connection) {
	/* 主循环 */
	for {
		fmt.Printf("manasql> ")

		/* 读取用户输入 */
		msg, err := user.ReadString('\n')
		if err != nil {
			fmt.Println("reading message error:", err)
			continue
		}

		/* 发送 SQL */
		err = db.SendSQL(msg)
		if err != nil {
			fmt.Println("send sql to db error:", err)
			continue
		}

		/* 接收 SQL 执行结果 */
		err = db.RecvSQLResult()
		if err != nil {
			fmt.Println("recv error:", err)
			continue
		}

		print(db)
		db.Desc = nil
		db.Rows = nil
		db.Err = ""
	}
}

func print(db *driver.Connection) {
	if db.Err != "" {
		fmt.Println(db.Err)
		return
	}

	/* 记录列宽 */
	width := make([]int, len(db.Desc))
	for i, desc := range db.Desc {
		width[i] = max(width[i], len(desc.Name))
	}
	for _, row := range db.Rows {
		for i, value := range row {
			width[i] = max(width[i], len(value))
		}
	}

	for i := range width {
		width[i] += 2
	}

	/* 打印描述符 */
	for i, desc := range db.Desc {
		if i > 0 {
			fmt.Print("|")
		}
		space := width[i] - len(desc.Name)
		left := space / 2
		printOne(desc.Name, left, space-left)
	}
	fmt.Println()

	/* 打印分割线 */
	for i, w := range width {
		if i > 0 {
			fmt.Print("+")
		}
		fmt.Print(strings.Repeat("-", w))
	}
	fmt.Println()

	/* 打印行数据 */
	for _, row := range db.Rows {
		for i, value := range row {
			if i > 0 {
				fmt.Print("|")
			}
			space := width[i] - len(value)
			switch metadata.OID(db.Desc[i].TypeID) {
			case metadata.CharID:
				fallthrough
			case metadata.TextID:
				fallthrough
			case metadata.VarcharID:
				printOne(value, 1, space-1)
			default:
				printOne(value, space-1, 1)
			}
		}
		fmt.Println()
	}
}

func printOne(s string, l int, r int) {
	fmt.Print(strings.Repeat(" ", l))
	fmt.Print(s)
	fmt.Print(strings.Repeat(" ", r))
}
