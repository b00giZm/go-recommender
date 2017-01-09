package lib

import (
	b "github.com/kr/beanstalk"
	"fmt"
	"os"
)

type Worker struct {
	Id uint
	WorkerPool chan chan Job
	JobChannel chan Job
	conn *b.Conn
	quit chan bool
}

func NewWorker(id uint, pool chan chan Job, conn *b.Conn) *Worker {
	jobChan := make(chan Job)
	quitChan := make(chan bool)

	return &Worker{id, pool, jobChan, conn, quitChan}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				err := job.Process()
				if err != nil {
					fmt.Fprintf(os.Stderr, fmt.Sprintf("Worker #%d: Error processing job #%d", w.Id, job.Id))

					return
				}

				fmt.Println(fmt.Sprintf("Worker: #%d: Finished processing job #%d", w.Id, job.Id));
				w.conn.Delete(job.Id)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
