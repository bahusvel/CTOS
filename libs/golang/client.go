package ctos

import (
	"encoding/json"
	"net"
	"sync"
)

type Client struct {
	host       string
	conn       net.Conn
	writeMutex sync.Mutex
}

func NewClient(host string) (*Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		conn = nil
		return nil, err
	}
	return &Client{host: host, conn: conn}, nil
}

func (this *Client) SendLog(line *LogLine) error {
	this.writeMutex.Lock()
	defer this.writeMutex.Unlock()
	if this.conn == nil {
		var err error
		this.conn, err = net.Dial("tcp", this.host)
		if err != nil {
			this.conn = nil
			return err
		}
	}
	data, err := json.Marshal(line)
	if err != nil {
		return err
	}
	_, err = this.conn.Write(data)
	if err != nil {
		// reconnect on error
		this.conn = nil
	}
	return err
}

func (this *Client) SendMetric() {
}

func (this *Client) SendEvent() {

}
