package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

// 给定一系列的task,要求在规定的timeout内跑完,不然就报错
// 如果操作系统给了中断信息,也报错
type Runner struct {
	tasks     []func(int) //task列表
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time //计时(过规定时间)
}

func New(t time.Duration) *Runner {
	return &Runner{
		tasks:     make([]func(int), 0),
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
	}
}

// 添加任务
func (r *Runner) AddTask(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// relay interrupt from OS
	signal.Notify(r.interrupt, os.Interrupt)

	// run the tasks
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}
