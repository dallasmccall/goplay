package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Command struct {
	X int
	Y int
}

type Response struct {
	V int
}

const limit = 1000000

var child *exec.Cmd
var writer io.WriteCloser
var reader io.ReadCloser

func startChild() {
	child = exec.Command(os.Args[0], "child")

	var err error
	writer, err = child.StdinPipe()
	if nil != err {
		fmt.Println(err)
	}
	reader, err = child.StdoutPipe()
	if nil != err {
		fmt.Println(err)
	}

	err = child.Start()
	if nil != err {
		fmt.Println(err)
	}
}

func runParent() {
	encoder := gob.NewEncoder(writer)
	decoder := gob.NewDecoder(reader)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var response Response
		for i := 0; i < limit; i++ {
			decoder.Decode(&response)
		}
		child.Wait()
		wg.Done()
	}()

	start := time.Now()

	var command Command
	for i := 0; i < limit; i++ {
		command.X = i
		command.Y = limit - i
		err := encoder.Encode(&command)
		if nil != err {
			fmt.Println(err)
		}
	}
	wg.Wait()

	total := time.Since(start)
	nanosPerOp := total.Nanoseconds() / limit

	fmt.Println("Nanoseconds per operation: ", nanosPerOp)
}

func runChild() {
	encoder := gob.NewEncoder(os.Stdout)
	decoder := gob.NewDecoder(os.Stdin)

	var command Command
	var response Response
	for i := 0; i < limit; i++ {
		decoder.Decode(&command)
		response.V = command.X + command.Y
		encoder.Encode(&response)
	}
}

func main() {
	if len(os.Args) == 1 {
		startChild()
		runParent()
	} else {
		runChild()
	}
}
