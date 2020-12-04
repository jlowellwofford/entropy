// +build armbe arm64be ppc64 mips mips64 mips64p32 ppc s390 s390x sparc sparc64

/* binary_arm64.go: sets the host byte order for arm64
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
