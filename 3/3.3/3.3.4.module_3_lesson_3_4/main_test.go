package pool

import (
	"sync"
	"testing"
)

func TestWorkerPool_NewPool(t *testing.T) {
	if _, err := NewWorkerPool(0, 0); err == nil {
		t.Fatalf("expected error when creating pool with 0 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(-1, 0); err == nil {
		t.Fatalf("expected error when creating pool with -1 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(1, -1); err == nil {
		t.Fatalf("expected error when creating pool with -1 channel size, got: %v", err)
	}

	if _, err := NewWorkerPool(5, 0); err != nil {
		t.Fatalf("expected no error creating pool, got: %v", err)
	}
}

func TestPool_MultipleStartStop(t *testing.T) {
	p, err := NewWorkerPool(5, 0)
	if err != nil {
		t.Fatal("error creating pool:", err)
	}

	p.Start()
	p.Start()

	p.Stop()
	p.Stop()
}

type testT struct {
	executeFunc func() error

	shouldErr bool
	wg        *sync.WaitGroup

	mFailure       *sync.Mutex
	failureHandled bool
}

func newTestTask(executeFunc func() error, shouldErr bool, wg *sync.WaitGroup) *testT {

}
