package main

import (
	"github.com/nzlov/goal/al"
	"github.com/nzlov/goal/alut"
)

func main() {
	alut.Init()
	helloBuffer := alut.CreateBufferHelloWorld()
	helloSource := al.GenSource()
	helloSource.Sourcei(al.BUFFER, int(helloBuffer))
	helloSource.Play()
	alut.Sleep(1)
	alut.Exit()
}
