package internal_test

import (
	"io"
	"testing"

	"github.com/aileron-projects/go-l4proxy/internal"
	"github.com/aileron-projects/go-tester"
)

type testCloser struct {
	name     string
	closed   int
	closeErr error
}

func (c *testCloser) Close() error {
	c.closed++
	return c.closeErr
}

func TestCloserStore(t *testing.T) {
	t.Parallel()
	t.Run("store single value", func(t *testing.T) {
		s := internal.CloserStore[*testCloser]{}
		foo := &testCloser{name: "foo"}
		s.Store(foo)
		tester.AssertEqual(t, 1, s.Length())
		s.Delete(foo)
		tester.AssertEqual(t, 0, s.Length())
	})
	t.Run("store different value", func(t *testing.T) {
		s := internal.CloserStore[*testCloser]{}
		foo := &testCloser{name: "foo"}
		bar := &testCloser{name: "bar"}
		s.Store(foo)
		tester.AssertEqual(t, 1, s.Length())
		s.Store(bar)
		tester.AssertEqual(t, 2, s.Length())
		s.Delete(foo)
		s.Delete(bar)
		tester.AssertEqual(t, 0, s.Length())
	})
	t.Run("close", func(t *testing.T) {
		s := internal.CloserStore[*testCloser]{}
		foo := &testCloser{name: "foo"}
		bar := &testCloser{name: "bar"}
		s.Store(foo)
		s.Store(bar)
		tester.AssertEqual(t, 2, s.Length())
		err := s.CloseAll()
		tester.AssertEqualErr(t, nil, err)
		tester.AssertEqual(t, 0, s.Length())
		tester.AssertEqual(t, 1, foo.closed)
		tester.AssertEqual(t, 1, bar.closed)
	})
	t.Run("close error", func(t *testing.T) {
		s := internal.CloserStore[*testCloser]{}
		foo := &testCloser{name: "foo", closeErr: io.EOF}              // Return dummy error
		bar := &testCloser{name: "bar", closeErr: io.ErrUnexpectedEOF} // Return dummy error
		s.Store(foo)
		s.Store(bar)
		err := s.CloseAll()
		errs := err.(interface{ Unwrap() []error }).Unwrap()
		tester.AssertEqual(t, 2, len(errs))
		tester.AssertEqual(t, 1, foo.closed)
		tester.AssertEqual(t, 1, bar.closed)
	})
}
