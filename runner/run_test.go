package runner_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/tarunKoyalwar/matthew/runner"
)

func TestCmdRunner(t *testing.T) {
	r, e := runner.GetcmdStruct("cat /proc/version", "")
	if e != nil {
		panic(e)
	}

	var z bytes.Buffer
	r.Stdout = &z

	er := r.Run()

	if er != nil {
		t.Errorf("Command Failed to run %v\n", er.Error())
	}

	fmt.Println(z.String())

}
