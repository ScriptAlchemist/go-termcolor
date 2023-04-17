package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	c "github.com/ScriptAlchemist/go-termcolor"
)

const (
	ExitCodeSuccess   = 0
	ExitCodeInterrupt = 1
)

func main() {
	fmt.Println("Starting program...")

	nameFlag := flag.String("n", "nyan", "name to print in random colors")
	flag.Parse()
	nameToPrint := *nameFlag

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		// Perform cleanup tasks before the program exits
		fmt.Print(c.CurOn)
		os.Exit(ExitCodeInterrupt)
	}()

	fmt.Print(c.CurOff)
	for {
		fmt.Print(c.RandomColor() + nameToPrint + c.Reset)
		time.Sleep(1 * time.Millisecond)
	}
}
