package driver

import (
	"bufio"
	"errors"
	"github/suixinpr/manadb/pkg/protocol"
)

type Connection struct {
	server *bufio.ReadWriter

	/* 接收到的错误 */
	Err string

	/* 接收到的数据集 */
	Desc []RowDesc
	Rows [][]string
}

func NewConnection(server *bufio.ReadWriter) *Connection {
	return &Connection{server: server}
}

/********************************************************************************
*
*  Error
*
********************************************************************************/

/* 接收 error 信息 */
func (conn *Connection) RecvError() error {
	s, err := protocol.RecvString(conn.server.Reader)
	conn.Err = s
	return err
}

/********************************************************************************
*
*  SQL
*
********************************************************************************/

/* 发送 SQL */
func (conn *Connection) SendSQL(sql string) error {
	protocol.SendMessage(conn.server.Writer, protocol.MessageSQL)
	protocol.SendString(conn.server.Writer, sql)
	return conn.server.Writer.Flush()
}

/* 接收 SQL 执行结果 */
func (conn *Connection) RecvSQLResult() error {
	for {
		msg, err := protocol.RecvMessage(conn.server.Reader)
		if err != nil {
			return err
		}

		switch msg {
		case protocol.MessageError:
			return conn.RecvError()
		case protocol.MessageFinish:
			return nil
		case protocol.MessageRowData:
			err = conn.RecvRowsData()
		case protocol.MessageRowDesc:
			err = conn.RecvRowDescs()
		default:
			err = errors.New("")
		}

		if err != nil {
			return err
		}
	}
}

/* 行描述符 */
type RowDesc struct {
	TypeID uint64
	Name   string
}

/* 接收行描述符 */
func (conn *Connection) RecvRowDescs() error {
	num, err := protocol.RecvInt16(conn.server.Reader)
	if err != nil {
		return err
	}

	desc := make([]RowDesc, num)
	for i := 0; i < int(num); i++ {
		desc[i].TypeID, err = protocol.RecvUint64(conn.server.Reader)
		if err != nil {
			return err
		}
		desc[i].Name, err = protocol.RecvString(conn.server.Reader)
		if err != nil {
			return err
		}
	}
	conn.Desc = desc
	return nil
}

/* 接收 SQL 执行数据集合 */
func (conn *Connection) RecvRowsData() error {
	num := len(conn.Desc)
	row := make([]string, num)
	for i := 0; i < num; i++ {
		value, err := protocol.RecvDatum(conn.server.Reader)
		if err != nil {
			return err
		}
		row[i] = string(value)
	}
	conn.Rows = append(conn.Rows, row)
	return nil
}
