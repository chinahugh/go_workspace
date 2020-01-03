package main

import (
	"flag"
	"fmt"
	"github.com/hpcloud/tail"
	"net"
	"time"
)

var addr string
var filename string
var msgChan chan string

func init() {
	flag.StringVar(&addr, "a", "localhost:80", ":addr")
	flag.StringVar(&filename, "f", "./test.log", ":filename")
}

func main() {
	flag.Parse()
	msgChan = make(chan string, 10000)
	go startClient()
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filenam:%s\n", tails, filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
		msgChan <- msg.Text
	}
}
func startClient() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		msg := <-msgChan
		_, err = conn.Write([]byte(msg))
		if err != nil {
			panic(err)
		}
	}
}
