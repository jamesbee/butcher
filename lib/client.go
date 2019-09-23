package lib

import (
	"bytes"
	"context"
	"fmt"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"io"
	"os"
	"time"
)

var CRLF = []byte{'\r', '\n'}

func Dial(host string, port int) (*TelnetClient, error) {
	caller := TelnetClient{writer: os.Stdout}
	conn, err := telnet.DialTo(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	var w telnet.Writer = conn
	var r telnet.Reader = conn
	caller.w = w
	caller.r = r
	caller.conn = conn

	go func() {
		caller.CallTELNET(nil, w, r)
	}()
	return &caller, nil
}

type TelnetClient struct {
	writer io.Writer
	r      telnet.Reader
	w      telnet.Writer
	cancel context.CancelFunc
	conn   *telnet.Conn
}

func (c *TelnetClient) Cmd(cmd string) {
	var buffer bytes.Buffer
	buffer.WriteString(cmd)
	buffer.Write(CRLF)

	fmt.Println(cmd)
	_, err := oi.LongWrite(c.w, buffer.Bytes())
	if nil != err {
		panic(err)
	}
}

func (c *TelnetClient) CallTELNET(tctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	// Wait a bit to receive data from the server (that we would send to io.Stdout).
	time.Sleep(3 * time.Millisecond)

	// Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	var buffer [1]byte
	p := buffer[:]
	for {
		// Read 1 byte.
		n, err := c.r.Read(p)
		if n <= 0 && nil == err {
			continue
		} else if n <= 0 && nil != err {
			break
		}

		oi.LongWrite(c.writer, p)
	}
}

func (c *TelnetClient) Close() error {
	return c.conn.Close()
}
