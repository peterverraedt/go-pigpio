package pigpio

// GPIO levels

type Level uint32

const (
	OFF, LOW, CLEAR Level = 0, 0, 0
	ON, HIGH, SET   Level = 1, 1, 1
	TIMEOUT         Level = 2
)

// GPIO edges

type Edge uint32

const (
	RISING_EDGE Edge = iota
	FALLING_EDGE
	EITHER_EDGE
)

// GPIO Modes

type Mode uint32

const (
	INPUT Mode = iota
	OUTPUT
	ALT5
	ALT4
	ALT0
	ALT1
	ALT2
)

// GPIO Pins

type Pin uint32

const (
	GPIO17 Pin = 17
)

// pigpio command numbers

type CmdNum uint32

const (
	CmdModeSet CmdNum = 0
	CmdModeGet CmdNum = 1
	CmdRead    CmdNum = 3
	CmdWrite   CmdNum = 4
	CmdServo   CmdNum = 8
	CmdBR1     CmdNum = 10
	CmdBR2     CmdNum = 11
	CmdTick    CmdNum = 16
	CmdHwVer   CmdNum = 17
	CmdPigpV   CmdNum = 26
)

var NoErrorCmds = []CmdNum{
	CmdBR1,
	CmdBR2,
	CmdTick,
	CmdHwVer,
	CmdPigpV,
}

// pigpio error numbers

const (
	_PI_INIT_FAILED      = -1
	PI_BAD_USER_GPIO     = -2
	PI_BAD_GPIO          = -3
	PI_BAD_MODE          = -4
	PI_BAD_LEVEL         = -5
	PI_BAD_PUD           = -6
	PI_BAD_PULSEWIDTH    = -7
	PI_BAD_DUTYCYCLE     = -8
	_PI_BAD_TIMER        = -9
	_PI_BAD_MS           = -10
	_PI_BAD_TIMETYPE     = -11
	_PI_BAD_SECONDS      = -12
	_PI_BAD_MICROS       = -13
	_PI_TIMER_FAILED     = -14
	PI_BAD_WDOG_TIMEOUT  = -15
	_PI_NO_ALERT_FUNC    = -16
	_PI_BAD_CLK_PERIPH   = -17
	_PI_BAD_CLK_SOURCE   = -18
	_PI_BAD_CLK_MICROS   = -19
	_PI_BAD_BUF_MILLIS   = -20
	PI_BAD_DUTYRANGE     = -21
	_PI_BAD_SIGNUM       = -22
	_PI_BAD_PATHNAME     = -23
	PI_NO_HANDLE         = -24
	PI_BAD_HANDLE        = -25
	_PI_BAD_IF_FLAGS     = -26
	_PI_BAD_CHANNEL      = -27
	_PI_BAD_PRIM_CHANNEL = -27
	_PI_BAD_SOCKET_PORT  = -28
	_PI_BAD_FIFO_COMMAND = -29
	_PI_BAD_SECO_CHANNEL = -30
	_PI_NOT_INITIALISED  = -31
	_PI_INITIALISED      = -32
	_PI_BAD_WAVE_MODE    = -33
	_PI_BAD_CFG_INTERNAL = -34
	PI_BAD_WAVE_BAUD     = -35
	PI_TOO_MANY_PULSES   = -36
	PI_TOO_MANY_CHARS    = -37
	PI_NOT_SERIAL_GPIO   = -38
	_PI_BAD_SERIAL_STRUC = -39
	_PI_BAD_SERIAL_BUF   = -40
	PI_NOT_PERMITTED     = -41
	PI_SOME_PERMITTED    = -42
	PI_BAD_WVSC_COMMND   = -43
	PI_BAD_WVSM_COMMND   = -44
	PI_BAD_WVSP_COMMND   = -45
	PI_BAD_PULSELEN      = -46
	PI_BAD_SCRIPT        = -47
	PI_BAD_SCRIPT_ID     = -48
	PI_BAD_SER_OFFSET    = -49
	PI_GPIO_IN_USE       = -50
	PI_BAD_SERIAL_COUNT  = -51
	PI_BAD_PARAM_NUM     = -52
	PI_DUP_TAG           = -53
	PI_TOO_MANY_TAGS     = -54
	PI_BAD_SCRIPT_CMD    = -55
	PI_BAD_VAR_NUM       = -56
	PI_NO_SCRIPT_ROOM    = -57
	PI_NO_MEMORY         = -58
	PI_SOCK_READ_FAILED  = -59
	PI_SOCK_WRIT_FAILED  = -60
	PI_TOO_MANY_PARAM    = -61
	PI_SCRIPT_NOT_READY  = -62
	PI_BAD_TAG           = -63
	PI_BAD_MICS_DELAY    = -64
	PI_BAD_MILS_DELAY    = -65
	PI_BAD_WAVE_ID       = -66
	PI_TOO_MANY_CBS      = -67
	PI_TOO_MANY_OOL      = -68
	PI_EMPTY_WAVEFORM    = -69
	PI_NO_WAVEFORM_ID    = -70
	PI_I2C_OPEN_FAILED   = -71
	PI_SER_OPEN_FAILED   = -72
	PI_SPI_OPEN_FAILED   = -73
	PI_BAD_I2C_BUS       = -74
	PI_BAD_I2C_ADDR      = -75
	PI_BAD_SPI_CHANNEL   = -76
	PI_BAD_FLAGS         = -77
	PI_BAD_SPI_SPEED     = -78
	PI_BAD_SER_DEVICE    = -79
	PI_BAD_SER_SPEED     = -80
	PI_BAD_PARAM         = -81
	PI_I2C_WRITE_FAILED  = -82
	PI_I2C_READ_FAILED   = -83
	PI_BAD_SPI_COUNT     = -84
	PI_SER_WRITE_FAILED  = -85
	PI_SER_READ_FAILED   = -86
	PI_SER_READ_NO_DATA  = -87
	PI_UNKNOWN_COMMAND   = -88
	PI_SPI_XFER_FAILED   = -89
	_PI_BAD_POINTER      = -90
	PI_NO_AUX_SPI        = -91
	PI_NOT_PWM_GPIO      = -92
	PI_NOT_SERVO_GPIO    = -93
	PI_NOT_HCLK_GPIO     = -94
	PI_NOT_HPWM_GPIO     = -95
	PI_BAD_HPWM_FREQ     = -96
	PI_BAD_HPWM_DUTY     = -97
	PI_BAD_HCLK_FREQ     = -98
	PI_BAD_HCLK_PASS     = -99
	PI_HPWM_ILLEGAL      = -100
	PI_BAD_DATABITS      = -101
	PI_BAD_STOPBITS      = -102
	PI_MSG_TOOBIG        = -103
	PI_BAD_MALLOC_MODE   = -104
	_PI_TOO_MANY_SEGS    = -105
	_PI_BAD_I2C_SEG      = -106
	PI_BAD_SMBUS_CMD     = -107
	PI_NOT_I2C_GPIO      = -108
	PI_BAD_I2C_WLEN      = -109
	PI_BAD_I2C_RLEN      = -110
	PI_BAD_I2C_CMD       = -111
	PI_BAD_I2C_BAUD      = -112
	PI_CHAIN_LOOP_CNT    = -113
	PI_BAD_CHAIN_LOOP    = -114
	PI_CHAIN_COUNTER     = -115
	PI_BAD_CHAIN_CMD     = -116
	PI_BAD_CHAIN_DELAY   = -117
	PI_CHAIN_NESTING     = -118
	PI_CHAIN_TOO_BIG     = -119
	PI_DEPRECATED        = -120
	PI_BAD_SER_INVERT    = -121
	_PI_BAD_EDGE         = -122
	_PI_BAD_ISR_INIT     = -123
	PI_BAD_FOREVER       = -124
	PI_BAD_FILTER        = -125
	PI_BAD_PAD           = -126
	PI_BAD_STRENGTH      = -127
	PI_FIL_OPEN_FAILED   = -128
	PI_BAD_FILE_MODE     = -129
	PI_BAD_FILE_FLAG     = -130
	PI_BAD_FILE_READ     = -131
	PI_BAD_FILE_WRITE    = -132
	PI_FILE_NOT_ROPEN    = -133
	PI_FILE_NOT_WOPEN    = -134
	PI_BAD_FILE_SEEK     = -135
	PI_NO_FILE_MATCH     = -136
	PI_NO_FILE_ACCESS    = -137
	PI_FILE_IS_A_DIR     = -138
	PI_BAD_SHELL_STATUS  = -139
	PI_BAD_SCRIPT_NAME   = -140
	PI_BAD_SPI_BAUD      = -141
	PI_NOT_SPI_GPIO      = -142
	PI_BAD_EVENT_ID      = -143
	PI_CMD_INTERRUPTED   = -144
	PI_NOT_ON_BCM2711    = -145
	PI_ONLY_ON_BCM2711   = -146
)

// pigpio error text

var errorStrings = map[int]string{
	_PI_INIT_FAILED:      "pigpio initialisation failed",
	PI_BAD_USER_GPIO:     "GPIO not 0-31",
	PI_BAD_GPIO:          "GPIO not 0-53",
	PI_BAD_MODE:          "mode not 0-7",
	PI_BAD_LEVEL:         "level not 0-1",
	PI_BAD_PUD:           "pud not 0-2",
	PI_BAD_PULSEWIDTH:    "pulsewidth not 0 or 500-2500",
	PI_BAD_DUTYCYCLE:     "dutycycle not 0-range (default 255)",
	_PI_BAD_TIMER:        "timer not 0-9",
	_PI_BAD_MS:           "ms not 10-60000",
	_PI_BAD_TIMETYPE:     "timetype not 0-1",
	_PI_BAD_SECONDS:      "seconds < 0",
	_PI_BAD_MICROS:       "micros not 0-999999",
	_PI_TIMER_FAILED:     "gpioSetTimerFunc failed",
	PI_BAD_WDOG_TIMEOUT:  "timeout not 0-60000",
	_PI_NO_ALERT_FUNC:    "DEPRECATED",
	_PI_BAD_CLK_PERIPH:   "clock peripheral not 0-1",
	_PI_BAD_CLK_SOURCE:   "DEPRECATED",
	_PI_BAD_CLK_MICROS:   "clock micros not 1: 2: 4: 5: 8: or 10",
	_PI_BAD_BUF_MILLIS:   "buf millis not 100-10000",
	PI_BAD_DUTYRANGE:     "dutycycle range not 25-40000",
	_PI_BAD_SIGNUM:       "signum not 0-63",
	_PI_BAD_PATHNAME:     "can't open pathname",
	PI_NO_HANDLE:         "no handle available",
	PI_BAD_HANDLE:        "unknown handle",
	_PI_BAD_IF_FLAGS:     "ifFlags > 4",
	_PI_BAD_CHANNEL:      "DMA channel not 0-14",
	_PI_BAD_SOCKET_PORT:  "socket port not 1024-30000",
	_PI_BAD_FIFO_COMMAND: "unknown fifo command",
	_PI_BAD_SECO_CHANNEL: "DMA secondary channel not 0-14",
	_PI_NOT_INITIALISED:  "function called before gpioInitialise",
	_PI_INITIALISED:      "function called after gpioInitialise",
	_PI_BAD_WAVE_MODE:    "waveform mode not 0-1",
	_PI_BAD_CFG_INTERNAL: "bad parameter in gpioCfgInternals call",
	PI_BAD_WAVE_BAUD:     "baud rate not 50-250000(RX)/1000000(TX)",
	PI_TOO_MANY_PULSES:   "waveform has too many pulses",
	PI_TOO_MANY_CHARS:    "waveform has too many chars",
	PI_NOT_SERIAL_GPIO:   "no bit bang serial read in progress on GPIO",
	PI_NOT_PERMITTED:     "no permission to update GPIO",
	PI_SOME_PERMITTED:    "no permission to update one or more GPIO",
	PI_BAD_WVSC_COMMND:   "bad WVSC subcommand",
	PI_BAD_WVSM_COMMND:   "bad WVSM subcommand",
	PI_BAD_WVSP_COMMND:   "bad WVSP subcommand",
	PI_BAD_PULSELEN:      "trigger pulse length not 1-100",
	PI_BAD_SCRIPT:        "invalid script",
	PI_BAD_SCRIPT_ID:     "unknown script id",
	PI_BAD_SER_OFFSET:    "add serial data offset > 30 minute",
	PI_GPIO_IN_USE:       "GPIO already in use",
	PI_BAD_SERIAL_COUNT:  "must read at least a byte at a time",
	PI_BAD_PARAM_NUM:     "script parameter id not 0-9",
	PI_DUP_TAG:           "script has duplicate tag",
	PI_TOO_MANY_TAGS:     "script has too many tags",
	PI_BAD_SCRIPT_CMD:    "illegal script command",
	PI_BAD_VAR_NUM:       "script variable id not 0-149",
	PI_NO_SCRIPT_ROOM:    "no more room for scripts",
	PI_NO_MEMORY:         "can't allocate temporary memory",
	PI_SOCK_READ_FAILED:  "socket read failed",
	PI_SOCK_WRIT_FAILED:  "socket write failed",
	PI_TOO_MANY_PARAM:    "too many script parameters (> 10)",
	PI_SCRIPT_NOT_READY:  "script initialising",
	PI_BAD_TAG:           "script has unresolved tag",
	PI_BAD_MICS_DELAY:    "bad MICS delay (too large)",
	PI_BAD_MILS_DELAY:    "bad MILS delay (too large)",
	PI_BAD_WAVE_ID:       "non existent wave id",
	PI_TOO_MANY_CBS:      "No more CBs for waveform",
	PI_TOO_MANY_OOL:      "No more OOL for waveform",
	PI_EMPTY_WAVEFORM:    "attempt to create an empty waveform",
	PI_NO_WAVEFORM_ID:    "No more waveform ids",
	PI_I2C_OPEN_FAILED:   "can't open I2C device",
	PI_SER_OPEN_FAILED:   "can't open serial device",
	PI_SPI_OPEN_FAILED:   "can't open SPI device",
	PI_BAD_I2C_BUS:       "bad I2C bus",
	PI_BAD_I2C_ADDR:      "bad I2C address",
	PI_BAD_SPI_CHANNEL:   "bad SPI channel",
	PI_BAD_FLAGS:         "bad i2c/spi/ser open flags",
	PI_BAD_SPI_SPEED:     "bad SPI speed",
	PI_BAD_SER_DEVICE:    "bad serial device name",
	PI_BAD_SER_SPEED:     "bad serial baud rate",
	PI_BAD_PARAM:         "bad i2c/spi/ser parameter",
	PI_I2C_WRITE_FAILED:  "I2C write failed",
	PI_I2C_READ_FAILED:   "I2C read failed",
	PI_BAD_SPI_COUNT:     "bad SPI count",
	PI_SER_WRITE_FAILED:  "ser write failed",
	PI_SER_READ_FAILED:   "ser read failed",
	PI_SER_READ_NO_DATA:  "ser read no data available",
	PI_UNKNOWN_COMMAND:   "unknown command",
	PI_SPI_XFER_FAILED:   "SPI xfer/read/write failed",
	_PI_BAD_POINTER:      "bad (NULL) pointer",
	PI_NO_AUX_SPI:        "no auxiliary SPI on Pi A or B",
	PI_NOT_PWM_GPIO:      "GPIO is not in use for PWM",
	PI_NOT_SERVO_GPIO:    "GPIO is not in use for servo pulses",
	PI_NOT_HCLK_GPIO:     "GPIO has no hardware clock",
	PI_NOT_HPWM_GPIO:     "GPIO has no hardware PWM",
	PI_BAD_HPWM_FREQ:     "invalid hardware PWM frequency",
	PI_BAD_HPWM_DUTY:     "hardware PWM dutycycle not 0-1M",
	PI_BAD_HCLK_FREQ:     "invalid hardware clock frequency",
	PI_BAD_HCLK_PASS:     "need password to use hardware clock 1",
	PI_HPWM_ILLEGAL:      "illegal: PWM in use for main clock",
	PI_BAD_DATABITS:      "serial data bits not 1-32",
	PI_BAD_STOPBITS:      "serial (half) stop bits not 2-8",
	PI_MSG_TOOBIG:        "socket/pipe message too big",
	PI_BAD_MALLOC_MODE:   "bad memory allocation mode",
	_PI_TOO_MANY_SEGS:    "too many I2C transaction segments",
	_PI_BAD_I2C_SEG:      "an I2C transaction segment failed",
	PI_BAD_SMBUS_CMD:     "SMBus command not supported",
	PI_NOT_I2C_GPIO:      "no bit bang I2C in progress on GPIO",
	PI_BAD_I2C_WLEN:      "bad I2C write length",
	PI_BAD_I2C_RLEN:      "bad I2C read length",
	PI_BAD_I2C_CMD:       "bad I2C command",
	PI_BAD_I2C_BAUD:      "bad I2C baud rate: not 50-500k",
	PI_CHAIN_LOOP_CNT:    "bad chain loop count",
	PI_BAD_CHAIN_LOOP:    "empty chain loop",
	PI_CHAIN_COUNTER:     "too many chain counters",
	PI_BAD_CHAIN_CMD:     "bad chain command",
	PI_BAD_CHAIN_DELAY:   "bad chain delay micros",
	PI_CHAIN_NESTING:     "chain counters nested too deeply",
	PI_CHAIN_TOO_BIG:     "chain is too long",
	PI_DEPRECATED:        "deprecated function removed",
	PI_BAD_SER_INVERT:    "bit bang serial invert not 0 or 1",
	_PI_BAD_EDGE:         "bad ISR edge value: not 0-2",
	_PI_BAD_ISR_INIT:     "bad ISR initialisation",
	PI_BAD_FOREVER:       "loop forever must be last chain command",
	PI_BAD_FILTER:        "bad filter parameter",
	PI_BAD_PAD:           "bad pad number",
	PI_BAD_STRENGTH:      "bad pad drive strength",
	PI_FIL_OPEN_FAILED:   "file open failed",
	PI_BAD_FILE_MODE:     "bad file mode",
	PI_BAD_FILE_FLAG:     "bad file flag",
	PI_BAD_FILE_READ:     "bad file read",
	PI_BAD_FILE_WRITE:    "bad file write",
	PI_FILE_NOT_ROPEN:    "file not open for read",
	PI_FILE_NOT_WOPEN:    "file not open for write",
	PI_BAD_FILE_SEEK:     "bad file seek",
	PI_NO_FILE_MATCH:     "no files match pattern",
	PI_NO_FILE_ACCESS:    "no permission to access file",
	PI_FILE_IS_A_DIR:     "file is a directory",
	PI_BAD_SHELL_STATUS:  "bad shell return status",
	PI_BAD_SCRIPT_NAME:   "bad script name",
	PI_BAD_SPI_BAUD:      "bad SPI baud rate: not 50-500k",
	PI_NOT_SPI_GPIO:      "no bit bang SPI in progress on GPIO",
	PI_BAD_EVENT_ID:      "bad event id",
	PI_CMD_INTERRUPTED:   "pigpio command interrupted",
	PI_NOT_ON_BCM2711:    "not available on BCM2711",
	PI_ONLY_ON_BCM2711:   "only available on BCM2711",
}
