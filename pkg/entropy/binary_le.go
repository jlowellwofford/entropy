// +build 386 amd64 amd64p32 arm arm64 ppc64le mipsle mips64p32le riscv riscv64 wasm

/* binary_amd64.go: sets the host byte order for amd64
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, J. Lowell Wofford.
 * See LICENSE file for details.
 */

package entropy

import "encoding/binary"

var hbo = binary.LittleEndian
