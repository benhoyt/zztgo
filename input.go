package main // unit: Input

import (
	"github.com/gdamore/tcell"
)

const (
	KEY_BACKSPACE = '\x08'
	KEY_TAB       = '\t'
	KEY_ENTER     = '\r'
	KEY_CTRL_Y    = '\x19'
	KEY_ESCAPE    = '\x1b'
	KEY_F1        = '\xbb'
	KEY_F2        = '\xbc'
	KEY_F3        = '\xbd'
	KEY_F4        = '\xbe'
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

	keyChan chan byte
)

// implementation uses: Dos, Crt, Keys, Sounds

var (
	InputLastDeltaX, InputLastDeltaY int16
	InputKeyBuffer                   string
)

func InputUpdate() {
	InputUpdateWithKey(0)
}

func InputUpdateWithKey(keyRead byte) {
	InputDeltaX = 0
	InputDeltaY = 0
	InputShiftPressed = false

	if keyRead == 0 {
		checkForKeys := true
		for checkForKeys {
			select {
			case key := <-keyChan:
				InputKeyPressed = key
				InputKeyBuffer += string([]byte{InputKeyPressed})
			default:
				checkForKeys = false
			}
		}
	} else {
		InputKeyPressed = keyRead
		InputKeyBuffer += string([]byte{InputKeyPressed})
	}

	// TODO
	//for KeyPressed() {
	//	InputKeyPressed = ReadKey()
	//	if InputKeyPressed == '\x00' || InputKeyPressed == '\x01' || InputKeyPressed == '\x02' {
	//		InputKeyBuffer += Chr(Ord(ReadKey()) | 0x80)
	//	} else {
	//		InputKeyBuffer += string([]byte{InputKeyPressed})
	//	}
	//}

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
	key := <-keyChan
	InputUpdateWithKey(key)
}

func InputStartPoller(screen tcell.Screen) {
	keyChan = make(chan byte)
	go InputKeyPoller(screen, keyChan)
}

func InputKeyPoller(screen tcell.Screen, keyChan chan byte) {
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyRune:
				r := event.Rune()
				if r >= 32 && r <= 126 {
					keyChan <- byte(r)
				}
			default:
				key := tcellToKey[event.Key()]
				if key != 0 {
					keyChan <- key
				}
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

var tcellToKey = map[tcell.Key]byte{
	tcell.KeyBackspace: KEY_BACKSPACE,
	tcell.KeyCtrlY:     KEY_CTRL_Y,
	tcell.KeyDelete:    KEY_DELETE,
	tcell.KeyDown:      KEY_DOWN,
	tcell.KeyEnd:       KEY_END,
	tcell.KeyEnter:     KEY_ENTER,
	tcell.KeyEscape:    KEY_ESCAPE,
	tcell.KeyF1:        KEY_F1,
	tcell.KeyF2:        KEY_F2,
	tcell.KeyF3:        KEY_F3,
	tcell.KeyF4:        KEY_F4,
	tcell.KeyHome:      KEY_HOME,
	tcell.KeyInsert:    KEY_INSERT,
	tcell.KeyLeft:      KEY_LEFT,
	tcell.KeyPgDn:      KEY_PAGE_DOWN,
	tcell.KeyPgUp:      KEY_PAGE_UP,
	tcell.KeyRight:     KEY_RIGHT,
	tcell.KeyTab:       KEY_TAB,
	tcell.KeyUp:        KEY_UP,
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
