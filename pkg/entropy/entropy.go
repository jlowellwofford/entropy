/* entropy.go: package interface for Linux kernel entropy management
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, J. Lowell Wofford.
 * See LICENSE file for details.
 */

package entropy

/* GetEntCnt returns the current count for the system.
 *
 * This is the same as reading the contents of `/proc/sys/kernel/random/entropy_avail`, but is accomplished through the RNDGETENTCNT IOCTL.
 *
 * GetEntCnt is a wrapper around the RNDGETENTCNT IOCTL on `/dev/(u)random`.
 */
func GetEntCnt() (int, error) {
	return getEntCnt()
}

/* AddToEntCnt adds the specified integer to the entropy count.
 *
 * Note: this does not directly add to the value, but adds by an algorithm that asymptotically
 *       approaches the pool size.  See `devices/char/random.c` in the kernel source code for details.
 *
 * AddToEntCnt is a wrapper around the RNDADDTOENTCNT IOCTL on `/dev/(u)random`.
 */
func AddToEntCnt(add int) error {
	return addToEntCnt(add)
}

/* AddEntropy will add the contents of `buf` to the entropy pool.  The kernel takes these bytes and "mixes" tthem
 * using a CRC-like algorithm.  Additionally, cnt is added to the entropy count (see `AddToEntCnt()`).
 *
 * This is like writing data to `/dev/(u)random`, then calling RNDADDTOENTCOUNT.
 *
 * AddEntropy is a wrapper around the RNDADDENTROPY IOCTL on `/dev/(u)random`.
 */
func AddEntropy(cnt int, buf []byte) error {
	return addEntropy(cnt, buf)
}

/*
 * ZapEntCnt clears the entropy pool counters (i.e the entropy count).  This might be useful if, for instance, you
 * suspect your entropy pool is tainted or your entropy count has been artificially inflated.
 *
 * ZapEntCnt is a wrapper around the RNDZAPENTCNT IOCTL on `/dev/(u)random`.
 */
func ZapEntCnt() error {
	return zapEntCnt()
}

/*
 * ClearPool clears the entropy pool counters (i.e. the entropy count).  Historically, this also cleared all of the
 * bytes in the entropy pool, but on modern kernels this is just an alias for ZapEntCnt.
 *
 * ClearPool is a wrapper around the RNDCLEARPOOL IOCTL on `/dev/(u)random`.
 */
func ClearPool() error {
	return clearPool()
}

/*
 * ReseedCrng will re-seed the CRNG used to generate `/dev/urandom`.
 *
 * ReseedCrng is a wrapper around the RNDRESEEDCRNG IOCTL on `/dev/(u)random`
 */
func ReseedCrng() error {
	return reseedCrng()
}
