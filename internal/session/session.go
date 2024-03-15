package session

import (
	"bufio"
	"context"
	"github/suixinpr/manadb/internal/executor"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/parser"
	"github/suixinpr/manadb/internal/planner"
	"github/suixinpr/manadb/pkg/debug"
	"github/suixinpr/manadb/pkg/protocol"
)

type Session struct {
	cli *bufio.ReadWriter
}

/* 生成一个新的会话, 需要包含客户端的信息 */
func NewSesion(cli *bufio.ReadWriter) *Session {
	return &Session{
		cli: cli,
	}
}

/*
 * 数据库运行的错误全部传递上来在该函数中完成, 下面的函数不能直
 * 接发送错误给客户端.
 *
 * 当出现的错误是与客户端发送相关的错误时, 将错误再向上传递处理
 */
func (sess *Session) Main() error {
	for {
		/* 接收客户端消息 */
		msg, err := protocol.RecvMessage(sess.cli.Reader)
		if err != nil {
			return err
		}

		/* 处理客户端消息 */
		switch msg {
		case protocol.MessageSQL:
			/* 接收 SQL */
			sql, err := protocol.RecvString(sess.cli.Reader)
			if err != nil {
				return err
			}

			/* 执行 SQL */
			err = sess.ExecuteSQL(sql)
			if err != nil {
				sendError(sess.cli.Writer, err)
				continue
			}

		default:
			return errlog.New("unexpected message")
		}

		/* 发送消息表示完成 */
		protocol.SendMessage(sess.cli.Writer, protocol.MessageFinish)
		err = sess.cli.Writer.Flush()
		if err != nil {
			return err
		}
	}
}

func (sess *Session) ExecuteSQL(sql string) error {
	/* 语法分析 */
	stmtNode, err := parser.Parse(sql)
	if err != nil {
		return err
	}

	/* 生成计划 */
	plan, err := planner.BuildPlan(stmtNode)
	if err != nil {
		return err
	}

	debug.Print("physical plan", plan)

	/* 生成执行器 */
	exec, err := executor.BuildExecutor(plan)
	if err != nil {
		return err
	}

	/* 执行 */
	ch := executor.Run(context.Background(), exec)

	/* 发送数据 */
	desc := exec.Descriptor()
	err = sendDescription(sess.cli.Writer, desc)
	if err != nil {
		return err
	}

	return sendResultSQL(sess.cli.Writer, desc, ch)
}
