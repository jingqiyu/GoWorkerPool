
package main

import (
	"workerPool"
	"time"
)

func main() {
	dispatcher := workerPool.NewDispatcher(3)
	dispatcher.Run()
	// produce job to be done
	producer := workerPool.NewProducer(40)
	producer.Run()
	time.Sleep(time.Second * 100)
}


/*
import (
	"time"
	"math/rand"
	"fmt"
)

type Processor interface {
	Do() error
}

type JobProcessor struct {
	JobId string   //每一个job分配一个id
	StartTime time.Time //job的开始时间
	Processor Processor //负责处理job的接口
}

var JobQueue chan JobProcessor

type Worker struct {
	WorkerPool chan chan JobProcessor
	JobChannel chan JobProcessor
	quit chan bool
}

func NewWorker(workerPool chan chan JobProcessor) Worker{
	return Worker{
		WorkerPool: workerPool, //工人在哪个对象池里工作,可以理解成部门
		JobChannel:make(chan JobProcessor),//工人的任务
		quit:make(chan bool),
	}
}

func (w *Worker) Start(){
	//开一个新的协程
	go func(){
		for{
			w.WorkerPool <-w.JobChannel
			select {
			//接收到了新的任务
			case job :=<- w.JobChannel:
				job.Processor.Do()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
				//接收到了任务
			case <-w.quit:
				return
			}
		}
	}()
}


type Dispatcher struct {
	maxWorkers int //最多的工人数
	WorkerPool chan chan JobProcessor //worker 池
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool :=make(chan chan JobProcessor,maxWorkers)
	return &Dispatcher{
		WorkerPool:pool,// 将工人放到一个池中
		maxWorkers:maxWorkers,//工作池的数量
	}
}
func (d *Dispatcher) Run(){
	// 开始运行
	for i :=0;i<d.maxWorkers;i++{
		worker := NewWorker(d.WorkerPool)
		//开始工作
		worker.Start()
	}
	//监控
	go d.dispatch()

}

func (d *Dispatcher) dispatch()  {
	for {
		select {
		case job :=<-JobQueue:
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			// 调度者接收到一个工作任务
			go func (job JobProcessor) {
				//从现有的对象池中拿出一个
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		default:
			//fmt.Println("ok!!")
		}
	}
}

func initialize()  {
	maxWorkers := 2
	maxQueue := 4
	//初始化一个调试者,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue =make(chan JobProcessor,maxQueue) //指定任务的队列长度
	//并让它一直接运行
	dispatch.Run()
}

type MtJob struct {

}
func (mt *MtJob) Do() error {
	fmt.Println("Mt is doing")
	return nil
}

func main() {
	//初始化对象池
	initialize()

	for i:=0;i<10;i++{
		MtJob := &MtJob{}
		JobQueue <- JobProcessor{
			Processor:MtJob,
		}
		time.Sleep(time.Second)
	}
	close(JobQueue)
}*/
