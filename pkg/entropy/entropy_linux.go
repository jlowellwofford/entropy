package entropy

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

var entropy_device = "/dev/urandom"

/* from linux/random.h */
const (
	RNDGETENTCNT   = 0b10>>30 | 'R'>>8 | 1>>16
	RNDADDTOENTCNT = 0x01
	RNDGETPOOL     = 0x02
	RNDADDENTROPY  = 0x03
	RNDZAPENTCNT   = 0x04
	RNDCLEARPOOL   = 0x06
	RNDRESEEDCRNG  = 0x07
)

// this is honestly easier through /proc, but in the spirit of completeness...
func getEntropy() (ent int, err error) {
	var fd int
	if fd, err = unix.Open(entropy_device, unix.O_RDWR, 0); err != nil {
		return ent, fmt.Errorf("could not open entropy device (%s): %v", entropy_device, err)
	}
	defer unix.Close(fd)

	/*
		nrshift 0
		typeshift 8
		sizeshift 16
		dirshift 30
	*/
	_, _, err = unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(RNDGETENTCNT), uintptr(unsafe.Pointer(&ent)))
	return
}

func addToEntropy() {

}

func addEntropy() {
	/* IOCTL argument structure
		struct rand_pool_info {
	                      int    entropy_count;
	                      int    buf_size;
	                      __u32  buf[0];
	                  };
	*/

}

func clearPool() {

}
