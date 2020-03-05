package main

import (
	"fmt"

	"github.com/liam-b/robocup-2020/Logger"
)

func main() {
	fmt.Println(Logger.ColourString("abc", Logger.Bold|Logger.Red))
	fmt.Println(Logger.ColourString("abc", Logger.End|Logger.Red))
}
