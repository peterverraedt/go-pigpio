package pigpio

//Command Code Refer: http://abyz.co.uk/rpi/pigpio/sif.html
const (
	CmdWrite uint32 = 4
	CmdServo uint32 = 8
)

//Command Struct Refer: http://abyz.co.uk/rpi/pigpio/sif.html#cmdCmd_t
type Command struct {
	Code uint32
	P1   uint32
	P2   uint32
	P3   uint32
}

//Sets the GPIO level, on or off.
func (cli *Client) Write(pin uint32, isHigh bool) (Command, error) {
	cmd := Command{CmdWrite, pin, 1, 0}
	if isHigh {
		cmd.P2 = 0
	}

	return cli.runCommand(cmd)
}

//Starts servo pulses, pulseWidth should be 0 (off), 500 (most anti-clockwise) to 2500 (most clockwise).
func (cli *Client) SetServoPulseWidth(pin uint32, pulseWidth uint32) (Command, error) {
	cmd := Command{CmdServo, pin, pulseWidth, 0}
	return cli.runCommand(cmd)
}
