package udp

import (
	"net"
	"sync"
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestConn_Read(t *testing.T) {
	t.Parallel()
	// Obtain dummy PacketConn.
	pc, _ := net.ListenPacket("udp4", ":0")
	pc.Close()
	t.Run("read", func(t *testing.T) {
		packets := make(chan []byte, 1)
		c := &conn{
			pc:      pc,
			packets: packets,
		}
		tester.AssertEqual(t, pc.LocalAddr(), c.LocalAddr())

		go func() {
			packets <- []byte("test")
		}()
		buf := make([]byte, 10)
		n, err := c.Read(buf)
		tester.AssertEqual(t, 4, n)
		tester.AssertEqual(t, "test", string(buf[:n]))
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("read from closed conn", func(t *testing.T) {
		packets := make(chan []byte, 1)
		raddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345}
		c := &conn{
			pc:       pc,
			packets:  packets,
			channels: &sync.Map{},
			raddr:    raddr,
		}
		tester.AssertEqual(t, pc.LocalAddr(), c.LocalAddr())

		c.Close()
		go func() {
			packets <- []byte("test")
		}()
		buf := make([]byte, 10)
		n, err := c.Read(buf)
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, "", string(buf[:n]))
		tester.AssertEqualErr(t, net.ErrClosed, err)
	})
}

type recordPacketConn struct {
	net.PacketConn
	raddr   net.Addr
	written []byte
}

func (pc *recordPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	pc.raddr = addr
	pc.written = p
	return len(p), nil
}

func TestConn_Write(t *testing.T) {
	t.Parallel()
	// Obtain dummy PacketConn.
	pc, _ := net.ListenPacket("udp4", ":0")
	pc.Close()

	t.Run("write", func(t *testing.T) {
		packets := make(chan []byte, 1)
		rpc := &recordPacketConn{PacketConn: pc}
		raddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345}
		c := &conn{
			pc:      rpc,
			packets: packets,
			raddr:   raddr,
		}
		tester.AssertEqual(t, raddr.String(), c.RemoteAddr().String())

		n, err := c.Write([]byte("test"))
		tester.AssertEqual(t, 4, n)
		tester.AssertEqualErr(t, nil, err)
		tester.AssertEqual(t, "test", string(rpc.written))
		tester.AssertEqual(t, raddr.String(), rpc.raddr.String())
	})
	t.Run("write to closed conn", func(t *testing.T) {
		packets := make(chan []byte, 1)
		rpc := &recordPacketConn{PacketConn: pc}
		raddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345}
		c := &conn{
			pc:       rpc,
			packets:  packets,
			channels: &sync.Map{},
			raddr:    raddr,
		}
		tester.AssertEqual(t, raddr.String(), c.RemoteAddr().String())

		c.Close()
		go func() {
			packets <- []byte("test")
		}()
		n, err := c.Write([]byte("test"))
		tester.AssertEqual(t, 0, n)
		tester.AssertEqual(t, "", string(rpc.written))
		tester.AssertEqualErr(t, net.ErrClosed, err)
	})
}

func TestConn_Close(t *testing.T) {
	t.Parallel()
	t.Run("close once", func(t *testing.T) {
		channels := &sync.Map{}
		addr := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345}
		c := &conn{
			raddr:    &addr,
			channels: channels,
		}
		channels.Store(addr.String(), "test")
		_, ok := channels.Load(addr.String())
		tester.AssertEqual(t, true, ok)
		c.Close()
		_, ok = channels.Load(addr.String())
		tester.AssertEqual(t, false, ok)
	})
	t.Run("close multiple times", func(t *testing.T) {
		channels := &sync.Map{}
		addr := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12345}
		c := &conn{
			raddr:    &addr,
			channels: channels,
		}
		channels.Store(addr.String(), "test")
		_, ok := channels.Load(addr.String())
		tester.AssertEqual(t, true, ok)
		c.Close()
		_, ok = channels.Load(addr.String())
		tester.AssertEqual(t, false, ok)

		channels.Store(addr.String(), "test") // Store value again.
		c.Close()                             // Value should not be removed.
		_, ok = channels.Load(addr.String())
		tester.AssertEqual(t, true, ok)
	})
}
