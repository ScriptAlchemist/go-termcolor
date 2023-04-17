package termcolors

import (
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/term"
)

var (
	forcedColorOn = false
	Black         = ""
	K             = ""
	Red           = ""
	R             = ""
	Green         = ""
	G             = ""
	Yellow        = ""
	Y             = ""
	Blue          = ""
	B             = ""
	Magenta       = ""
	M             = ""
	Cyan          = ""
	C             = ""
	White         = ""
	W             = ""
	Reset         = ""
	X             = ""
	Clear         = ""
	CurOn         = ""
	CurOff        = ""
)

func init() {
	checkingInteractive(SetColor)
}

// SetColor will place the color ANSI escape characters in place of the empty
// strings. Allowing color to be used.
func SetColor() {
	Black = "\033[30m"
	K = Black
	Red = "\033[31m"
	R = Red
	Green = "\033[32m"
	G = Green
	Yellow = "\033[33m"
	Y = Yellow
	Blue = "\033[34m"
	B = Blue
	Magenta = "\033[35m"
	M = Magenta
	Cyan = "\033[36m"
	C = Cyan
	White = "\033[37m"
	W = White
	Reset = "\033[0m"
	X = Reset
	Clear = "\033[H\033[2J"
	CurOn = "\033[?25h"
	CurOff = "\033[?25l"
}

// Checking if interactive terminal and running the function taken in from
// input if the terminal is interactive
func checkingInteractive(fn func()) {
	if term.IsTerminal(int(os.Stdout.Fd())) && term.IsTerminal(int(os.Stdin.Fd())) {
		forcedColorOn = true
		fn()
	}
}

// RandomColor returns a random color ANSI escape between Black (30) and White (37)
func RandomColor(forceColorOns ...bool) string {
	if len(forceColorOns) > 0 {
		forcedColorOn = forceColorOns[0]
	}

	if forcedColorOn {
		var b [1]byte
		rand.Read(b[:])
		colorCode := int(b[0]) % 8
		color := fmt.Sprintf("\033[3%dm", colorCode)
		return color
	}
	return ""
}
