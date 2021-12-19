package pigpio

// SetMode sets the pin mode.
func (cli *Client) SetMode(pin Pin, mode Mode) error {
	request := Request{
		Code: CmdModeSet,
		P1:   uint32(pin),
		P2:   uint32(mode),
	}

	_, err := cli.runCommand(request)

	return err
}

// GetMode gets the pin mode.
func (cli *Client) GetMode(pin Pin) (Mode, error) {
	request := Request{
		Code: CmdModeGet,
		P1:   uint32(pin),
	}

	response, err := cli.runCommand(request)

	return Mode(response.Result), err
}

// Read gets the GPIO level, high or low.
func (cli *Client) Read(pin Pin) (isHigh bool, err error) {
	request := Request{
		Code: CmdRead,
		P1:   uint32(pin),
	}

	response, err := cli.runCommand(request)

	return response.Result > 0, err
}

// Write sets the GPIO level, high or low.
func (cli *Client) Write(pin Pin, isHigh bool) error {
	request := Request{
		Code: CmdWrite,
		P1:   uint32(pin),
		P2:   1,
	}

	if isHigh {
		request.P2 = 0
	}

	_, err := cli.runCommand(request)

	return err
}

// SetServoPulseWidth starts servo pulses, pulseWidth should be 0 (off), 500 (most anti-clockwise) to 2500 (most clockwise).
func (cli *Client) SetServoPulseWidth(pin uint32, pulseWidth uint32) error {
	request := Request{
		Code: CmdServo,
		P1:   uint32(pin),
		P2:   pulseWidth,
	}

	_, err := cli.runCommand(request)

	return err
}
