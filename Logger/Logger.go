package Logger

import "fmt"

const (
	End       = 1 << 0
	Bold      = 1 << 1
	Underline = 1 << 2
	Black     = 1 << 3
	Red       = 1 << 4
	Green     = 1 << 5
	Yellow    = 1 << 6
	Blue      = 1 << 7
	Purple    = 1 << 8
	Cyan      = 1 << 9
	White     = 1 << 10
)

var mods [12]string = [12]string{
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
}

func colourString(input string, formats int) string {
	var final string = ""
	for i := 0; i < 11; i++ {
		if formats&(1<<i) != 0 {
			final += mods[i]
		}
	}
	final += input
	final += mods[End]

	return final
}
