package pooler

import (
	"hash/fnv"
	"time"
)

type Work struct {
	ID  int
	Job string
	JS  JobSpec
}

type Worker interface {
	Start()
	Stop()
	DoWork(params Params)
}

type (
	worker struct {
		ID            int
		WorkerChannel chan chan Work
		Channel       chan Work
		End           chan bool

		JS JobSpec
	}
)

func start(w *worker) {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:

				// gather parameters

				pp := make(map[string]Any)
				pp["job"] = job.Job
				pp["id"] = w.ID

				p := NewParams(false, pp)

				// do work
				w.DoWork(p)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *worker) Start() {
	// go func() {
	// 	for {
	// 		w.WorkerChannel <- w.Channel
	// 		select {
	// 		case job := <-w.Channel:
	// 			// do work
	// 			w.DoWork(params{job.Job, w.ID})
	// 		case <-w.End:
	// 			return
	// 		}
	// 	}
	// }()
	start(w)

}

func (w *worker) Stop() {
	dbLogf("worker [%d] is stopping", w.ID)
	w.End <- true
}

func (w *worker) DoWork(params Params) {
	ExampleWork(params)
}

// ExampleWork mimics any type of job that can be run concurrently
func ExampleWork(params Params) error {

	w, err := params.Get("job")
	if err != nil {
		return err
	}
	word := w.(string)

	id2, err := params.Get("id")
	if err != nil {
		return err
	}

	id := id2.(int)

	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)

	dbLogf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)

	return nil
}
