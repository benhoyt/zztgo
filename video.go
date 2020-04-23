
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
}

func VideoMove(x, y, chars int16, data interface{}, toVideo bool) {
	// TODO
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
	CrtBlack = 0
	CrtBlue = 1
	CrtGreen = 2
	CrtCyan = 3
	CrtRed = 4
	CrtMagenta = 5
	CrtBrown = 6
	CrtLightGray = 7
	CrtDarkGray = 8
	CrtLightBlue = 9
	CrtLightGreen = 10
	CrtLightCyan = 11
	CrtLightRed = 12
	CrtLightMagenta = 13
	CrtYellow = 14
	CrtWhite = 15
)

var crtColorToTcell = [16]tcell.Color{
	CrtBlack:       tcell.ColorBlack,
	CrtBlue:        tcell.ColorNavy,
	CrtGreen:       tcell.ColorGreen,
	CrtCyan:        tcell.ColorTeal,
	CrtRed:         tcell.ColorMaroon,
	CrtMagenta:      tcell.ColorPurple,
	CrtBrown:      tcell.ColorOlive,
	CrtLightGray:       tcell.ColorSilver,
	CrtDarkGray:        tcell.ColorGray,
	CrtLightBlue:   tcell.ColorBlue,
	CrtLightGreen:  tcell.ColorLime,
	CrtLightCyan:   tcell.ColorAqua,
	CrtLightRed:    tcell.ColorRed,
	CrtLightMagenta: tcell.ColorFuchsia,
	CrtYellow: tcell.ColorYellow,
	CrtWhite:  tcell.ColorWhite,
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
	33:  '\u0021', // TODO: would be nicer if 32-126 were done as '!' etc
	34:  '\u0022',
	35:  '\u0023',
	36:  '\u0024',
	37:  '\u0025',
	38:  '\u0026',
	39:  '\u0027',
	40:  '\u0028',
	41:  '\u0029',
	42:  '\u002A',
	43:  '\u002B',
	44:  '\u002C',
	45:  '\u002D',
	46:  '\u002E',
	47:  '\u002F',
	48:  '\u0030',
	49:  '\u0031',
	50:  '\u0032',
	51:  '\u0033',
	52:  '\u0034',
	53:  '\u0035',
	54:  '\u0036',
	55:  '\u0037',
	56:  '\u0038',
	57:  '\u0039',
	58:  '\u003A',
	59:  '\u003B',
	60:  '\u003C',
	61:  '\u003D',
	62:  '\u003E',
	63:  '\u003F',
	64:  '\u0040',
	65:  '\u0041',
	66:  '\u0042',
	67:  '\u0043',
	68:  '\u0044',
	69:  '\u0045',
	70:  '\u0046',
	71:  '\u0047',
	72:  '\u0048',
	73:  '\u0049',
	74:  '\u004A',
	75:  '\u004B',
	76:  '\u004C',
	77:  '\u004D',
	78:  '\u004E',
	79:  '\u004F',
	80:  '\u0050',
	81:  '\u0051',
	82:  '\u0052',
	83:  '\u0053',
	84:  '\u0054',
	85:  '\u0055',
	86:  '\u0056',
	87:  '\u0057',
	88:  '\u0058',
	89:  '\u0059',
	90:  '\u005A',
	91:  '\u005B',
	92:  '\u005C',
	93:  '\u005D',
	94:  '\u005E',
	95:  '\u005F',
	96:  '\u0060',
	97:  '\u0061',
	98:  '\u0062',
	99:  '\u0063',
	100: '\u0064',
	101: '\u0065',
	102: '\u0066',
	103: '\u0067',
	104: '\u0068',
	105: '\u0069',
	106: '\u006A',
	107: '\u006B',
	108: '\u006C',
	109: '\u006D',
	110: '\u006E',
	111: '\u006F',
	112: '\u0070',
	113: '\u0071',
	114: '\u0072',
	115: '\u0073',
	116: '\u0074',
	117: '\u0075',
	118: '\u0076',
	119: '\u0077',
	120: '\u0078',
	121: '\u0079',
	122: '\u007A',
	123: '\u007B',
	124: '\u007C',
	125: '\u007D',
	126: '\u007E',
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