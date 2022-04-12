package gogithub

// The majority of this code in this file is taken from an example GitHub repository
// there is no License attached to the repository, so I am assuming based on the small
// size of the 'example' reposity and its relationship to a tutorial blog post
// that it is in the Public Domain.
//
// Reference: https://github.com/godoylucase/workers-pool
// Blog post: https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

const (
	jobsCount   = 10
	workerCount = 2
)

func TestWorkerPool(t *testing.T) {
	wp := New(workerCount)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go wp.GenerateFrom(testJobs())

	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}

			i, err := strconv.ParseInt(string(r.Descriptor.ID), 10, 64)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			val := r.Value.(int)
			if val != int(i)*2 {
				t.Fatalf("wrong value %v; expected %v", val, int(i)*2)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}

func TestWorkerPool_TimeOut(t *testing.T) {
	wp := New(workerCount)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Nanosecond*10)
	defer cancel()

	go wp.Run(ctx)

	for {
		select {
		case r := <-wp.Results():
			if r.Err != nil && r.Err != context.DeadlineExceeded {
				t.Fatalf("expected error: %v; got: %v", context.DeadlineExceeded, r.Err)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}

func TestWorkerPool_Cancel(t *testing.T) {
	wp := New(workerCount)

	ctx, cancel := context.WithCancel(context.TODO())

	go wp.Run(ctx)
	cancel()

	for {
		select {
		case r := <-wp.Results():
			if r.Err != nil && r.Err != context.Canceled {
				t.Fatalf("expected error: %v; got: %v", context.Canceled, r.Err)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}

func testJobs() []Job {
	jobs := make([]Job, jobsCount)
	for i := 0; i < jobsCount; i++ {
		jobs[i] = Job{
			Descriptor: JobDescriptor{
				ID:       JobID(fmt.Sprintf("%v", i)),
				JType:    "anyType",
				Metadata: nil,
			},
			ExecFn: execFn,
			Args:   i,
		}
	}
	return jobs
}
