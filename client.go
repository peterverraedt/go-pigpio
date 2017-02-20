package pigpio

import (
	"encoding/binary"
	"fmt"
	"net"
)

const DefaultHost = "127.0.0.1:8888"

type Client struct {
	Host       string
	serverAddr *net.TCPAddr
}

func New(host string) (*Client, error) {
	if host == "" {
		host = DefaultHost
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		return nil, fmt.Errorf("Invalid address %s: %s", host, err)
	}

	return &Client{serverAddr: tcpAddr}, nil
}

func (cli *Client) runCommand(cmd Command) (Command, error) {
	conn, err := net.DialTCP("tcp4", nil, cli.serverAddr)
	if err != nil {
		return cmd, fmt.Errorf("Connect addr failed %s", err)
	}
	defer conn.Close()

	err = binary.Write(conn, binary.LittleEndian, cmd)
	if err != nil {
		return cmd, err
	}

	retCmd := Command{}
	err = binary.Read(conn, binary.LittleEndian, &retCmd)
	return retCmd, err
}
