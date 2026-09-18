package l4proxy

import (
	"context"
	"errors"
	"time"
)

// ErrShutdownTimeout tells that shutdown has timeout.
var ErrShutdownTimeout = errors.New("go-l4proxy/l4proxy: shutdown timeout")

// ServerRunner runs a server with an ability of graceful shutdown.
//
// Example:
//
//	svr := &http.Server{Addr: ":8080"}
//	r := &ServerRunner{
//		Serve:           svr.ListenAndServe,
//		Shutdown:        svr.Shutdown,
//		Close:           svr.Close,
//		ShutdownTimeout: 30 * time.Second,
//	}
//
//	sigCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
//	defer cancel()
//
//	if err := r.Run(sigCtx); err != nil {
//		panic(err)
//	}
type ServerRunner struct {
	// Serve starts a server.
	// Serve must blocks until the server closed.
	// Serve must not be nil.
	Serve func() error
	// Shutdown gracefully shutdowns the server.
	// Shutdown must not be nil.
	// It will be called only when the context given to
	// the ServerRunner.Run is done.
	// Otherwise, Shutdown is not called even the Server exited.
	// Typically [http.Server.Shutdown] should be set.
	Shutdown func(context.Context) error
	// Close immediately closes a server.
	// Unlike Shutdown, it should not block.
	// Close, if non-nil, will be called only when shutdown timeout occurred.
	// Typically [net/http.Server.Close] should be set.
	// Note that the [net/http.Server.Shutdown] does not close remaining
	// connection after shutdown timeout occurred but this Runner try to.
	Close func() error
	// ShutdownTimeout is the shutdown timeout duration.
	// ShutdownTimeout = 0 means no timeout.
	// ShutdownTimeout < 0 means shutdown immediately.
	ShutdownTimeout time.Duration
}

// Run runs a server.
// A server will be shutdown when the sigCtx is done.
// It returns non-nil error if r.Serve returns non-nil error.
// When a timeout occurred while shutting down, a [ErrShutdownTimeout] can be returned.
func (r *ServerRunner) Run(sigCtx context.Context) error {
	shutdownError := make(chan error, 1)
	cancelShutdown := make(chan struct{})

	go func() {
		select {
		case <-sigCtx.Done(): // May be SigTerm, SigInt, etc.
		case <-cancelShutdown: // Graceful shutdown canceled due to server error.
			return
		}
		shutdownCtx := context.WithoutCancel(sigCtx)
		if timeout := r.ShutdownTimeout; timeout != 0 {
			timeoutCtx, cancel := context.WithTimeoutCause(shutdownCtx, timeout, ErrShutdownTimeout)
			defer cancel()
			shutdownCtx = timeoutCtx
		}
		err := r.Shutdown(shutdownCtx) // Graceful shutdown.
		if context.Cause(shutdownCtx) == ErrShutdownTimeout {
			err = ErrShutdownTimeout
			if closeFn := r.Close; closeFn != nil {
				_ = closeFn() // Force close the server.
			}
		}
		shutdownError <- err
	}()

	err := r.Serve()
	select {
	case <-sigCtx.Done():
		return <-shutdownError // Shutdown is in progress.
	default:
		close(cancelShutdown)
		return err
	}
}
