package executor

import (
	"context"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/planner/physical"
)

/*
 * 执行器的实现是 manadb 中极其重要的一部分, 它协调着协程的并发与下层非并发
 * 操作的交互.
 *
 * 当执行器执行时会调用函数 Run, 它是并发的执行的. 它会接收从 inputs 中传递的
 * 数据, 并在执行器做相应处理, 再将其传递给 output 通道. 当接收到错误或者自身
 * 产生错误时, 会依次进行如下处理:
 *   1. 通过 Cancel 函数取消所有的子节点协程. 不同时退出所有执行器节点是因为
 *      退出协程时资源的释放需要一次进行.
 *   2. 如果有必要, 进一步包装错误信息.
 *   3. 将错误向上传递.
 * 执行器间的数据以及错误信息全是通过 output 通道向上传递的.
 */
type Executor interface {
	base() *executor

	/* 执行 */
	run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error

	/* 执行器产生结果的行描述符 */
	Descriptor() *common.EntryDesc
}

func BuildExecutor(plan physical.PhysicalPlan) (Executor, error) {
	b := NewExecutorBuilder()
	return b.Build(plan)
}

/*
 * 启动一个执行器
 */
func Run(parent context.Context, exec Executor) <-chan *Result {
	if exec == nil {
		return nil
	}
	base := exec.base()

	/* 为父节点分配接收管道 */
	output := make(chan *Result)

	/* 启动子节点 */
	ctx, cancel := context.WithCancel(parent)
	outer := Run(ctx, base.outer)
	inner := Run(ctx, base.inner)

	/* 启动该节点 */
	go func() {
		defer func() {
			cancel()
			close(output)
		}()

		/* 执行并处理错误 */
		if err := exec.run(ctx, output, outer, inner); err != nil {
			output <- &Result{Err: err}
		}
	}()

	return output
}

type executor struct {
	/* 不同执行器的子节点数会不同, 可能为 0, 1, 2. 不会有其他情况 */
	outer Executor /* outer 节点 */
	inner Executor /* inner 节点 */

	/* 当前执行器的结果列描述符 */
	desc *common.EntryDesc
}

func (exec *executor) base() *executor {
	return exec
}

func (exec *executor) Descriptor() *common.EntryDesc {
	return exec.desc
}

/*
 * 使用该结构体在执行器之间传递结果
 */
type Result struct {
	Slot *row.EntrySlot
	Err  error
}
