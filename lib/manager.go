package lib

import (
	b "github.com/kr/beanstalk"
	"time"
	"fmt"
)

type Manager struct {
	WorkerPool chan chan Job
	conn *b.Conn
	maxWorkers int
}

func NewManager(conn *b.Conn, maxWorkers int) *Manager {
	pool := make(chan chan Job, maxWorkers)

	return &Manager{pool, conn, maxWorkers}
}

func (m *Manager) dispatch() {
	for {
		id, body, err := m.conn.Reserve(5 * time.Second)
		if err == nil {
			fmt.Println(fmt.Sprintf("Add job #%d (Body: %s)", id, string(body)));
			job := NewJob(id, string(body))
			go func(job *Job) {
				jobChannel := <-m.WorkerPool
				jobChannel <- *job
			}(job)
		}
	}
}

func (m *Manager) Run() {
	for i := 0; i < m.maxWorkers; i++ {
		fmt.Println(fmt.Sprintf("Starting worker #%d", i+1))
		worker := NewWorker(uint(i+1), m.WorkerPool, m.conn)
		worker.Start()
	}

	m.dispatch()
}
