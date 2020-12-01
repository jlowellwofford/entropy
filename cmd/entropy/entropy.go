package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/jlowellwofford/entropy/pkg/entropy"
)

func usage() {
	fmt.Printf(`
Usage: %s <command> [<opts>...]

Commands:

	get(entropy)                    - get the current system entropy.
	addto(entcnt) <num>             - (superuser) Add <num> bits to the current entropy count.
									  Note: this does not literally increase entropy count by <num>.  The kernel adds using an asymptotic algorithm.  
									  See <drivers/char/random.c> for details.
	add(entropy) <file> [<quality>] - (superuser) Add the contents of <file> to entropy, incrementing entropy by the byte-length of the file. 
									  The optional <quality> specifies the percentage of total data to count as Shannon entropy (default: 1, which is highly unlikely).
	zap(entcnt)                     - (superuser) Clear the kernel entropy count.
	clear(pool)                     - (superuser) Clear the entropy pool and counters (on modern linux, this just does zapentcnt).
	reseed(crng)                    - (superuser) Reseed the CRNG. 

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
	case "add":
		fallthrough
	case "addentropy":
		if len(os.Args) != 3 {
			usageFatal("addtoentcnt requires a file path as an option")
		}
		var buf []byte
		if buf, err = ioutil.ReadFile(os.Args[2]); err != nil {
			fatal("could not read file %s: %v", os.Args[2], err)
		}
		if err = entropy.AddEntropy(len(buf), buf); err != nil {
			fatal("failed to add entropy: %v", err)
		}
	case "zap":
		fallthrough
	case "zapentcnt":
		if err = entropy.ZapEntCnt(); err != nil {
			fatal("failed to zap entropy count: %v", err)
		}
	case "clear":
		fallthrough
	case "clearpool":
		if err = entropy.ClearPool(); err != nil {
			fatal("failed to clear entropy pool: %v", err)
		}
	case "reseed":
		fallthrough
	case "reseedcrng":
		if err = entropy.ReseedCrng(); err != nil {
			fatal("failed to reseed CRNG: %v", err)
		}
	case "help":
		fallthrough
	case "usage":
		usage()
	default:
		usageFatal("unrecognized command: %s", os.Args[1])
	}
}
