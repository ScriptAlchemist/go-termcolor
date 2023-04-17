package termcolors_test

import (
	"fmt"
	"strings"

	C "github.com/ScriptAlchemist/go-termcolor"
)

func ExampleSetColor_noninteractive() {
	fmt.Printf("%q\n", C.Black+"black"+C.Reset)
	fmt.Printf("%q\n", C.Red+"red"+C.Reset)
	fmt.Printf("%q\n", C.Green+"green"+C.Reset)
	fmt.Printf("%q\n", C.Yellow+"yellow"+C.Reset)
	fmt.Printf("%q\n", C.Blue+"blue"+C.Reset)
	fmt.Printf("%q\n", C.Magenta+"magenta"+C.Reset)
	fmt.Printf("%q\n", C.Cyan+"cyan"+C.Reset)
	fmt.Printf("%q\n", C.White+"white"+C.Reset)

	// Output:
	// "black"
	// "red"
	// "green"
	// "yellow"
	// "blue"
	// "magenta"
	// "cyan"
	// "white"

}

func ExampleSetColor_shortcutnoninteractive() {
	fmt.Printf("%q\n", C.K+"black"+C.X)
	fmt.Printf("%q\n", C.R+"red"+C.X)
	fmt.Printf("%q\n", C.G+"green"+C.X)
	fmt.Printf("%q\n", C.Y+"yellow"+C.X)
	fmt.Printf("%q\n", C.B+"blue"+C.X)
	fmt.Printf("%q\n", C.M+"magenta"+C.X)
	fmt.Printf("%q\n", C.C+"cyan"+C.X)
	fmt.Printf("%q\n", C.W+"white"+C.X)

	// Output:
	// "black"
	// "red"
	// "green"
	// "yellow"
	// "blue"
	// "magenta"
	// "cyan"
	// "white"
}

func ExampleRandomColor_noninteractive() {
	fmt.Printf("%q\n", C.RandomColor())

	// Output:
	// ""
}

func ExampleSetColor_interactive() {
	C.SetColor()
	fmt.Printf("%q\n", C.Black+"black"+C.Reset)
	fmt.Printf("%q\n", C.Red+"red"+C.Reset)
	fmt.Printf("%q\n", C.Green+"green"+C.Reset)
	fmt.Printf("%q\n", C.Yellow+"yellow"+C.Reset)
	fmt.Printf("%q\n", C.Blue+"blue"+C.Reset)
	fmt.Printf("%q\n", C.Magenta+"magenta"+C.Reset)
	fmt.Printf("%q\n", C.Cyan+"cyan"+C.Reset)
	fmt.Printf("%q\n", C.White+"white"+C.Reset)

	// Output:
	// "\x1b[30mblack\x1b[0m"
	// "\x1b[31mred\x1b[0m"
	// "\x1b[32mgreen\x1b[0m"
	// "\x1b[33myellow\x1b[0m"
	// "\x1b[34mblue\x1b[0m"
	// "\x1b[35mmagenta\x1b[0m"
	// "\x1b[36mcyan\x1b[0m"
	// "\x1b[37mwhite\x1b[0m"

}
func ExampleSetColor_shortcutinteractive() {
	fmt.Printf("%q\n", C.K+"black"+C.X)
	fmt.Printf("%q\n", C.R+"red"+C.X)
	fmt.Printf("%q\n", C.G+"green"+C.X)
	fmt.Printf("%q\n", C.Y+"yellow"+C.X)
	fmt.Printf("%q\n", C.B+"blue"+C.X)
	fmt.Printf("%q\n", C.M+"magenta"+C.X)
	fmt.Printf("%q\n", C.C+"cyan"+C.X)
	fmt.Printf("%q\n", C.W+"white"+C.X)

	// Output:
	// "\x1b[30mblack\x1b[0m"
	// "\x1b[31mred\x1b[0m"
	// "\x1b[32mgreen\x1b[0m"
	// "\x1b[33myellow\x1b[0m"
	// "\x1b[34mblue\x1b[0m"
	// "\x1b[35mmagenta\x1b[0m"
	// "\x1b[36mcyan\x1b[0m"
	// "\x1b[37mwhite\x1b[0m"
}

func ExampleRandomColor_interactive() {
	string := C.RandomColor(true)
	colorAdded := strings.HasPrefix(string, "\x1b[3")
	fmt.Println(colorAdded)

	// Output:
	// true
}
