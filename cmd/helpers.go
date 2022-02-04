package cmd

import "github.com/tarunKoyalwar/matthew/stdin"

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
