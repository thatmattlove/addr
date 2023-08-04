package whois

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"
)

const (
	TCP string = "tcp"
)

type Whois struct {
	Host       string
	Port       uint
	TCPAddr    *net.TCPAddr
	Connection *net.TCPConn
}

func (w *Whois) Query(q string) (string, error) {
	err := w.Open()
	defer w.Close()
	if err != nil {
		return "", err
	}
	q = strings.Trim(q, "\r\n")
	q += "\r\n"
	d := fmt.Sprintf(" -v %s", q)
	_, err = w.Connection.Write([]byte(d))
	if err != nil {
		return "", err
	}
	rx := make([]byte, 1024)
	_, err = w.Connection.Read(rx)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	rx = bytes.Trim(rx, "\x00")
	return string(rx), nil
}

func (w *Whois) Open() error {
	conn, err := net.DialTCP(TCP, nil, w.TCPAddr)
	if err != nil {
		return err
	}
	w.Connection = conn
	now := time.Now()
	deadline := now.Add(time.Second * 10)
	w.Connection.SetReadDeadline(deadline)
	return nil
}

func (w *Whois) Close() error {
	if w.Connection == nil {
		return nil
	}
	err := w.Connection.Close()
	if err != nil {
		return err
	}
	w.Connection = nil
	return nil
}

func New(host string, port uint) (*Whois, error) {
	server, err := net.ResolveTCPAddr(TCP, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(TCP, nil, server)
	if err != nil {
		return nil, err
	}
	whois := &Whois{
		Host:       host,
		Port:       port,
		TCPAddr:    server,
		Connection: nil,
	}
	err = conn.Close()
	if err != nil {
		return nil, err
	}
	return whois, nil
}
