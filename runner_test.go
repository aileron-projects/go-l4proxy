package l4proxy_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/aileron-projects/go-l4proxy"
	"github.com/aileron-projects/go-tester"
)

func TestServerRunner(t *testing.T) {
	t.Parallel()
	t.Run("no error", func(t *testing.T) {
		r := &l4proxy.ServerRunner{
			Serve:    func() error { return nil },
			Shutdown: func(context.Context) error { return nil },
		}
		err := r.Run(context.Background())
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("serve error", func(t *testing.T) {
		t.Run("already closed", func(t *testing.T) {
			r := &l4proxy.ServerRunner{
				Serve:    func() error { return http.ErrServerClosed },
				Shutdown: func(context.Context) error { return nil },
			}
			err := r.Run(context.Background())
			tester.AssertEqualErr(t, http.ErrServerClosed, err)
		})
		t.Run("non-nil error", func(t *testing.T) {
			testErr := errors.New("serve error")
			r := &l4proxy.ServerRunner{
				Serve:    func() error { return testErr },
				Shutdown: func(context.Context) error { return nil },
			}
			err := r.Run(context.Background())
			tester.AssertEqualErr(t, testErr, err)
		})
	})
	t.Run("shutdown error", func(t *testing.T) {
		t.Run("timeout", func(t *testing.T) {
			closeCalled := false
			r := &l4proxy.ServerRunner{
				Serve:           func() error { return http.ErrServerClosed },
				Shutdown:        func(ctx context.Context) error { <-ctx.Done(); return ctx.Err() },
				Close:           func() error { closeCalled = true; return nil },
				ShutdownTimeout: 10 * time.Millisecond,
			}
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			err := r.Run(ctx)
			tester.AssertEqualErr(t, l4proxy.ErrShutdownTimeout, err)
			tester.AssertEqual(t, true, closeCalled)
		})
		t.Run("non-nil error", func(t *testing.T) {
			testErr := errors.New("shutdown error")
			closeCalled := false
			r := &l4proxy.ServerRunner{
				Serve:    func() error { return http.ErrServerClosed },
				Shutdown: func(context.Context) error { return testErr },
				Close:    func() error { closeCalled = true; return nil },
			}
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			err := r.Run(ctx)
			tester.AssertEqualErr(t, testErr, err)
			tester.AssertEqual(t, false, closeCalled)
		})
	})
}
