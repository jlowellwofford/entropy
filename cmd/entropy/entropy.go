package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jlowellwofford/entropy/pkg/entropy"
)

func usage() {
	fmt.Printf(`
Usage: %s <command> [<opts>...]

Commands:

	get(entropy)        - get the current system entropy.
	addto(entcnt) <num> - Add <num> to the current entropy count (must be root).

`, os.Args[0])
}

func usageFatal(str string, args ...interface{}) {
	fmt.Printf("\n"+str+"\n", args...)
	usage()
	os.Exit(1)
}

func fatal(str string, args ...interface{}) {
	fmt.Printf(str+"\n", args...)
	os.Exit(1)
}

func main() {
	var err error
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		fallthrough
	case "getentropy":
		var ent int
		if ent, err = entropy.GetEntCnt(); err != nil {
			fatal("failed to get entropy: %v", err)
		}
		fmt.Printf("%d\n", ent)
	case "addto":
		fallthrough
	case "addtoentcnt":
		if len(os.Args) != 3 {
			usageFatal("addtoentcnt requires a number to add as an option")
		}
		var add int
		if add, err = strconv.Atoi(os.Args[2]); err != nil {
			usageFatal("provided option does not appear to be a valid number: %v", err)
		}
		if err = entropy.AddToEntCnt(add); err != nil {
			fatal("entropy addition failed: %v", err)
		}
	case "help":
		fallthrough
	case "usage":
		usage()
	default:
		usageFatal("unrecognized command: %s", os.Args[1])
	}
}
