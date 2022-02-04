package stdin

import "os"

//WIll Check if a stdin source exists
func CheckStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	} else {
		return true
	}
}
