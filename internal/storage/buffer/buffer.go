package buffer

import (
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/storage/page"
	"hash/fnv"
	"sync"
	"sync/atomic"
)

var BufferPool, _ = NewBufferPool(10, 10)

type (
	/* BufferManager */
	BufferManager struct {
		bucketNum uint64
		capacity  bufferNumber
		victim    bufferNumber /* 受害者 */
		maxUsage  uint32

		bufferMap  []*bucket /* map: mapKey -> bufId */
		bufferPool []*Buffer
	}

	/* 通过把数据放到不同的 bucket 来降低冲突 */
	bucket struct {
		index uint64 /* 该 bucket 的在 BufferManager 中的序号 */

		mu    sync.RWMutex
		items map[string]bufferNumber
	}

	/*
	 * bufferNumber 是 buf id 的集合
	 * Buffer元素的索引号, 从0开始
	 */
	bufferNumber uint64

	// chunk store Page
	Buffer struct {
		pi *PageIdentifier /* 页面标识符, 使用时 bufferMap 持有锁 */

		isUsed   bool   /* 该 buffer 是否被使用过, 如果使用过, 那么在 bufferMap 中存在映射, 使用时 bufferMap 持有锁 */
		refNum   uint32 /* 引用数, 赋值操作都在锁住对应的 bucket 后, 原子操作 */
		usageNum uint32 /* usageNum 时钟扫描需要用到的引用数, 原子操作 */

		ioRoutine sync.WaitGroup /* 记录 io 进程 */
		isValid   bool           /* 页面是否有效, 被 ioRoutine 锁住 */

		Mu      sync.RWMutex /* 读写锁 */
		isDirty uint32       /* 是否为脏页 */
		Data    page.Page    /* 页面数据 */
	}
)

/*
 * BufferManager 相关的函数
 */

func NewBufferPool(capacity uint64, bucketNum uint64) (*BufferManager, error) {
	var bmgr = &BufferManager{
		bucketNum: bucketNum,
		capacity:  bufferNumber(capacity),
		victim:    0,
		maxUsage:  5,
	}

	bmgr.bufferMap = make([]*bucket, bucketNum)
	for i := uint64(0); i < bucketNum; i++ {
		bmgr.bufferMap[i] = &bucket{index: i, items: make(map[string]bufferNumber)}
	}

	bmgr.bufferPool = make([]*Buffer, capacity)
	for i := uint64(0); i < capacity; i++ {
		bmgr.bufferPool[i] = &Buffer{isUsed: false}
		bmgr.bufferPool[i].Data = page.NewPage()
	}

	return bmgr, nil
}

func (bmgr *BufferManager) getBucket(key string) *bucket {
	h := fnv.New64()
	h.Write([]byte(key))
	return bmgr.bufferMap[h.Sum64()%bmgr.bucketNum]
}

/* 获取buffer, pi 为页面标识符 */
func (bmgr *BufferManager) GetBuffer(pi *PageIdentifier) (*Buffer, error) {
	key := pi.hash()
	newBucket := bmgr.getBucket(key)
	newBucket.mu.RLock()

	// 在缓冲池已经存在对应的Buffer, 找到
	if bufId, ok := newBucket.items[key]; ok {
		var buf = bmgr.bufferPool[bufId]

		atomic.AddUint32(&buf.refNum, 1)
		buf.usageNumIncrement(bmgr.maxUsage)
		newBucket.mu.RUnlock()

		// 等待io线程
		buf.ioRoutine.Wait()
		if !buf.isValid {
			atomic.AddUint32(&buf.refNum, ^uint32(0))
			return nil, errlog.New("failed to read this Data into buffer pool")
		}
		return buf, nil
	}
	newBucket.mu.RUnlock()

	// 未找到, 需要自己获取Buffer
	var bufId bufferNumber
	var buf *Buffer
	var oldBucket *bucket
	var oldKey string
	var isUsed bool
	for {
		// 获取新的buffer, 淘汰算法
		bufId = bmgr.evict()
		buf = bmgr.bufferPool[bufId]
		atomic.AddUint32(&buf.refNum, 1)

		// 找到的是否为空闲buffer
		if isUsed = buf.isUsed; !isUsed {
			newBucket.mu.Lock()
			// 是否已经有线程找到buffer了
			if oldBufId, ok := newBucket.items[key]; ok {
				var oldBuf = bmgr.bufferPool[oldBufId]
				atomic.AddUint32(&buf.refNum, ^uint32(0))
				atomic.AddUint32(&oldBuf.refNum, 1)
				oldBuf.usageNumIncrement(bmgr.maxUsage)

				newBucket.mu.Unlock()

				oldBuf.ioRoutine.Wait()
				if !oldBuf.isValid {
					atomic.AddUint32(&oldBuf.refNum, ^uint32(0))
					return nil, errlog.New("failed to read this Data into buffer pool")
				}
				return oldBuf, nil
			}
		} else {
			// 写出脏页
			if atomic.LoadUint32(&buf.isDirty) == 1 {
				buf.Mu.RLock()
				// 如果isDirty赋值为0失败, 说明已经有其他线程在刷脏了
				if atomic.CompareAndSwapUint32(&buf.isDirty, 1, 0) {
					if err := writePage(buf.Data, buf.pi); err != nil {
						// 刷脏失败
						buf.Mu.RUnlock()
						atomic.AddUint32(&buf.refNum, ^uint32(0))
						return nil, err
					}
				}
				buf.Mu.RUnlock()
			}

			// 获取旧buffer所在的bucket
			oldKey = buf.pi.hash()
			oldBucket = bmgr.getBucket(oldKey)

			// 从左往右锁住bucekt, 避免死锁
			if oldBucket.index < newBucket.index {
				oldBucket.mu.Lock()
				newBucket.mu.Lock()
			} else if oldBucket.index > newBucket.index {
				newBucket.mu.Lock()
				oldBucket.mu.Lock()
			} else {
				newBucket.mu.Lock()
			}

			// 如果已经有线程找到buffer了, 那么返回它并撤销我们之前做的操作
			if oldBufId, ok := newBucket.items[key]; ok {
				var oldBuf = bmgr.bufferPool[oldBufId]
				atomic.AddUint32(&buf.refNum, ^uint32(0))
				atomic.AddUint32(&oldBuf.refNum, 1)

				oldBucket.mu.Unlock()
				if newBucket.index != oldBucket.index {
					newBucket.mu.Unlock()
				}

				oldBuf.ioRoutine.Wait()
				if !oldBuf.isValid {
					atomic.AddUint32(&oldBuf.refNum, ^uint32(0))
					return nil, errlog.New("failed to read this Data into buffer pool")
				}

				oldBuf.usageNumIncrement(bmgr.maxUsage)
				return oldBuf, nil
			}
		}

		// 如果没有其他线程引用该缓存区, 并且buf中的pi
		// 没有改变, 那么就完成对buf的选择, 否则重新获取
		if atomic.LoadUint32(&buf.refNum) == 1 &&
			isUsed == buf.isUsed &&
			(!buf.isUsed || oldKey == buf.pi.hash()) {
			break
		}

		// 如果线程进行到这里, 那么只能重新获取
		if !isUsed {
			newBucket.mu.Unlock()
		} else {
			oldBucket.mu.Unlock()
			if newBucket.index != oldBucket.index {
				newBucket.mu.Unlock()
			}
		}

		atomic.AddUint32(&buf.refNum, ^uint32(0))
	}

	// Okay, it's finally safe to rename the buffer.

	// 添加io写入任务
	buf.ioRoutine.Add(1)
	defer buf.ioRoutine.Done()
	buf.pi = pi

	// 修改bufferMap
	if !isUsed {
		newBucket.items[key] = bufId
		newBucket.mu.Unlock()
		buf.isUsed = true
	} else {
		// 在bufferMap中删除buffer原有映射, 添加新映射
		delete(oldBucket.items, oldKey)
		newBucket.items[key] = bufId

		// 解锁对应的bucket
		oldBucket.mu.Unlock()
		if newBucket.index != oldBucket.index {
			newBucket.mu.Unlock()
		}
	}

	// IO获取页面
	if err := readPage(buf.Data, pi); err != nil {
		// 获取页面失败, 取消引用
		buf.isValid = false
		atomic.AddUint32(&buf.refNum, ^uint32(0))
		return nil, err
	}

	buf.isValid = true
	buf.usageNumIncrement(bmgr.maxUsage)

	return buf, nil
}

/* 淘汰算法 clock */
func (bmgr *BufferManager) evict() bufferNumber {
	for {
		var bufId = bufferNumber(atomic.AddUint64((*uint64)(&bmgr.victim), 1) - 1)
		if bufId >= bmgr.capacity {
			if bufId == bmgr.capacity {
				atomic.AddUint64((*uint64)(&bmgr.victim), ^uint64(bmgr.capacity-1))
			}
			bufId = bufId % bmgr.capacity
		}

		var buf = bmgr.bufferPool[bufId]
		if atomic.LoadUint32(&buf.refNum) == 0 && !buf.usageNumDecrement(bmgr.maxUsage) {
			return bufId
		}
	}
}

// 对缓冲池刷脏
func (bmgr *BufferManager) Flush() {
	for i := bufferNumber(0); i < bmgr.capacity; i++ {
		// 找到buf, 引用增加1
		buf := bmgr.bufferPool[i]
		atomic.AddUint32(&buf.refNum, 1)

		// 如果没有被使用, 解除引用, 跳过
		if !buf.isUsed {
			atomic.AddUint32(&buf.refNum, ^uint32(0))
			continue
		}

		// 写出脏页
		if atomic.LoadUint32(&buf.isDirty) == 1 {
			buf.Mu.RLock()
			// 如果isDirty赋值为0失败, 说明已经有其他线程在刷脏了
			if atomic.CompareAndSwapUint32(&buf.isDirty, 1, 0) {
				// 不处理写出错误
				_ = writePage(buf.Data, buf.pi)
			}
			buf.Mu.RUnlock()
		}

		// 解除引用
		atomic.AddUint32(&buf.refNum, ^uint32(0))
	}
}

/*
 * Buffer 相关的函数
 */

/* buf 使用次数加 1 */
func (buf *Buffer) usageNumIncrement(maxUsage uint32) {
	for atomic.LoadUint32(&buf.usageNum) < maxUsage {
		for i := uint32(0); i < maxUsage; i++ {
			if atomic.CompareAndSwapUint32(&buf.usageNum, i, i+1) {
				return
			}
		}
	}
}

/* buf 使用次数减 1 */
func (buf *Buffer) usageNumDecrement(maxUsage uint32) bool {
	for atomic.LoadUint32(&buf.usageNum) > 0 {
		for i := uint32(0); i < maxUsage; i++ {
			if atomic.CompareAndSwapUint32(&buf.usageNum, i+1, i) {
				return true
			}
		}
	}
	return false
}

/* 标记为脏页 */
func (buf *Buffer) MarkDirty() {
	atomic.StoreUint32(&buf.isDirty, 1)
}

/* 释放对该 buf 的引用 */
func (buf *Buffer) Release() {
	atomic.AddUint32(&buf.refNum, ^uint32(0))
}
