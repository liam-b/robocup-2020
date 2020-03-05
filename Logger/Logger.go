package Logger

import (
	"fmt"
	"time"
)

const (
	End         = 0 // goddamn edge cases making code look ugly
	Bold        = 1 << 1
	Underline   = 1 << 2
	Black       = 1 << 3
	Red         = 1 << 4
	Green       = 1 << 5
	Yellow      = 1 << 6
	Blue        = 1 << 7
	Purple      = 1 << 8
	Cyan        = 1 << 9
	White       = 1 << 10
	LightGreen  = 1 << 11
	LightYellow = 1 << 12
)

var mods [14]string = [14]string{
	"\x1b[0m",
	"\x1b[1m",
	"\x1b[4m",
	"\x1b[30m",
	"\x1b[31m",
	"\x1b[32m",
	"\x1b[33m",
	"\x1b[34m",
	"\x1b[35m",
	"\x1b[36m",
	"\x1b[37m",
	"\x1b[92m",
	"\x1b[93m",
}

// Example Usage: ColourString("123", Red | Bold)
func ColourString(input string, formats int) string {
	var final string = ""
	for i := 0; i < 13; i++ {
		if formats&(1<<i) != 0 { // fancy bit magic, if (i-1)th bit selected
			final += mods[i]
		}
	}
	return final + input + mods[End]
}

// self documenting
func Log(input string) {
	var final string = ""
	currTime := time.Now()
	final += ColourString("["+currTime.Format("15:04:05")+"]", Yellow|Bold) // who knows how the time thing works but it does
	final += " " + input
	fmt.Println(final)
}

func Error(input string) {
	fmt.Println(ColourString("ERROR: "+input, Bold|Red))
}

func Success(input string) {
	fmt.Println(ColourString("[*] ", Bold) + ColourString(input, LightGreen))
}
