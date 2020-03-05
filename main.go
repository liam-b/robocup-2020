package main

import (
	"fmt"

	"github.com/liam-b/robocup-2020/Logger"
	"github.com/liam-b/robocup-2020/Utils"
)

// 10,000 IQ function
func keepImports() {
	fmt.Print("")
	Logger.ColourString("", 0)
	Utils.WaitUntil(func() bool { return 1 == 1 })
}

func main() {
	keepImports()
}
