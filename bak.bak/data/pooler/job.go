package pooler

import (
	"time"
)

type (
	Any interface{}

	JobSpec interface {

		// Identifier is the main unique identifier associated
		// with this specific job specification. It may be
		// considered analogous to a primary key in a database.
		Identifier() Any

		// Name is an optional text name for this job that will
		// be used in error messages and reports.
		Name() string

		// ID is an optional unique identifier for this job. It
		// may be used when sorting and searching performance
		// will be hindered by using an interface{}.
		Id() int

		// Completed tags the job specification as completed.
		// Not all workers may have finished, but the goals of
		// the job specification have been reached. All work is
		// stopped and no further work will be performed.
		Completed() bool

		// Deadline is an optional time at which the job spec
		// expires and no more work will be performed towards
		// the goals of this job specification.
		Deadline() time.Time
	}

	jobExampleSpec struct {
		identifier Any       // primary unique identifier
		name       string    // optional name
		id         int       // optional unique identifier
		completed  bool      // specification is completed.
		deadline   time.Time // optional deadline
		workerType Worker    // the type of worker implemented
	}
)

// CreateExampleJobs mimics the creation of 'amount'
// concurrent jobs.
// func CreateJobs(amount int, js JobSpec) []string {
// 	var jobs []string

// 	for i := 0; i < amount; i++ {
// 		jobs = append(jobs, RandStringRunes(8))
// 	}
// 	return jobs
// }

// CreateExampleJobs mimics the creation of 'amount'
// concurrent jobs.
func CreateExampleJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}
