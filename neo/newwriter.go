package neo

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

type StreamWriter struct {
	Buffer      bytes.Buffer
	OtherBuffer bytes.Buffer
	ChunkSize   int
	function    bool
	OnTime      func(dat []byte)
	counter     int
}

func NewStreamWriter() *StreamWriter {
	s := StreamWriter{}
	s.Buffer = bytes.Buffer{}

	return &s
}

//similar to tee however instead or writing only on stdout
//writes continously on  stdout and sends to http address
func NewStreamWriterWithCondition(ops int) *StreamWriter {
	s := StreamWriter{
		function: true,
		counter:  0,
	}
	s.Buffer = bytes.Buffer{}
	s.OtherBuffer = bytes.Buffer{}
	s.ChunkSize = ops

	return &s

}

func (s *StreamWriter) Write(b []byte) (n int, err error) {
	nx, err2 := os.Stdout.Write(b)
	nx2, err3 := s.Buffer.Write(b)

	//If funciton execution is on
	if s.function {
		if s.ChunkSize == 1 {
			s.OnTime(b)
		} else {
			s.OtherBuffer.Write(b)
			if s.counter%s.ChunkSize == 0 {
				s.OnTime(s.OtherBuffer.Bytes())
				s.OtherBuffer.Reset()
			}
		}
		s.counter += 1
	}

	if nx != nx2 {
		fmt.Println("This should not happen")
	}
	if err2 != nil {
		return nx, err2
	}
	if err3 != nil {
		return nx, err3
	}

	return nx2, err3
}

func NewHttpClient() *http.Client {
	c := http.Client{
		Timeout:       time.Duration(15) * time.Second,
		CheckRedirect: nil,
	}

	return &c
}
