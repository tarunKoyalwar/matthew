package cmd

import (
	"fmt"

	"github.com/tarunKoyalwar/matthew/stdin"
)

var ch chan stdin.Receive

func GetInput() {
	stdin.ChunkSize = chunk

	// s := neo.NewStreamWriterWithCondition(4)

	if stdin.CheckStdin() {
		//has stdin
		ch = stdin.GetStdinPipe()
		defer stdin.Wg.Wait()

	}
}

func DebugPrint(text string) {
	if Debug {
		fmt.Println(text)
	}
}

//Will Automatically Add '\n' For Convinience
func DebugPrintWithArgs(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", args...)
	}
}
