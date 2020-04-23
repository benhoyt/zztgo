package main // unit: Input

const (
	KEY_BACKSPACE = '\x08'
	KEY_TAB       = '\t'
	KEY_ENTER     = '\r'
	KEY_CTRL_Y    = '\x19'
	KEY_ESCAPE    = '\x1b'
	KEY_ALT_P     = '\x99'
	KEY_F1        = '\xbb'
	KEY_F2        = '\xbc'
	KEY_F3        = '\xbd'
	KEY_F4        = '\xbe'
	KEY_F5        = '\xbf'
	KEY_F6        = '\xc0'
	KEY_F7        = '\xc1'
	KEY_F8        = '\xc2'
	KEY_F9        = '\xc3'
	KEY_F10       = '\xc4'
	KEY_UP        = '\xc8'
	KEY_PAGE_UP   = '\xc9'
	KEY_LEFT      = '\xcb'
	KEY_RIGHT     = '\xcd'
	KEY_DOWN      = '\xd0'
	KEY_PAGE_DOWN = '\xd1'
	KEY_INSERT    = '\xd2'
	KEY_DELETE    = '\xd3'
	KEY_HOME      = '\xd4'
	KEY_END       = '\xd5'
)

var (
	InputDeltaX, InputDeltaY int16
	InputShiftPressed        bool
	InputShiftAccepted       bool
	InputKeyPressed          byte
)

// implementation uses: Dos, Crt, Keys, Sounds

var (
	InputLastDeltaX, InputLastDeltaY int16
	InputKeyBuffer                   string
)

func InputUpdate() {
	InputDeltaX = 0
	InputDeltaY = 0
	InputShiftPressed = false
	for KeyPressed() {
		InputKeyPressed = ReadKey()
		if InputKeyPressed == '\x00' || InputKeyPressed == '\x01' || InputKeyPressed == '\x02' {
			InputKeyBuffer += Chr(Ord(ReadKey()) | 0x80)
		} else {
			InputKeyBuffer += string(InputKeyPressed)
		}
	}
	if Length(InputKeyBuffer) != 0 {
		InputKeyPressed = InputKeyBuffer[0]
		if Length(InputKeyBuffer) == 1 {
			InputKeyBuffer = ""
		} else {
			InputKeyBuffer = Copy(InputKeyBuffer, Length(InputKeyBuffer)-1, 1)
		}
		switch InputKeyPressed {
		case KEY_UP, '8':
			InputDeltaX = 0
			InputDeltaY = -1
		case KEY_LEFT, '4':
			InputDeltaX = -1
			InputDeltaY = 0
		case KEY_RIGHT, '6':
			InputDeltaX = 1
			InputDeltaY = 0
		case KEY_DOWN, '2':
			InputDeltaX = 0
			InputDeltaY = 1
		}
	} else {
		InputKeyPressed = '\x00'
	}
	if InputDeltaX != 0 || InputDeltaY != 0 {
		KeysUpdateModifiers()
		InputShiftPressed = KeysShiftHeld
	}

	if InputDeltaX != 0 || InputDeltaY != 0 {
		InputLastDeltaX = InputDeltaX
		InputLastDeltaY = InputDeltaY
	}
}

func InputReadWaitKey() {
	for {
		InputUpdate()
		if InputKeyPressed != '\x00' {
			break
		}
	}
}

func init() {
	InputLastDeltaX = 0
	InputLastDeltaY = 0
	InputDeltaX = 0
	InputDeltaY = 0
	InputShiftPressed = false
	InputShiftAccepted = false
	InputKeyBuffer = ""
}
