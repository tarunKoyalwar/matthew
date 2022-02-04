package stdin

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
)

//Get Input From Stdin
//Asynchronously

var Wg *sync.WaitGroup = &sync.WaitGroup{}
var ChunkSize int = 1

type Receive struct {
	ByteData   []byte
	StringData string
	Done       bool
}

func GetStdinPipe() chan Receive {
	ch := make(chan Receive, 10)
	if ChunkSize == 1 {
		go ReadFromStdinReaderAsync(ch)
	} else {
		go ReadFromStdinBuffered(ch)
	}
	return ch
}

//Send Data From Stdin to Channel
func ReadFromStdinAsync(ch chan<- Receive) {
	Wg.Add(1)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		r := Receive{}

		r.StringData = fmt.Sprintf("%v\n", sc.Text())
		r.ByteData = []byte(r.StringData)
		//add newline
		// r.StringData = sc.Text()
		r.Done = false

		//send data to channel
		ch <- r
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	Wg.Done()
	defer close(ch)
}

func ReadFromStdinReaderAsync(ch chan<- Receive) {
	Wg.Add(1)
	defer Wg.Done()
	defer close(ch)

	r := bufio.NewReader(os.Stdin)

	for {
		bytes, err := r.ReadBytes(byte('\n'))

		rec := Receive{}
		rec.ByteData = bytes
		rec.StringData = string(bytes)

		ch <- rec

		if err != nil {
			break
		}
	}

	// r.
}

func ReadFromStdinBuffered(ch chan<- Receive) {
	Wg.Add(1)
	defer Wg.Done()
	defer close(ch)

	r := bufio.NewReader(os.Stdin)
	counter := 0

	buff := bytes.Buffer{}

	for {
		bytess, err := r.ReadBytes(byte('\n'))
		buff.Write(bytess)

		counter += 1

		if counter%ChunkSize == 0 {
			rec := Receive{}
			rec.ByteData = buff.Bytes()
			rec.StringData = string(rec.ByteData)
			ch <- rec

			buff = bytes.Buffer{}

		}

		if err != nil {
			break
		}
	}

	dat := buff.Bytes()
	if len(dat) != 0 {
		rec := Receive{}
		rec.ByteData = dat
		rec.StringData = string(rec.ByteData)
		ch <- rec
	}
}
