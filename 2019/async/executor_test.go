package async

import (
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	mock_async "github.com/mfesenko/adventofcode/.mocks/2019/async"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func withExecutor(t *testing.T, test func(executor *Executor, executable *mock_async.MockExecutable)) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	defer goleak.VerifyNone(t)

	executable := mock_async.NewMockExecutable(ctrl)
	executor := NewExecutor(executable)
	test(executor, executable)
}

func TestRunningReturnsFalseIfExecutionWasNotStarted(t *testing.T) {
	withExecutor(t, func(executor *Executor, executable *mock_async.MockExecutable) {
		assert.False(t, executor.Running())
	})
}

func TestRunningReturnsTrueWhileExecutionIsInProgress(t *testing.T) {
	withExecutor(t, func(executor *Executor, executable *mock_async.MockExecutable) {
		execution := make(chan bool)
		executable.EXPECT().Execute().Do(func() {
			execution <- true
		})

		done := newWaitGroup()

		executor.ExecuteAsync(done)
		assert.True(t, executor.Running())

		<-execution
		assert.False(t, executor.Running())
	})
}

func TestWaitGroupIsMarkedAsDoneAfterExecutionCompletes(t *testing.T) {
	withExecutor(t, func(executor *Executor, executable *mock_async.MockExecutable) {
		done := newWaitGroup()
		executable.EXPECT().Execute()

		executor.ExecuteAsync(done)

		done.Wait()
	})
}

func newWaitGroup() *sync.WaitGroup {
	done := &sync.WaitGroup{}
	done.Add(1)
	return done
}

func TestCannotStartSecondExecutionIfFirstExecutionIsStillRunning(t *testing.T) {
	withExecutor(t, func(executor *Executor, executable *mock_async.MockExecutable) {
		execution := make(chan bool)
		executable.EXPECT().Execute().Do(func() {
			execution <- true
		})

		done := newWaitGroup()

		executor.ExecuteAsync(done)
		executor.ExecuteAsync(done)
		<-execution
	})
}

func TestCanStartSecondExecutionIfFirstExecutionCompleted(t *testing.T) {
	withExecutor(t, func(executor *Executor, executable *mock_async.MockExecutable) {
		execution := make(chan bool)
		executable.EXPECT().Execute().Times(2).Do(func() {
			execution <- true
		})

		firstDone := newWaitGroup()
		secondDone := newWaitGroup()

		executor.ExecuteAsync(firstDone)
		<-execution
		firstDone.Wait()
		executor.ExecuteAsync(secondDone)
		<-execution
		secondDone.Wait()
	})
}
