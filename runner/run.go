package runner

import (
	"os"
	"os/exec"
	"strings"
)

func GetcmdStruct(command string, dir string) (*exec.Cmd, error) {
	splitdata := strings.Split(command, " ")
	filter := []string{}
	for _, v := range splitdata {
		if strings.TrimSpace(v) != "" {
			filter = append(filter, v)
		}
	}

	path, err := exec.LookPath(filter[0])
	if err != nil {
		return &exec.Cmd{}, err
	}

	if dir == "" {
		dir, _ = os.Getwd()
	}

	ecmd := exec.Cmd{
		Path: path,
		Args: filter,
		Dir:  dir,
	}

	return &ecmd, nil

}
