package misc

import (
	"context"
	"log"
	"sync"
)

type (
	handlerImpl struct {
		errorChan chan error
		ctx       context.Context
		canceler  context.CancelFunc
		waiter    sync.WaitGroup
	}

	// Handler is context handler with sync.WaitGroup and error notifier.
	Handler interface {
		// NotifyError is an error notifier method within a concurrency environment that avoids falling in deadlock.
		NotifyError(err error)
		// Error method is a Error channel for cancellation.
		Error() <-chan error
		// Done method is a Done channel for cancellation.
		Done() <-chan struct{}
		// GracefulWait is a graceful shutdown method while the dependents are safely released.
		// It acts like a "WaitGroup.Wait" method.
		GracefulWait()
		// IncreaseWait is an acquiring method for a dependent.
		// It acts like a "WaitGroup.Add(1)" method.
		IncreaseWait()
		// DecreaseWait is an release method for a dependent.
		// It acts like a "WaitGroup.Done" method.
		DecreaseWait()
	}
)

// NewHandler returns an initialized Handler interface.
func NewHandler(ctx context.Context) Handler {
	ctx, canceler := context.WithCancel(ctx)
	return &handlerImpl{
		ctx:       ctx,
		canceler:  canceler,
		errorChan: make(chan error, 5),
	}
}

func (h *handlerImpl) NotifyError(err error) { h.errorChan <- err }
func (h *handlerImpl) Error() <-chan error   { return h.errorChan }
func (h *handlerImpl) Done() <-chan struct{} { return h.ctx.Done() }
func (h *handlerImpl) GracefulWait() {
	if h.ctx.Err() == nil {
		h.canceler()
	}

	h.waiter.Wait()
	close(h.errorChan)

	for remainError := range h.errorChan {
		log.Println("remain errors ", remainError)
	}
}

func (h *handlerImpl) IncreaseWait() {
	h.waiter.Add(1)
}
func (h *handlerImpl) DecreaseWait() {
	h.waiter.Done()
}
