# github.com/jlowellwofford/entropy/pkg/entropy

# Overview

This package provides an API that wraps all of the IOCTL calls on the `/dev/(u)random` devices.  These IOCTLs require important functionality beyond just reading/writing `/dev/(u)random`.  Of particular imporance, they allow for adding to and clearing the entropy count on the system.

The entropy count is intended to provide an estimate of how much information (in the Shannon sense) is stored in the entropy pool.  The `/dev/random` device will only provide at maximum the number of bits in the entropy count.

Note: all entropy count values are in bits, not bytes.

The kernel makes no attempt to estimate the entropy of data.  It's up to the user of the API to provide those estimates. That is why, e.g. the `AddEntropy` function, which adds bytes to the pool, requires the user to also provide the entropy count.

# Intended use

This package and the associated command was originaly created to provide an easy interface for artificially injecting entropy into the kernel to accelerate entropy gathering when booting large numbers of VMs for test clusters.  This pkg provides a generic interface that could be used to, e.g. create a goland version of programs like [rng-trools](https://github.com/nhorman/rng-tools) or [haveged](http://www.issihosts.com/haveged/).  

# See also

Command documentation [README](cmd/entropy/README.md)

Kernel source `devices/char/random.c`

Man page `random(4)`

# Authors

- J. Lowell Wofford <lowell@lanl.gov