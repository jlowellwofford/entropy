/* entropy.go: this provides a commandline enterface to linux entropy management
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, J. Lowell Wofford.
 * See LICENSE file for details.
 */

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
		if len(os.Args) < 3 || len(os.Args) > 4 {
			usageFatal("addtoentcnt requires a file path as an option")
		}
		var quality = 1.0
		if len(os.Args) == 4 {
			quality, err = strconv.ParseFloat(os.Args[3], 32)
			if err != nil {
				fatal("couldn't parse quality option, must be a float from 0.0 - 1.0: %v", err)
			}
			if quality < 0 || quality > 1 {
				fatal("quality must be in the range of 0.0 - 1.0")
			}
		}
		var buf []byte
		if buf, err = ioutil.ReadFile(os.Args[2]); err != nil {
			fatal("could not read file %s: %v", os.Args[2], err)
		}
		bits := int(float64(len(buf)) * 8 * quality)
		fmt.Printf("adding %d bytes with %d bits of entropy\n", len(buf), bits)
		if err = entropy.AddEntropy(bits, buf); err != nil {
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
