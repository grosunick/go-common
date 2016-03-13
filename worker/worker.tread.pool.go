package worker

import (
	"github.com/grosunick/go-common/queue"
	"sync"
	"time"
)

// Thread pool handler struct
type TreadPoolWorker struct {
	maxWorkerAmount int32        // maximum thread amount
	queue           queue.IQueue // queue object
	mutex           sync.Cond
	executingTask   int32 // maximum number of threads
}

// Thread pool handler factory
func NewTreadPoolWorker(workerAmount int32, queue queue.IQueue) *TreadPoolWorker {
	return &TreadPoolWorker{
		workerAmount,
		queue,
		sync.Cond{L: &sync.Mutex{}},
		0,
	}
}

// Adds task cmd to thread pool
func (this *TreadPoolWorker) Add(cmd IHandler) {
	this.queue.Add(cmd)
}

// Run thread pool handler
func (this *TreadPoolWorker) Run() {
	// getting tasks from the queue in a loop
	for {
		// Fall asleep if number of simulteniosly excecuting tasks more then this.maxWorkerAmount
		this.mutex.L.Lock()
		if this.executingTask >= this.maxWorkerAmount {
			this.mutex.Wait()
		}
		this.mutex.L.Unlock()

		// excecute next task
		this.execute()

		// if queue is empty fall asleep
		if this.queue.Size() == 0 {
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// Excecute task in separate thread
func (this *TreadPoolWorker) execute() {
	if this.queue.Size() > 0 {
		// increase number of simulteniosly excecuting tasks
		this.mutex.L.Lock()
		this.executingTask++
		this.mutex.L.Unlock()

		// getting task from the queue
		if handler, ok := this.queue.Remove(); ok {
			go func() {
				_, err := handler.(IHandler).Run()
				this.notify(err)
			}()
		}
	}
}

// Notifies handler that task have been already done
func (this *TreadPoolWorker) notify(error error) {
	this.mutex.L.Lock()
	this.executingTask--
	this.mutex.L.Unlock()
	this.mutex.Signal()
}
