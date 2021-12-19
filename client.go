package pigpio

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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

// See http://abyz.co.uk/rpi/pigpio/sif.html

type Request struct {
	Code              CmdNum
	P1                uint32
	P2                uint32
	Extension         interface{}                   // Pass struct{} with extension values
	ResponseExtension func(size uint32) interface{} // Returns struct{} with appropriate size
}

type Response struct {
	Result    uint32
	Extension interface{} // A copy of ResponseExtension(size)
}

type Command struct {
	Code uint32
	P1   uint32
	P2   uint32
	P3   uint32
}

func (cli *Client) runCommand(request Request) (response Response, err error) {
	conn, err := net.DialTCP("tcp4", nil, cli.serverAddr)
	if err != nil {
		return response, fmt.Errorf("Connect addr failed: %w", err)
	}

	defer conn.Close()

	cmd := Command{
		Code: uint32(request.Code),
		P1:   request.P1,
		P2:   request.P2,
	}

	if request.Extension != nil {
		cmd.P3 = uint32(binary.Size(request.Extension))
	}

	err = binary.Write(conn, binary.LittleEndian, cmd)
	if err != nil {
		return response, err
	}

	if request.Extension != nil {
		err = binary.Write(conn, binary.LittleEndian, request.Extension)
		if err != nil {
			return response, err
		}
	}

	retCmd := Command{}

	err = binary.Read(conn, binary.LittleEndian, &retCmd)
	if err != nil {
		return response, err
	}

	if retCmd.Code != cmd.Code || retCmd.P1 != cmd.P1 || retCmd.P2 != cmd.P2 {
		return response, fmt.Errorf("%w: {Code, P1, P2} do not match", ErrInvalidResponse)
	}

	err = cli.checkError(retCmd)
	if err != nil {
		return response, err
	}

	if request.ResponseExtension == nil {
		response.Result = retCmd.P3
	} else {
		buf := request.ResponseExtension(retCmd.P3)

		err = binary.Read(conn, binary.LittleEndian, &buf)

		response.Extension = buf
	}

	return response, err
}

var (
	ErrInvalidResponse = errors.New("invalid response for command")
	ErrInvalidError    = errors.New("invalid error number")
	ErrPigpioError     = errors.New("pigpio error")
)

func (cli *Client) checkError(cmd Command) error {
	for _, c := range NoErrorCmds {
		if cmd.Code == uint32(c) {
			return nil
		}
	}

	if cmd.P3 <= 2147483647 {
		return nil
	}

	var errNum int32

	r, w := io.Pipe()
	defer w.Close()

	binary.Write(w, binary.LittleEndian, cmd.P3)
	binary.Read(r, binary.LittleEndian, &errNum)

	errStr, ok := errorStrings[int(errNum)]
	if !ok {
		return fmt.Errorf("%w: %d", ErrInvalidError, errNum)
	}

	return fmt.Errorf("%w [%d]: %s", ErrPigpioError, errNum, errStr)
}
