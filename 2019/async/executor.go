package async

import (
	"sync"
)

type (
	// Executable is an interface that can be executed by the Executor
	Executable interface {
		Execute()
	}

	// Executor starts an async execution of the Executable
	Executor struct {
		executable Executable
		lock       *sync.RWMutex
		running    bool
	}
)

// NewExecutor creates a new Executor
func NewExecutor(executable Executable) *Executor {
	return &Executor{
		executable: executable,
		lock:       &sync.RWMutex{},
	}
}

// Running returns true if the async execution is in progress
func (e *Executor) Running() bool {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.running
}

// ExecuteAsync starts an execution
func (e *Executor) ExecuteAsync(done *sync.WaitGroup) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.running {
		return
	}

	e.running = true
	go e.execute(done)
}

func (e *Executor) execute(done *sync.WaitGroup) {
	e.executable.Execute()

	e.lock.Lock()
	e.running = false
	e.lock.Unlock()

	done.Done()
}
