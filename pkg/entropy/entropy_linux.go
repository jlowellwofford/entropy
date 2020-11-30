package entropy

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

var entropy_device = "/dev/urandom"

// this is honestly easier through /proc, but in the spirit of completeness...
func getEntropy() (ent int, err error) {
	var fd int
	if fd, err = unix.Open(entropy_device, unix.O_RDWR, 0); err != nil {
		return ent, fmt.Errorf("could not open entropy device (%s): %v", entropy_device, err)
	}
	defer unix.Close(fd)

	var errno syscall.Errno
	_, _, errno = unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(RNDGETENTCNT), uintptr(unsafe.Pointer(&ent)))
	if errno != 0 {
		err = errno
	}
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