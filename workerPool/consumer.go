package workerPool

import "fmt"

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	Len        int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, Len: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	fmt.Println("len of workerPool", len(d.WorkerPool))
	for i := 0; i < d.Len; i++ {
		fmt.Println("Processor generate worker to do job ", i)
		worker := NewWorker(d.WorkerPool)
		fmt.Println("Generate NewWorker done ", i)
		worker.Start()
		fmt.Println("Worker started", i)
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			fmt.Println("Store a job into jobChannel")
			go func(job Job) {
				//try to obtain a worker job channel that is available.
				//this will block until a worker is idle
				//最开始初始化好maxnum个worker之后,其实就初始化了对应数量的jobChannel,并且每个worker都阻塞在了获取任务的位置.
				//当一个实际的任务到来时，会从workerPool中取一个worker消息（其实就是一个jobChannel类型）, 然后把job写入到这个jobChannel中去
				//之前阻塞的jobChannel收到了这条job之后，会实际处理job工作，并在代码设计牛逼之处，就在此，一个worker会死循环的去尝试将一个jobChannel作为消息写入到workerPool中
				//在阻塞的消息处理完后，这一步又被执行了，相当于，这个worker又处于空闲状态了..
				//所以整体流程是 初始化好N个worker后，每一个任务来临时，会消耗到一个worker，这个worker将job分配给自己的jobchannel执行，待执行完毕后，重新补充了worker的数量
				jobChannel := <-d.WorkerPool
				//dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
