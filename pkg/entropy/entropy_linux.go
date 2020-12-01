package entropy

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

var entropy_device = "/dev/random"

func entropyIoctl(request int, data uintptr) (err error) {
	var fd int
	if fd, err = unix.Open(entropy_device, unix.O_RDWR, 0); err != nil {
		return fmt.Errorf("could not open entropy device (%s): %v", entropy_device, err)
	}
	defer unix.Close(fd)

	var errno syscall.Errno
	_, _, errno = unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(request), data)
	if errno != 0 {
		err = errno
	}
	return err
}

// this is honestly easier through /proc, but in the spirit of completeness...
func getEntCnt() (ent int, err error) {
	err = entropyIoctl(RNDGETENTCNT, uintptr(unsafe.Pointer(&ent)))
	return
}

func addToEntCnt(add int) (err error) {
	return entropyIoctl(RNDADDTOENTCNT, uintptr(unsafe.Pointer(&add)))
}

/* IOCTL argument structure
	struct rand_pool_info {
                      int    entropy_count;
                      int    buf_size;
                      __u32  buf[0];
                  };
*/
type randPoolInfo struct {
	entropyCount int
	bufSize      int
	buf          uint32 // first 4 bytes, followed by the rest
}

func addEntropy(cnt int, buf []byte) (err error) {
	blen := len(buf)
	// we need to pad to 4-byte chunks since this is a uint32 array
	if blen%4 != 0 {
		for i := 0; i < 4-(blen%4); i++ {
			buf = append(buf, 0x00)
		}
	}
	blen = len(buf)

	// make a byte slice and pack it
	// this may not be the cleanest way to do this...
	const structSize = int(unsafe.Sizeof(randPoolInfo{}))

	rpi := make([]byte, structSize+blen-1)

	hbo.PutUint32(rpi[0:], uint32(cnt))
	hbo.PutUint32(rpi[4:], uint32(blen))
	copy(rpi[8:], buf)

	err = entropyIoctl(RNDADDENTROPY, uintptr(unsafe.Pointer(&rpi[0])))
	return
}

func zapEntCnt() (err error) {
	return entropyIoctl(RNDZAPENTCNT, uintptr(unsafe.Pointer(nil)))
}

func clearPool() (err error) {
	return entropyIoctl(RNDCLEARPOOL, uintptr(unsafe.Pointer(nil)))
}

func reseedCrng() (err error) {
	return entropyIoctl(RNDRESEEDCRNG, uintptr(unsafe.Pointer(nil)))
}
