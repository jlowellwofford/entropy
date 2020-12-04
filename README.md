# About

Entropy is a simple go pkg and cmdline interface for manipulating entropy in the linux kernel.

It achieves this by using the the IOCTL interface to /dev/(u)random.

# Some basic theory

The linux kernel simulates randomness by keeping a pool of data, estimating its information entropy, and generating data out of it using SHA hashes.

When information is added to the pool it gets "mixed" into the pool with a CRC-like algorithm; it doesn't actually add the raw data.

Optionally, when information is added, the entropy count of the pool can be incrememnted.  This isn't a literal add, but rather an asymptotic algorithm that approaches the pool size.

Information can be added to the pool by writing to `/dev/random` or `/dev/urandom`, but this will not increment the entropy count.  This package/command provide an interface to the IOCTLs that provide extended userspace functionality for manipulating randomness.  Of particular note, the `AddToEntCnt()` function adds bits to the entropy count and the `AddEntropy()` function adds bytes to the pool while also incrementing entropy bits.


# Command
The `entropy` command has the following usage information:

```
Usage: entropy <command> [<opts>...]

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
```

# Package

The `entropy` package provides a basic wrapper for all IOCTL functions provided by the kernel.

# Authors

- J. Lowell Wofford <lowell@lanl.gov>