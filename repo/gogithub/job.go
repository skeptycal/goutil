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
)

type (
	JobID       string
	jobType     string
	jobMetadata map[string]interface{}

	ExecutionFn func(ctx context.Context, args interface{}) (interface{}, error)

	JobDescriptor struct {
		ID       JobID
		JType    jobType
		Metadata map[string]interface{}
	}

	Result struct {
		Value      interface{}
		Err        error
		Descriptor JobDescriptor
	}

	Job struct {
		Descriptor JobDescriptor
		ExecFn     ExecutionFn
		Args       interface{}
	}
)

func (j Job) execute(ctx context.Context) Result {
	value, err := j.ExecFn(ctx, j.Args)
	if err != nil {
		return Result{
			Err:        err,
			Descriptor: j.Descriptor,
		}
	}

	return Result{
		Value:      value,
		Descriptor: j.Descriptor,
	}
}
