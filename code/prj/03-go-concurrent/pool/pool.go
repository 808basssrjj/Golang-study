package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClosed = errors.New("pool has been closed")
)

// 使goroutine安全的共享资源
type Pool struct {
	factory  func() (io.Closer, error)
	resources chan io.Closer
	mtx      sync.Mutex
	closed   bool
}

func New(factory func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("invalid size for the resources pool")
	}

	return &Pool{
		factory:  factory,
		resources: make(chan io.Closer),
		closed:   false,
	}, nil
}

// 获取资源
func (p *Pool) AcquireResource() (io.Closer, error) {
	select {
	case res, ok := <-p.resources:
		// 如果取不到代表 p.resource已关闭
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("acquire resource from the pool")
		// 能取到就返回
		return res, nil
	default:
		// p.resource取不到资源 就用factory新建一个资源
		fmt.Println("acquire new resource")
		return p.factory()
	}
}

// 释放资源
func (p *Pool) ReleaseResource(resource io.Closer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		resource.Close()
		return
	}
	select {
	case p.resources <- resource: //如果能丢进去
		fmt.Println("release resource back to pool")
	default:
		fmt.Println("release reource closed")
		resource.Close()
	}
}

// 关闭数据池
func (p *Pool) Close() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources) //要先关闭channel 再读取
	for resource := range p.resources {
		resource.Close()
	}
}
