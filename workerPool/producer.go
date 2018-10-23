package workerPool

import (
	"fmt"
)

type Producer struct {
	MaxJob int
}

type MtProcessor struct {
	Id int
}
func (p *MtProcessor) Do() error {
	fmt.Printf("mt %d is running \n",p.Id)
	return nil
}
func NewProducer(maxJob int) (*Producer) {
	return &Producer{MaxJob:maxJob}
}

func (p Producer) Run() {

	for i:=1; i < int(p.MaxJob); i++{
		mt := MtProcessor{i}
		job := Job{Payload:&mt}
		fmt.Println("Producer: job " , i)
		JobQueue <- job
		//time.Sleep(time.Second*1)
	}
}


//func payloadHandler(){
//	//Go through each payload and queue items individually to be posted to S3
//	//for _, payload := range
//	payload := Payload(4)
//	work := Job{Payload: payload}
//
//	//Push the work onto the queue.
//	JobQueue <- work
//}