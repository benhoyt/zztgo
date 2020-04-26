package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

var screen tcell.Screen

func VideoInstall() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating screen: %v", err)
		os.Exit(1)
	}
	err = screen.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing screen: %v", err)
		os.Exit(1)
	}
	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	screen.Clear()

	// TODO: doesn't really belong in "video" install, but oh well
	InputStartPoller(screen)
}

func VideoClrScr() {
	screen.Clear()
}

func VideoWriteText(x, y int16, color byte, text string) {
	fg := color & 0x0F
	bg := color >> 4
	style := tcell.StyleDefault.Foreground(crtColorToTcell[fg]).
		Background(crtColorToTcell[bg])
	for i := 0; i < len(text); i++ {
		ch := text[i]
		r := codePage437ToRune[ch]
		screen.SetContent(int(x)+i, int(y), r, nil, style)
	}
	VideoShow() // TODO: is this inefficient?
}

type VideoCell struct {
	Rune  rune
	Style tcell.Style
}

func VideoMoveToVideo(x, y, width int16, cells []VideoCell) {
	for i := 0; i < int(x+width); i++ {
		cell := cells[i]
		screen.SetContent(int(x)+i, int(y), cell.Rune, nil, cell.Style)
	}
	VideoShow()
}

func VideoMoveToBuffer(x, y, width int16, cells []VideoCell) {
	for i := 0; i < int(x+width); i++ {
		r, _, style, _ := screen.GetContent(int(x)+i, int(y))
		cells[i] = VideoCell{r, style}
	}
}

func VideoShow() {
	screen.Show()
}

func VideoHideCursor() {
	screen.HideCursor()
}

func VideoUninstall() {
	screen.Fini()
}

const (
	CrtBlack        = 0
	CrtBlue         = 1
	CrtGreen        = 2
	CrtCyan         = 3
	CrtRed          = 4
	CrtMagenta      = 5
	CrtBrown        = 6
	CrtLightGray    = 7
	CrtDarkGray     = 8
	CrtLightBlue    = 9
	CrtLightGreen   = 10
	CrtLightCyan    = 11
	CrtLightRed     = 12
	CrtLightMagenta = 13
	CrtYellow       = 14
	CrtWhite        = 15
)

var crtColorToTcell = [16]tcell.Color{
	CrtBlack:        tcell.ColorBlack,
	CrtBlue:         tcell.ColorNavy,
	CrtGreen:        tcell.ColorGreen,
	CrtCyan:         tcell.ColorTeal,
	CrtRed:          tcell.ColorMaroon,
	CrtMagenta:      tcell.ColorPurple,
	CrtBrown:        tcell.ColorOlive,
	CrtLightGray:    tcell.ColorSilver,
	CrtDarkGray:     tcell.ColorGray,
	CrtLightBlue:    tcell.ColorBlue,
	CrtLightGreen:   tcell.ColorLime,
	CrtLightCyan:    tcell.ColorAqua,
	CrtLightRed:     tcell.ColorRed,
	CrtLightMagenta: tcell.ColorFuchsia,
	CrtYellow:       tcell.ColorYellow,
	CrtWhite:        tcell.ColorWhite,
}

var codePage437ToRune = [256]rune{
	0:   ' ',
	1:   '\u263A',
	2:   '\u263B',
	3:   '\u2665',
	4:   '\u2666',
	5:   '\u2663',
	6:   '\u2660',
	7:   '\u2022',
	8:   '\u25D8',
	9:   '\u25CB',
	10:  '\u25D9',
	11:  '\u2642',
	12:  '\u2640',
	13:  '\u266A',
	14:  '\u266B',
	15:  '\u263C',
	16:  '\u25BA',
	17:  '\u25C4',
	18:  '\u2195',
	19:  '\u203C',
	20:  '\u00B6',
	21:  '\u00A7',
	22:  '\u25AC',
	23:  '\u21A8',
	24:  '\u2191',
	25:  '\u2193',
	26:  '\u2192',
	27:  '\u2190',
	28:  '\u221F',
	29:  '\u2194',
	30:  '\u25B2',
	31:  '\u25BC',
	32:  ' ',
	33:  '!',
	34:  '"',
	35:  '#',
	36:  '$',
	37:  '%',
	38:  '&',
	39:  '\'',
	40:  '(',
	41:  ')',
	42:  '*',
	43:  '+',
	44:  ',',
	45:  '-',
	46:  '.',
	47:  '/',
	48:  '0',
	49:  '1',
	50:  '2',
	51:  '3',
	52:  '4',
	53:  '5',
	54:  '6',
	55:  '7',
	56:  '8',
	57:  '9',
	58:  ':',
	59:  ';',
	60:  '<',
	61:  '=',
	62:  '>',
	63:  '?',
	64:  '@',
	65:  'A',
	66:  'B',
	67:  'C',
	68:  'D',
	69:  'E',
	70:  'F',
	71:  'G',
	72:  'H',
	73:  'I',
	74:  'J',
	75:  'K',
	76:  'L',
	77:  'M',
	78:  'N',
	79:  'O',
	80:  'P',
	81:  'Q',
	82:  'R',
	83:  'S',
	84:  'T',
	85:  'U',
	86:  'V',
	87:  'W',
	88:  'X',
	89:  'Y',
	90:  'Z',
	91:  '[',
	92:  '\\',
	93:  ']',
	94:  '^',
	95:  '_',
	96:  '`',
	97:  'a',
	98:  'b',
	99:  'c',
	100: 'd',
	101: 'e',
	102: 'f',
	103: 'g',
	104: 'h',
	105: 'i',
	106: 'j',
	107: 'k',
	108: 'l',
	109: 'm',
	110: 'n',
	111: 'o',
	112: 'p',
	113: 'q',
	114: 'r',
	115: 's',
	116: 't',
	117: 'u',
	118: 'v',
	119: 'w',
	120: 'x',
	121: 'y',
	122: 'z',
	123: '{',
	124: '|',
	125: '}',
	126: '~',
	127: '\u2302',
	128: '\u00C7',
	129: '\u00FC',
	130: '\u00E9',
	131: '\u00E2',
	132: '\u00E4',
	133: '\u00E0',
	134: '\u00E5',
	135: '\u00E7',
	136: '\u00EA',
	137: '\u00EB',
	138: '\u00E8',
	139: '\u00EF',
	140: '\u00EE',
	141: '\u00EC',
	142: '\u00C4',
	143: '\u00C5',
	144: '\u00C9',
	145: '\u00E6',
	146: '\u00C6',
	147: '\u00F4',
	148: '\u00F6',
	149: '\u00F2',
	150: '\u00FB',
	151: '\u00F9',
	152: '\u00FF',
	153: '\u00D6',
	154: '\u00DC',
	155: '\u00A2',
	156: '\u00A3',
	157: '\u00A5',
	158: '\u20A7',
	159: '\u0192',
	160: '\u00E1',
	161: '\u00ED',
	162: '\u00F3',
	163: '\u00FA',
	164: '\u00F1',
	165: '\u00D1',
	166: '\u00AA',
	167: '\u00BA',
	168: '\u00BF',
	169: '\u2310',
	170: '\u00AC',
	171: '\u00BD',
	172: '\u00BC',
	173: '\u00A1',
	174: '\u00AB',
	175: '\u00BB',
	176: '\u2591',
	177: '\u2592',
	178: '\u2593',
	179: '\u2502',
	180: '\u2524',
	181: '\u2561',
	182: '\u2562',
	183: '\u2556',
	184: '\u2555',
	185: '\u2563',
	186: '\u2551',
	187: '\u2557',
	188: '\u255D',
	189: '\u255C',
	190: '\u255B',
	191: '\u2510',
	192: '\u2514',
	193: '\u2534',
	194: '\u252C',
	195: '\u251C',
	196: '\u2500',
	197: '\u253C',
	198: '\u255E',
	199: '\u255F',
	200: '\u255A',
	201: '\u2554',
	202: '\u2569',
	203: '\u2566',
	204: '\u2560',
	205: '\u2550',
	206: '\u256C',
	207: '\u2567',
	208: '\u2568',
	209: '\u2564',
	210: '\u2565',
	211: '\u2559',
	212: '\u2558',
	213: '\u2552',
	214: '\u2553',
	215: '\u256B',
	216: '\u256A',
	217: '\u2518',
	218: '\u250C',
	219: '\u2588',
	220: '\u2584',
	221: '\u258C',
	222: '\u2590',
	223: '\u2580',
	224: '\u03B1',
	225: '\u00DF',
	226: '\u0393',
	227: '\u03C0',
	228: '\u03A3',
	229: '\u03C3',
	230: '\u00B5',
	231: '\u03C4',
	232: '\u03A6',
	233: '\u0398',
	234: '\u03A9',
	235: '\u03B4',
	236: '\u221E',
	237: '\u03C6',
	238: '\u03B5',
	239: '\u2229',
	240: '\u2261',
	241: '\u00B1',
	242: '\u2265',
	243: '\u2264',
	244: '\u2320',
	245: '\u2321',
	246: '\u00F7',
	247: '\u2248',
	248: '\u00B0',
	249: '\u2219',
	250: '\u00B7',
	251: '\u221A',
	252: '\u207F',
	253: '\u00B2',
	254: '\u25A0',
	255: ' ',
}
